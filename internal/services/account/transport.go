package account

import (
	"encoding/json"
	"github.com/Solar-2020/Account-Backend/pkg/models"
	"github.com/valyala/fasthttp"
	"strconv"
)

type Transport interface {
	GetByIDDecode(ctx *fasthttp.RequestCtx) (userID int, err error)
	GetByIDEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	GetByEmailDecode(ctx *fasthttp.RequestCtx) (email string, err error)
	GetByEmailEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	CreateDecode(ctx *fasthttp.RequestCtx) (createUser models.User, err error)
	CreateEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	EditDecode(ctx *fasthttp.RequestCtx) (createUser models.User, err error)
	EditEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	DeleteDecode(ctx *fasthttp.RequestCtx) (userID int, err error)
	DeleteEncode(ctx *fasthttp.RequestCtx) (err error)
}

type transport struct {
}

func NewTransport() Transport {
	return &transport{}
}

func (t transport) GetByIDDecode(ctx *fasthttp.RequestCtx) (userID int, err error) {
	userIDStr := ctx.UserValue("userID").(string)
	userID, err = strconv.Atoi(userIDStr)

	return
}

func (t transport) GetByIDEncode(ctx *fasthttp.RequestCtx, user models.User) (err error) {
	body, err := json.Marshal(user)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}

func (t transport) GetByEmailDecode(ctx *fasthttp.RequestCtx) (email string, err error) {
	email = ctx.UserValue("email").(string)
	return
}

func (t transport) GetByEmailEncode(ctx *fasthttp.RequestCtx, user models.User) (err error) {
	body, err := json.Marshal(user)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}

func (t transport) CreateDecode(ctx *fasthttp.RequestCtx) (createUser models.User, err error) {
	err = json.Unmarshal(ctx.Request.Body(), &createUser)
	return
}

func (t transport) CreateEncode(ctx *fasthttp.RequestCtx, user models.User) (err error) {
	body, err := json.Marshal(user)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}

func (t transport) EditDecode(ctx *fasthttp.RequestCtx) (editUser models.User, err error) {
	err = json.Unmarshal(ctx.Request.Body(), &editUser)
	return
}

func (t transport) EditEncode(ctx *fasthttp.RequestCtx, user models.User) (err error) {
	body, err := json.Marshal(user)
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}

func (t transport) DeleteDecode(ctx *fasthttp.RequestCtx) (userID int, err error) {
	userIDStr := ctx.UserValue("userID").(string)
	userID, err = strconv.Atoi(userIDStr)

	return
}

func (t transport) DeleteEncode(ctx *fasthttp.RequestCtx) (err error) {
	body, err := json.Marshal("OK")
	if err != nil {
		return
	}
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.Header.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBody(body)
	return
}
