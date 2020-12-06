package accountHandler

import (
	"github.com/valyala/fasthttp"
)

type Handler interface {
	GetByID(ctx *fasthttp.RequestCtx)
	GetByEmail(ctx *fasthttp.RequestCtx)
	GetByCookie(ctx *fasthttp.RequestCtx)

	Create(ctx *fasthttp.RequestCtx)
	CreateAdvance(ctx *fasthttp.RequestCtx)
	GetYandex(ctx *fasthttp.RequestCtx)
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

func (h *handler) GetByID(ctx *fasthttp.RequestCtx) {
	userID, err := h.accountTransport.GetByIDDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.GetByID(userID)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.GetByIDEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) GetByEmail(ctx *fasthttp.RequestCtx) {
	email, err := h.accountTransport.GetByEmailDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.GetByEmail(email)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.GetByEmailEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) GetByCookie(ctx *fasthttp.RequestCtx) {
	userID, err := h.accountTransport.GetByCookieDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.GetByID(userID)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.GetByIDEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) Create(ctx *fasthttp.RequestCtx) {
	createUser, err := h.accountTransport.CreateDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.Create(createUser)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.CreateEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) CreateAdvance(ctx *fasthttp.RequestCtx) {
	createUser, err := h.accountTransport.CreateAdvanceDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.CreateAdvance(createUser)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.CreateAdvanceEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) GetYandex(ctx *fasthttp.RequestCtx) {
	createUser, err := h.accountTransport.GetYandexDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.GetYandex(createUser)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.GetYandexEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) Edit(ctx *fasthttp.RequestCtx) {
	editUser, err := h.accountTransport.EditDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	user, err := h.accountService.Edit(editUser)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.EditEncode(ctx, user)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}

func (h *handler) Delete(ctx *fasthttp.RequestCtx) {
	userID, err := h.accountTransport.DeleteDecode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountService.Delete(userID)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}

	err = h.accountTransport.DeleteEncode(ctx)
	if err != nil {
		h.errorWorker.ServeJSONError(ctx, err)
		return
	}
}
