package machine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/luanguimaraesla/freegrow/pkg/node"
	"go.uber.org/zap"
)

var (
	requestTimeout = 10 * time.Second
)

func (m *Machine) getNodes(w http.ResponseWriter, r *http.Request) {

}

func (m *Machine) registerNode(w http.ResponseWriter, r *http.Request) {
	m.Logger().Info("registering node")

	w.Header().Set("Content-Type", "application/json")

	n := node.New()
	if err := json.NewDecoder(r.Body).Decode(n); err != nil {
		m.Logger().Error("failed decoding node", zap.Error(err))
		httpError(w, http.StatusBadRequest, err)
		return
	}

	m.Logger().Debug("saving node data", zap.String("node", n.Metadata.Name))

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	if err := m.Node(n.Metadata.Name).Put(ctx, n); err != nil {
		m.Logger().Error("failed saving object into storage", zap.Error(err))
		httpError(w, http.StatusInternalServerError, err)
		return
	}

	m.Logger().Info("node registered", zap.String("node", n.Metadata.Name))

	json.NewEncoder(w).Encode(n)
}

func (m *Machine) getNode(w http.ResponseWriter, r *http.Request) {

}

func (m *Machine) updateNode(w http.ResponseWriter, r *http.Request) {

}

func (m *Machine) deleteNode(w http.ResponseWriter, r *http.Request) {

}

func httpError(w http.ResponseWriter, code int, err error) {
	msg := map[string]string{
		"error": fmt.Sprintf("%v", err),
	}

	w.WriteHeader(code)
	json.NewEncoder(w).Encode(msg)
}
