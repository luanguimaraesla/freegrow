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
	ID      interface{} `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
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

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
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

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
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
		ID:      int64(user.ID),
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

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	g := gadget.New()

	if err := json.NewDecoder(r.Body).Decode(&g); err != nil {
		log.L.Fatal("unable to decode the request body", zap.Error(err))
	}

	if err := user.Gadgets().Register(g); err != nil {
		log.L.Fatal("unable to register a new gadget", zap.Error(err))
	}

	msg := fmt.Sprintf("gadget %s registered successfully", g.UUID)

	res := response{
		ID:      int64(user.ID),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// UnregisterUserGadget creates an user gadget row in the database
func UnregisterUserGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(r)

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	g, err := user.Gadgets().Get(params["gadget_uuid"])
	if err != nil {
		log.L.Fatal("unable to get gadget", zap.Error(err))
	}

	if err := user.Gadgets().Unregister(g); err != nil {
		log.L.Fatal("unable to unregister gadget", zap.Error(err))
	}

	msg := fmt.Sprintf("gadget %s unregistered successfully", g.UUID)

	res := response{
		ID:      g.UUID,
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

// GetUserGadgets returns a list of user gadget in the database
func GetUserGadgets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	all, err := user.Gadgets().All()
	if err != nil {
		log.L.Fatal("unable to get all gadgets", zap.Error(err))
	}

	json.NewEncoder(w).Encode(all)
}

// GetUserGadget a single gadget of this user in the database
func GetUserGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	UUID := params["gadget_uuid"]

	gadget, err := user.Gadgets().Get(UUID)
	if err != nil {
		log.L.Fatal("unable to get gadget", zap.Error(err))
	}

	json.NewEncoder(w).Encode(gadget)
}

// UpdateUserGadget updates gadget's detail in the postgres db
func UpdateUserGadget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(params["user_id"])
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	UUID := params["gadget_uuid"]

	gadget, err := user.Gadgets().Get(UUID)
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

// loadUser is a helper function that receives an string with the
// user_id and returns an user instance loaded from db
func loadUser(userID string) (*User, error) {
	id, err := strconv.Atoi(userID)
	if err != nil {
		return nil, fmt.Errorf("unable to parse user_id into int: %v", err)
	}

	users := NewUsers()

	user, err := users.Get(int64(id))
	if err != nil {
		return nil, fmt.Errorf("unable to get user: %v", err)
	}

	return user, nil
}
