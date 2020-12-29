package user

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/luanguimaraesla/freegrow/internal/log"
	"github.com/luanguimaraesla/freegrow/internal/session"
	"github.com/luanguimaraesla/freegrow/pkg/gadget"
	"go.uber.org/zap"
)

// response format
type response struct {
	ID      interface{} `json:"id,omitempty"`
	Message string      `json:"message,omitempty"`
}

// GetUser will return a single user by its ID
func GetUser(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	user, err := loadUser(userID)
	if err != nil {
		log.L.Fatal("unable to load user", zap.Error(err))
	}

	json.NewEncoder(w).Encode(user)
}

// UpdateUser update user's detail in the postgres db
func UpdateUser(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	user, err := loadUser(userID)
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
func DeleteUser(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	id, err := strconv.Atoi(userID)
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
func RegisterUserGadget(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	user, err := loadUser(userID)
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
func UnregisterUserGadget(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(userID)
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
func GetUserGadgets(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	user, err := loadUser(userID)
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
func GetUserGadget(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(userID)
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
func UpdateUserGadget(userID string, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	user, err := loadUser(userID)
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

// Signin is used to created users get access to the system
func Signin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	creds := struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}{}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.L.Error("unable to decode the request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	users := NewUsers()

	user, err := users.Search(creds.Identity)
	if err != nil {
		log.L.Error("unable to find user", zap.Error(err))
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.checkPassword(creds.Password) {
		log.L.Error("unable to login: wrong password")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := session.CreateSession(user.ID)
	if err != nil {
		log.L.Error("unable generate session token", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := struct {
		AccessToken  string
		RefreshToken string
	}{accessToken, refreshToken}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.L.Error("failed encoding response", zap.Error(err))
	}
}

// Signout is used to logout users, invalidating its JWT token
func Signout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	res := struct {
		Message string
	}{
		Message: "successfully logged out",
	}

	json.NewEncoder(w).Encode(res)
}

// Signup is used to register new users
func Signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")

	user := New()

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.L.Error("unable to decode the request body", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if err := user.Create(); err != nil {
		log.L.Error("unable to create user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res := response{
		ID:      user.ID,
		Message: "user created successfully",
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
