package brain

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/session"
	"github.com/luanguimaraesla/freegrow/pkg/user"
	"go.uber.org/zap"
)

var (
	originsOk = handlers.AllowedOrigins([]string{"*"})
	methodsOk = handlers.AllowedMethods(
		[]string{"GET", "HEAD", "POST", "UPDATE", "DELETE", "PUT", "OPTIONS"},
	)
	headersOk = handlers.AllowedHeaders(
		[]string{
			"Accept",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
	)
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

	r := handlers.CORS(originsOk, headersOk, methodsOk)(router)
	if err := http.ListenAndServe(bind, r); err != nil {
		return err
	}

	return nil
}
