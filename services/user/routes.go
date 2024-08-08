package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/lib/pq"
	"github.com/nasim0x1/bifrost/handlers"
	"github.com/nasim0x1/bifrost/models"
	"github.com/nasim0x1/bifrost/utils"
)

type Handler struct {
	DB *sql.DB
}
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type LoginResponse struct {
	Message string               `json:"message"`
	User    *models.UserResponse `json:"user,omitempty"`
}

func NewHandler(DB *sql.DB) *Handler {
	return &Handler{
		DB: DB,
	}
}

func (h *Handler) RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/forgot-password", h.handleForgotPassword).Methods("POST")
	router.HandleFunc("/protected", h.testProtected).Methods("GET")

}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user, err := getUserByEmail(h.DB, login.Email)
	if err != nil {
		handlers.SendErrorResponse(w, http.StatusNotFound, "User not found")
		return
	}
	if !utils.VerifyPasswordHash(user.Password, login.Password) {
		handlers.SendErrorResponse(w, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.CreateJwtToken(user.ID)
	if err != nil {
		handlers.SendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	responseUser := models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		JwtToken:  token,
	}
	loginResp := LoginResponse{
		Message: "Login successful",
		User:    &responseUser,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(loginResp)
}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user.CreatedAt = time.Now()

	hashedPassword, err := utils.GenaratePasswordHash(user.Password)
	if err != nil {
		handlers.SendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	err = h.DB.QueryRow(
		`INSERT INTO users ("firstName", "lastName", "email", "password", "createdAt") VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		user.FirstName, user.LastName, user.Email, hashedPassword, user.CreatedAt).Scan(&user.ID)

	if err != nil {
		var pqErr *pq.Error
		if ok := errors.As(err, &pqErr); ok {
			if pqErr.Code == "23505" && strings.Contains(pqErr.Message, "users_email_key") {
				handlers.SendErrorResponse(w, http.StatusConflict, "Email already exists")
				return
			}
		}
		handlers.SendErrorResponse(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.UserResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	})
}
func (h *Handler) handleForgotPassword(w http.ResponseWriter, r *http.Request) {}

func getUserByEmail(DB *sql.DB, email string) (models.User, error) {
	var user models.User
	err := DB.QueryRow(`SELECT * FROM users WHERE email = $1`, email).Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (h *Handler) testProtected(w http.ResponseWriter, r *http.Request) {
	
}
