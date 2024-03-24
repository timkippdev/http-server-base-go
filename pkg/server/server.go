package server

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/rs/cors"
)

type ContextKey string

const (
	UserContextKey ContextKey = "user"
)

type Server struct {
	*http.Server
	ctx  context.Context
	mux  *mux.Router
	port int
}

func NewServer(ctx context.Context, port int, allowedOrigins []string, allowedHeaders []string) *Server {
	mux := mux.NewRouter()
	mux.NotFoundHandler = notFoundHandler()
	c := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowCredentials: true,
		AllowedHeaders:   allowedHeaders,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut},
	})

	return &Server{
		&http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: loggingHandler(c.Handler(mux)),
		},
		ctx,
		mux,
		port,
	}
}

func (s *Server) GetRouter() *mux.Router {
	return s.mux
}

func (s *Server) Start() {
	log.Println("HTTP server starting")
	go func() {
		http.Handle("/", s.mux)
		s.ListenAndServe()
	}()
	log.Printf("HTTP server listening on port %d\n", s.port)
}

func (s *Server) Stop() {
	defer log.Println("HTTP server stopped")
	s.Shutdown(s.ctx)
}

func WriteErrorResponse(w http.ResponseWriter, errorResponse ErrorInterface) {
	writeResponse(w, response{Error: errorResponse}, errorResponse.GetStatus())
}

func WriteResponse(w http.ResponseWriter, data interface{}, metadata *Metadata, statusCode ...int) {
	responseStatusCode := http.StatusOK
	if len(statusCode) > 0 {
		responseStatusCode = statusCode[0]
	}
	writeResponse(w, response{Data: data, Metadata: metadata}, responseStatusCode)
}

func notFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		WriteErrorResponse(w, NewError("The path you requested was not found.", "PATH_NOT_FOUND", http.StatusNotFound))
	})
}

func loggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w}

		start := time.Now()
		next.ServeHTTP(rw, r)

		fullURL := r.URL.EscapedPath()
		if r.URL.RawQuery != "" {
			fullURL = fmt.Sprintf("%s?%s", fullURL, r.URL.RawQuery)
		}
		log.Printf("%s %s %d %s\n", r.Method, fullURL, rw.Status(), time.Since(start))
	})
}

func writeResponse(w http.ResponseWriter, response interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		w.WriteHeader(statusCode)
		json.NewEncoder(w).Encode(errors.Wrap(err, ErrorGenericError.Message))
	}
}
