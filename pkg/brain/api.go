package brain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/pkg/gadget"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	// user routes
	router.HandleFunc("/users", user.GetUsers).Methods("GET")
	router.HandleFunc("/users", user.CreateUser).Methods("POST")
	router.HandleFunc("/users/{user_id}", user.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}", user.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", user.UpdateUser).Methods("UPDATE")
	router.HandleFunc("/users/{user_id}/gadgets", user.RegisterUserGadget).Methods("POST")

	// gadget routes
	router.HandleFunc("/gadgets", gadget.GetGadgets).Methods("GET")
	router.HandleFunc("/gadgets", gadget.CreateGadget).Methods("POST")
	router.HandleFunc("/gadgets/{gadget_uuid}", gadget.GetGadget).Methods("GET")
	router.HandleFunc("/gadgets/{gadget_uuid}", gadget.DeleteGadget).Methods("DELETE")
	router.HandleFunc("/gadgets/{gadget_uuid}", gadget.UpdateGadget).Methods("UPDATE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}
