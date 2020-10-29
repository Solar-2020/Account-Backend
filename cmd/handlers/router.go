package handlers

import (
	accountHandler "github.com/Solar-2020/Account-Backend/cmd/handlers/account"
	httputils "github.com/Solar-2020/GoUtils/http"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(account accountHandler.Handler, middleware httputils.Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.PanicHandler =  httputils.PanicHandler

	router.Handle("GET", "/health", middleware.Log(httputils.HealthCheckHandler))
	clientside := httputils.ClientsideChain(middleware)

	router.Handle("GET", "/api/account/by-user/:userID", clientside(account.GetByID))
	router.Handle("GET", "/api/account/by-email/:email", clientside(account.GetByEmail))

	router.Handle("POST", "/api/account/user", clientside(account.Create))
	router.Handle("PUT", "/api/account/user", clientside(account.Edit))
	router.Handle("DELETE", "/api/account/user", clientside(account.Delete))

	serverside := httputils.ServersideChain(middleware)
	router.Handle("GET", "/api/internal/account/by-user/:userID", serverside(account.GetByID))
	router.Handle("GET", "/api/internal/account/by-email/:email", serverside(account.GetByEmail))
	router.Handle("POST", "/api/internal/account/user", serverside(account.Create))
	return router
}
