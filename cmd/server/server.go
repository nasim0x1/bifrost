package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/nasim0x1/bifrost/services/user"
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

const (
	reset   = "\033[0m"
	red     = "\033[31m"
	green   = "\033[32m"
	yellow  = "\033[33m"
	blue    = "\033[34m"
	magenta = "\033[35m"
	cyan    = "\033[36m"
	white   = "\033[37m"
)

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.statusCode = code
	rec.ResponseWriter.WriteHeader(code)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		next.ServeHTTP(rec, r)

		duration := time.Since(start)
		statusColor := red
		if rec.statusCode >= 200 && rec.statusCode < 300 {
			statusColor = green
		} else if rec.statusCode >= 300 && rec.statusCode < 400 {
			statusColor = yellow
		} else if rec.statusCode >= 400 && rec.statusCode < 500 {
			statusColor = magenta
		} else if rec.statusCode >= 500 {
			statusColor = red
		}

		log.Printf("%sMethod: %s, URL: %s, Status: %s%d%s, Duration: %v, Time: %s%s",
			statusColor, r.Method, r.URL.Path, statusColor, rec.statusCode, reset, duration, start.Format(time.RFC1123), reset)
	})
}
func (s *Server) Start(debug bool) error {
	router := mux.NewRouter()

	subRoute := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterUserRoutes(subRoute)

	router.HandleFunc("/", handleIndexEndpoint)
	log.Println("Listening on", s.Addr)
	if debug {
		logger := loggingMiddleware(router)
		return http.ListenAndServe(s.Addr, logger)
	}
	return http.ListenAndServe(s.Addr, router)
}

func handleIndexEndpoint(w http.ResponseWriter, r *http.Request) {
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
}
