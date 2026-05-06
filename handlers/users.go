package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	"go-backend-starter/models"
)

// UserHandler keeps user HTTP handlers together with their backing store.
// The in-memory store is intentionally small and can later be replaced by a
// database-backed repository without changing the routing layer.
type UserHandler struct {
	mu     sync.Mutex
	nextID int
	users  []models.User
}

// NewUserHandler creates a user handler with sample data.
func NewUserHandler() *UserHandler {
	return &UserHandler{
		nextID: 3,
		users: []models.User{
			{ID: 1, Name: "Alice", Email: "alice@example.com"},
			{ID: 2, Name: "Bob", Email: "bob@example.com"},
		},
	}
}

// List returns all known users.
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	users := make([]models.User, len(h.users))
	copy(users, h.users)

	writeJSON(w, http.StatusOK, users)
}

// Create adds a user from the request body.
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var input models.User
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeError(w, http.StatusBadRequest, "invalid JSON request body")
		return
	}

	if input.Name == "" || input.Email == "" {
		writeError(w, http.StatusBadRequest, "name and email are required")
		return
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	input.ID = h.nextID
	h.nextID++
	h.users = append(h.users, input)

	writeJSON(w, http.StatusCreated, input)
}
