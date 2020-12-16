package brain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	router.HandleFunc("/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", user.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}", user.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", user.UpdateUser).Methods("UPDATE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}
