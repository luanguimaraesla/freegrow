package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"github.com/luanguimaraesla/freegrow/pkg/gadget"
	"go.uber.org/zap"
)

// response format
type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// CreateUser creates an user row in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	user := New()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := user.Create(); err != nil {
		log.L.Fatal("unable to create user", zap.Error(err))
	}

	res := response{
		ID:      user.ID,
		Message: "user created successfully",
	}

	json.NewEncoder(w).Encode(res)
}

// GetUser will return a single user by its ID
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to parse user_id into int64", zap.Error(err))
	}

	users := NewUsers()

	user, err := users.Get(int64(id))
	if err != nil {
		log.L.Fatal("unable to get user", zap.Error(err))
	}

	json.NewEncoder(w).Encode(user)
}

// GetUsers will return all the users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	users := NewUsers()

	all, err := users.All()
	if err != nil {
		log.L.Fatal("unable to get all users", zap.Error(err))
	}

	json.NewEncoder(w).Encode(all)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to parse user_id into int", zap.Error(err))
	}

	users := NewUsers()
	user, err := users.Get(int64(id))
	if err != nil {
		log.L.Fatal("unable to get user", zap.Error(err))
	}

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := user.Update(); err != nil {
		log.L.Fatal("unable to update user", zap.Error(err))
	}

	log.L.Info("user updated successfully", zap.Int64("userID", user.ID))
	msg := "user updated successfully"

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// DeleteUser deletes an user from database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to parse user_id into int", zap.Error(err))
	}

	users := NewUsers()
	if err := users.Delete(int64(id)); err != nil {
		log.L.Fatal("unable to delete user", zap.Error(err))
	}

	log.L.Info("user deleted successfully", zap.Int64("userID", int64(id)))

	msg := fmt.Sprintf("user deleted successfully")

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// RegisterUserGadget creates an user gadget row in the database
func RegisterUserGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to parse user_id into int64", zap.Error(err))
	}

	users := NewUsers()

	user, err := users.Get(int64(id))
	if err != nil {
		log.L.Fatal("unable to get user", zap.Error(err))
	}

	g := gadget.New()

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := user.Gadgets().Register(g); err != nil {
		log.L.Fatal("unable to register a new gadget", zap.Error(err))
	}

	msg := fmt.Sprint("gadget %s registered successfully", g.UUID)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}
