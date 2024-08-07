package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Server struct {
	Addr string
	DB   *sql.DB
	HOST string
	PORT string
}

func NewServer(addr string, host string, port string, db *sql.DB) *Server {
	return &Server{
		Addr: addr,
		DB:   db,
		HOST: host,
		PORT: port,
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Method: %s, URL: %s, Time: %s", r.Method, r.URL.Path, start.Format(time.RFC1123))
		next.ServeHTTP(w, r)
	})
}

func (s *Server) Start() error {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := map[string]any{
			"status":      "success",
			"message":     "Server is running",
			"server_time": time.Now().Format("2006-01-02 02 Jan 2006 03:04PM"),
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	loggedRouter := loggingMiddleware(router)

	log.Println("Listening on", s.Addr)
	return http.ListenAndServe(s.Addr, loggedRouter)
}
