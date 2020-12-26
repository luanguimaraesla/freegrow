package brain

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/session"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	// auth routes
	router.HandleFunc("/signup", user.Signup).Methods("POST")
	router.HandleFunc("/signin", user.Signin).Methods("POST")

	// user routes
	router.HandleFunc("/users/{user_id}", session.CheckSession(user.GetUser)).Methods("GET")
	router.HandleFunc("/users/{user_id}", session.CheckSession(user.DeleteUser)).Methods("DELETE")
	router.HandleFunc("/users/{user_id}", session.CheckSession(user.UpdateUser)).Methods("UPDATE")

	// user gadget routes
	router.HandleFunc("/users/{user_id}/gadgets", session.CheckSession(user.GetUserGadgets)).Methods("GET")
	router.HandleFunc("/users/{user_id}/gadgets", session.CheckSession(user.RegisterUserGadget)).Methods("POST")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", session.CheckSession(user.GetUserGadget)).Methods("GET")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", session.CheckSession(user.UnregisterUserGadget)).Methods("DELETE")
	router.HandleFunc("/users/{user_id}/gadgets/{gadget_uuid}", session.CheckSession(user.UpdateUserGadget)).Methods("UPDATE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}
