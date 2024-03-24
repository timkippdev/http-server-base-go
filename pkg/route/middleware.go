package route

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/timkippdev/http-server-base-go/pkg/server"
)

const (
	authHeaderKey = "Authorization"
)

type AuthChecker interface {
	FindUserByIdentifier(ctx context.Context, identifier string) interface{}
	ValidateAuthToken(ctx context.Context, authToken string) (map[string]interface{}, *server.Error)
}

func (rh *Handler) checkAuthentication(next http.HandlerFunc, authChecker AuthChecker) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authTokenFromHeader := r.Header.Get(authHeaderKey)
		fmt.Println(authTokenFromHeader)
		if authTokenFromHeader == "" {
			server.WriteErrorResponse(w, server.ErrorMissingAuthToken)
			return
		}

		authTokenParts := strings.Split(authTokenFromHeader, " ")
		if len(authTokenParts) != 2 || authTokenParts[0] != "Bearer" {
			server.WriteErrorResponse(w, server.ErrorInvalidAuthToken)
			return
		}

		ctx := r.Context()

		claims, err := authChecker.ValidateAuthToken(ctx, authTokenParts[1])
		if err != nil {
			server.WriteErrorResponse(w, err)
			return
		}

		guid := fmt.Sprintf("%v", claims["sub"])
		user := authChecker.FindUserByIdentifier(ctx, guid)
		if user == nil {
			server.WriteErrorResponse(w, server.ErrorInvalidAuthToken)
			return
		}

		ctx = context.WithValue(ctx, server.UserContextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
