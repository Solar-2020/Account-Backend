package accountHandler

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

type Handler interface {
	Get(ctx *fasthttp.RequestCtx)
	Create(ctx *fasthttp.RequestCtx)
	Edit(ctx *fasthttp.RequestCtx)
	Delete(ctx *fasthttp.RequestCtx)
}

type handler struct {
	accountService   accountService
	accountTransport accountTransport
	errorWorker      errorWorker
}

func NewHandler(accountService accountService, accountTransport accountTransport, errorWorker errorWorker) Handler {
	return &handler{
		accountService:   accountService,
		accountTransport: accountTransport,
		errorWorker:      errorWorker,
	}
}

func (h *handler) Authorization(ctx *fasthttp.RequestCtx) {
	fmt.Println("New incoming request: POST /account/account")
	auth, err := h.authorizationTransport.AuthorizationDecode(ctx)
	if err != nil {
		fmt.Println("Create: cannot decode request")
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	cookie, err := h.authorizationService.Authorization(auth)
	if err != nil {
		fmt.Println("Create: bad usecase: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.authorizationTransport.AuthorizationEncode(ctx, cookie)
	if err != nil {
		fmt.Println("Create: cannot encode response: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}

func (h *handler) Registration(ctx *fasthttp.RequestCtx) {
	fmt.Println("New incoming request: POST /account/account")
	auth, err := h.authorizationTransport.RegistrationDecode(ctx)
	if err != nil {
		fmt.Println("Create: cannot decode request")
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	cookie, err := h.authorizationService.Registration(auth)
	if err != nil {
		fmt.Println("Create: bad usecase: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.authorizationTransport.RegistrationEncode(ctx, cookie)
	if err != nil {
		fmt.Println("Create: cannot encode response: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}

func (h *handler) GetUserIdByCookie(ctx *fasthttp.RequestCtx) {
	fmt.Println("New incoming request: POST /account/account")
	cookieValue, err := h.authorizationTransport.GetUserIdByCookieDecode(ctx)
	if err != nil {
		fmt.Println("Create: cannot decode request")
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	userID, err := h.authorizationService.GetUserIdByCookie(cookieValue)
	if err != nil {
		fmt.Println("Create: bad usecase: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}

	err = h.authorizationTransport.GetUserIdByCookieEncode(ctx, userID)
	if err != nil {
		fmt.Println("Create: cannot encode response: ", err)
		err = h.errorWorker.ServeJSONError(ctx, err)
		if err != nil {
			h.errorWorker.ServeFatalError(ctx)
		}
		return
	}
}
