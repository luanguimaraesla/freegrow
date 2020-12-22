package gadget

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"go.uber.org/zap"
)

// response format
type response struct {
	ID      interface{} `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
}

// CreateGadget creates a gadget row in the database
func CreateGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	gadget := New()

	err := json.NewDecoder(r.Body).Decode(&gadget)
	if err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := gadget.Create(); err != nil {
		log.L.Fatal("unable to create gadget", zap.Error(err))
	}

	res := response{
		ID:      gadget.UUID,
		Message: "gadget created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

// GetGadget will return a single gadget by its ID
func GetGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)
	UUID := params["gadget_uuid"]
	gadgets := NewGadgets()

	gadget, err := gadgets.Get(UUID)
	if err != nil {
		log.L.Fatal("unable to get gadget", zap.Error(err))
	}

	json.NewEncoder(w).Encode(gadget)
}

// GetGadget will return all the gadgets
func GetGadgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	gadgets := NewGadgets()

	all, err := gadgets.All()
	if err != nil {
		log.L.Fatal("unable to get all gadgets", zap.Error(err))
	}

	json.NewEncoder(w).Encode(all)
}

// Update updates gadget's detail in the postgres db
func UpdateGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	UUID := params["gadget_uuid"]
	gadgets := NewGadgets()

	gadget, err := gadgets.Get(UUID)
	if err != nil {
		log.L.Fatal("unable to get gadget", zap.Error(err))
	}

	if err := json.NewDecoder(r.Body).Decode(&gadget); err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := gadget.Update(); err != nil {
		log.L.Fatal("unable to update gadget", zap.Error(err))
	}

	log.L.Info("gadget updated successfully", zap.String("gadget_uuid", gadget.UUID))
	msg := "gadget updated successfully"

	res := response{
		ID:      UUID,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteGadget deletes an gadget from database
func DeleteGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	UUID := params["gadget_uuid"]

	gadgets := NewGadgets()
	if err := gadgets.Delete(UUID); err != nil {
		log.L.Fatal("unable to delete gadget", zap.Error(err))
	}

	log.L.Info("gadget deleted successfully", zap.String("gadget_uuid", UUID))

	msg := fmt.Sprintf("gadget deleted successfully")

	res := response{
		ID:      UUID,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
