package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/timkippdev/http-server-base-go/pkg/route"
	"github.com/timkippdev/http-server-base-go/pkg/server"
	"github.com/timkippdev/http-server-base-go/pkg/util"
)

func main() {
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGTERM, syscall.SIGINT)

	allowedOrigins := util.GetEnv("ALLOWED_ORIGINS", "")
	allowedHeaders := util.GetEnv("ALLOWED_HEADERS", "")
	port := util.GetEnvInt("PORT", 8000)
	s := server.NewServer(context.Background(), port, strings.Split(allowedOrigins, ","), strings.Split(allowedHeaders, ","))

	rh := route.NewHandler(&exampleAuthChecker{})
	rh.RegisterAllRoutes(s)

	s.Start()
	<-doneChan
	s.Stop()
}

type exampleAuthChecker struct{}

func (checker *exampleAuthChecker) FindUserByIdentifier(ctx context.Context, identifier string) interface{} {
	return "user"
}

func (checker *exampleAuthChecker) ValidateAuthToken(ctx context.Context, authToken string) (map[string]interface{}, *server.Error) {
	if authToken == "invalid" {
		return nil, server.ErrorInvalidAuthToken
	}
	return nil, nil
}
