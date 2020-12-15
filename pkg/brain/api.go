package brain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	router.HandleFunc("/users", user.Missing).Methods("GET")
	router.HandleFunc("/users", user.Missing).Methods("POST")
	router.HandleFunc("/users/{id}", user.Missing).Methods("GET")
	router.HandleFunc("/users/{id}", user.Missing).Methods("DELETE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}
