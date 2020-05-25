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
	router.HandleFunc("/nodes/{name}", m.deleteNode).Methods("DELETE")

	router.HandleFunc("/resources/{kind}", m.getResources).Methods("GET")
	router.HandleFunc("/resources/{kind}", m.registerResource).Methods("POST")
	router.HandleFunc("/resources/{kind}/{name}", m.getResource).Methods("GET")
	router.HandleFunc("/resources/{kind}/{name}", m.deleteResource).Methods("DELETE")

	m.Logger().With(
		zap.String("bind", m.Spec.Bind),
	).Info("listening")
	if err := http.ListenAndServe(m.Spec.Bind, router); err != nil {
		return err
	}

	return nil
}
