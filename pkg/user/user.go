package user

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/luanguimaraesla/freegrow/internal/global"
)

var (
	requestTimeout = 10 * time.Second
)

func Missing(w http.ResponseWriter, r *http.Request) {
	global.GlobalLogger.Info("calling missing function")

	w.Header().Set("Content-Type", "application/json")

	_, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	response := map[string]string{
		"message": "missing function",
	}

	json.NewEncoder(w).Encode(response)
}
