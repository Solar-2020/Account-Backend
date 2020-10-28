package accountHandler

import (
	"github.com/Solar-2020/GoUtils/context"
)

type Handler interface {
	GetByID(ctx context.Context)
	GetByEmail(ctx context.Context)
	Create(ctx context.Context)
	Edit(ctx context.Context)
	Delete(ctx context.Context)
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

func (h *handler) GetByID(ctx context.Context) {
	userID, err := h.accountTransport.GetByIDDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	user, err := h.accountService.GetByID(userID)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.accountTransport.GetByIDEncode(ctx.RequestCtx, user)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) GetByEmail(ctx context.Context) {
	email, err := h.accountTransport.GetByEmailDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	user, err := h.accountService.GetByEmail(email)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.accountTransport.GetByEmailEncode(ctx.RequestCtx, user)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) Create(ctx context.Context) {
	createUser, err := h.accountTransport.CreateDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	user, err := h.accountService.Create(createUser)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.accountTransport.CreateEncode(ctx.RequestCtx, user)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) Edit(ctx context.Context) {
	editUser, err := h.accountTransport.EditDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
	if editUser.ID != ctx.Session.Uid {
		editUser.ID = ctx.Session.Uid
	}


	user, err := h.accountService.Edit(editUser)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.accountTransport.EditEncode(ctx.RequestCtx, user)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) Delete(ctx context.Context) {
	userID, err := h.accountTransport.DeleteDecode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	if userID != ctx.Session.Uid {
		userID = ctx.Session.Uid
	}

	err = h.accountService.Delete(userID)
	if err != nil {
		h.handleError(err, ctx)
		return
	}

	err = h.accountTransport.DeleteEncode(ctx.RequestCtx)
	if err != nil {
		h.handleError(err, ctx)
		return
	}
}

func (h *handler) handleError(err error, ctx context.Context) {
	err = h.errorWorker.ServeJSONError(ctx.RequestCtx, err)
	if err != nil {
		h.errorWorker.ServeFatalError(ctx.RequestCtx)
	}
	return
}