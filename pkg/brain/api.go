package brain

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/global"
	"go.uber.org/zap"
)

func (b *Brain) Listen(bind string) error {
	router := mux.NewRouter()

	router.HandleFunc("/users", missing).Methods("GET")
	router.HandleFunc("/users", missing).Methods("POST")
	router.HandleFunc("/users/{id}", missing).Methods("GET")
	router.HandleFunc("/users/{id}", missing).Methods("DELETE")

	b.L.With(zap.String("bind", bind)).Info("listening")

	if err := http.ListenAndServe(bind, router); err != nil {
		return err
	}

	return nil
}

func missing(w http.ResponseWriter, r *http.Request) {
	global.GlobalLogger.Info("calling missing function")

	w.Header().Set("Content-Type", "application/json")

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response := map[string]string{
		"message": "missing function",
	}

	json.NewEncoder(w).Encode(response)
}
