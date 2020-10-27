package handlers

import (
	"fmt"
	accountHandler "github.com/Solar-2020/Account-Backend/cmd/handlers/account"
	"github.com/Solar-2020/Account-Backend/internal/errorWorker"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"runtime/debug"
)

func NewFastHttpRouter(account accountHandler.Handler, middleware Middleware) *fasthttprouter.Router {
	router := fasthttprouter.New()

	//router.Handle("GET", "/health", check)

	router.PanicHandler = panicHandler

	router.Handle("GET", "/account/by-user/:userID", middleware.CORS(account.GetByID))
	router.Handle("GET", "/account/by-email/:email", middleware.CORS(account.GetByEmail))

	router.Handle("POST", "/account/user", middleware.CORS(account.Create))
	router.Handle("PUT", "/account/user", middleware.CORS(account.Edit))
	router.Handle("DELETE", "/account/user", middleware.CORS(account.Delete))

	return router
}

func panicHandler(ctx *fasthttp.RequestCtx, err interface{}) {
	fmt.Printf("Request falied with panic: %s, error: %v\nTrace:\n", string(ctx.Request.RequestURI()), err)
	fmt.Println(string(debug.Stack()))
	errorWorker.NewErrorWorker().ServeFatalError(ctx)
}
