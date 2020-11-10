package handlers

import (
	accountHandler "github.com/Solar-2020/Account-Backend/cmd/handlers/account"
	httputils "github.com/Solar-2020/GoUtils/http"
	"github.com/buaazp/fasthttprouter"
)

func NewFastHttpRouter(account accountHandler.Handler, middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()
	router.PanicHandler = httputils.PanicHandler

	router.Handle("GET", "/health", middleware.Log(httputils.HealthCheckHandler))

	router.Handle("GET", "/api/account/by-user/:userID", middleware.Log(middleware.ExternalAuth(account.GetByID)))
	router.Handle("GET", "/api/account/by-email/:email", middleware.Log(middleware.ExternalAuth(account.GetByEmail)))
	router.Handle("GET", "/api/account/by-cookie", middleware.Log(middleware.ExternalAuth(account.GetByCookie)))

	router.Handle("POST", "/api/account/user", middleware.Log(middleware.ExternalAuth(account.Create)))
	router.Handle("PUT", "/api/account/user", middleware.Log(middleware.ExternalAuth(account.Edit)))
	router.Handle("DELETE", "/api/account/user", middleware.Log(middleware.ExternalAuth(account.Delete)))

	router.Handle("GET", "/api/internal/account/by-user/:userID", middleware.Log(middleware.InternalAuth(account.GetByID)))
	router.Handle("GET", "/api/internal/account/by-email/:email", middleware.Log(middleware.InternalAuth(account.GetByEmail)))
	router.Handle("POST", "/api/internal/account/user", middleware.Log(middleware.InternalAuth(account.Create)))
	router.Handle("GET", "/api/internal/account/yandex/:userToken", middleware.Log(middleware.InternalAuth(account.GetYandex)))
	return router
}
