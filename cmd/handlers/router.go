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

	router.Handle("POST", "/account/account", middleware.CORS(authorization.Authorization))
	router.Handle("GET", "/account/registration", middleware.CORS(authorization.Registration))

	router.Handle("GET", "/account/user-id", middleware.CORS(authorization.GetUserIdByCookie))



	return router
}

func panicHandler(ctx *fasthttp.RequestCtx, err interface{}) {
	fmt.Printf("Request falied with panic: %s, error: %v\nTrace:\n", string(ctx.Request.RequestURI()), err)
	fmt.Println(string(debug.Stack()))
	errorWorker.NewErrorWorker().ServeFatalError(ctx)
}
