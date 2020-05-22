package machine

import (
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func (m *Machine) Listen() error {
	router := mux.NewRouter()

	router.HandleFunc("/nodes", m.getNodes).Methods("GET")
	router.HandleFunc("/nodes", m.registerNode).Methods("POST")
	router.HandleFunc("/nodes/{name}", m.getNode).Methods("GET")
	router.HandleFunc("/nodes/{name}", m.updateNode).Methods("PUT")
	router.HandleFunc("/nodes/{name}", m.deleteNode).Methods("DELETE")

	m.Logger().With(
		zap.String("bind", m.Spec.Bind),
	).Info("listening")
	if err := http.ListenAndServe(m.Spec.Bind, router); err != nil {
		return err
	}

	return nil
}
