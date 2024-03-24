package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/timkippdev/http-server-base-go/pkg/server"
)

func (rh *Handler) RegisterExampleRoutes(router *mux.Router) {
	rh.Get(router, "/auth", func(w http.ResponseWriter, r *http.Request) {
		server.WriteResponse(w, "authenticated", nil)
	}, &accessOptions{authRequired: true})

	rh.Get(router, "/metadata", func(w http.ResponseWriter, r *http.Request) {
		server.WriteResponse(w, []string{"meta", "data"}, server.NewMetadata(2, &server.PaginationParams{
			Limit:  10,
			Offset: 0,
		}))
	}, nil)

	rh.Get(router, "/ping", func(w http.ResponseWriter, r *http.Request) {
		server.WriteResponse(w, "pong", nil)
	}, nil)
}
