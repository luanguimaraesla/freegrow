package machine

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/pkg/gadget/irrigator"
	"go.uber.org/zap"
)

type Named interface {
	Name() string
}

func (m *Machine) getResources(w http.ResponseWriter, r *http.Request) {}

func (m *Machine) registerResource(w http.ResponseWriter, r *http.Request) {
	m.Logger().Info("registering node")

	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	kind := vars["kind"]
	resource, err := decodeResource(kind, r.Body)
	if err != nil {
		m.Logger().Error("failed decoding generic resource", zap.Error(err))
		httpError(w, http.StatusBadRequest, err)
		return
	}

	m.Logger().Debug(
		"saving resource data",
		zap.String("kind", kind),
		zap.String("name", resource.Name()),
	)

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	if err := m.Resource(kind, resource.Name()).Put(ctx, resource); err != nil {
		m.Logger().Error("failed saving object into storage", zap.Error(err))
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	m.Logger().Debug(
		"resource registered",
		zap.String("kind", kind),
		zap.String("name", resource.Name()),
	)

	json.NewEncoder(w).Encode(resource)
}

func (m *Machine) getResource(w http.ResponseWriter, r *http.Request) {
	m.Logger().Info("getting resource")

	w.Header().Set("Content-Type", "application/json")

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	vars := mux.Vars(r)
	kind := vars["kind"]
	name := vars["name"]

	resource, err := m.Resource(kind, name).Get(ctx)
	if err != nil {
		m.Logger().Error("failed getting object from storage", zap.Error(err))
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	json.NewEncoder(w).Encode(resource)
}

func (m *Machine) deleteResource(w http.ResponseWriter, r *http.Request) {}

func decodeResource(kind string, r io.Reader) (Named, error) {
	switch kind {
	case "irrigator":
		resource := irrigator.New()
		if err := json.NewDecoder(r).Decode(resource); err != nil {
			return nil, err
		}

		return resource, nil
	default:
		return nil, fmt.Errorf("kind not found")
	}
}
