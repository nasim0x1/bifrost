package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
	router.HandleFunc("/forgot-password", h.handleForgotPassword).Methods("POST")

}
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request)          {}
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request)       {}
func (h *Handler) handleForgotPassword(w http.ResponseWriter, r *http.Request) {}
