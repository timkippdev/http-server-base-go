package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timkippdev/http-server-base-go/pkg/server"
)

type Handler struct {
	authChecker AuthChecker
}

func NewHandler(authChecker AuthChecker) *Handler {
	return &Handler{
		authChecker: authChecker,
	}
}

func (rh *Handler) RegisterAllRoutes(s *server.Server) {
	router := s.GetRouter().PathPrefix("/api/v1").Subrouter()
	rh.RegisterExampleRoutes(router)
}

func (rh *Handler) Delete(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), opts *accessOptions) *mux.Route {
	return rh.register(router, path, f, opts, []string{http.MethodDelete, http.MethodOptions})
}

func (rh *Handler) Get(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), opts *accessOptions) *mux.Route {
	return rh.register(router, path, f, opts, []string{http.MethodGet})
}

func (rh *Handler) Post(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), opts *accessOptions) *mux.Route {
	return rh.register(router, path, f, opts, []string{http.MethodPost, http.MethodOptions})
}

func (rh *Handler) Put(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), opts *accessOptions) *mux.Route {
	return rh.register(router, path, f, opts, []string{http.MethodPut, http.MethodOptions})
}

func (rh *Handler) register(router *mux.Router, path string, f func(w http.ResponseWriter, r *http.Request), opts *accessOptions, methods []string) *mux.Route {
	if opts == nil {
		opts = &accessOptions{}
	}

	if opts.authRequired {
		// authenticate user
		f = rh.checkAuthentication(f, rh.authChecker)
	}

	return router.HandleFunc(path, f).Methods(methods...)
}
