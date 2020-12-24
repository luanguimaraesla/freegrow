package brain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	// auth routes
	router.HandleFunc("/signup", user.Signup).Methods("POST")
	router.HandleFunc("/signin", user.Signin).Methods("POST")

	// user routes
	router.HandleFunc("/users/{user_id}", user.GetUser).Methods("GET")
	router.HandleFunc("/users/{user_id}", user.DeleteUser).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", user.UpdateUser).Methods("UPDATE")

	// user gadget routes
	router.HandleFunc("/users/{user_id}/gadgets", user.GetUserGadgets).Methods("GET")
	router.HandleFunc("/users/{user_id}/gadgets", user.RegisterUserGadget).Methods("POST")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", user.GetUserGadget).Methods("GET")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", user.UnregisterUserGadget).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", user.UpdateUserGadget).Methods("UPDATE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}
