package accountHandler

import (
	"github.com/Solar-2020/Account-Backend/pkg/models"
	"github.com/valyala/fasthttp"
)

type accountService interface {
	GetByID(userID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Create(createUser models.User) (user models.User, err error)
	CreateAdvance(createUser models.UserAdvance) (user models.User, err error)
	GetYandex(userToken string) (user models.User, err error)
	Edit(editUser models.User) (user models.User, err error)
	Delete(userID int) (err error)
}

type accountTransport interface {
	GetByIDDecode(ctx *fasthttp.RequestCtx) (userID int, err error)
	GetByIDEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	GetByCookieDecode(ctx *fasthttp.RequestCtx) (userID int, err error)

	GetByEmailDecode(ctx *fasthttp.RequestCtx) (email string, err error)
	GetByEmailEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	CreateDecode(ctx *fasthttp.RequestCtx) (createUser models.User, err error)
	CreateEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	CreateAdvanceDecode(ctx *fasthttp.RequestCtx) (createUserAdvance models.UserAdvance, err error)
	CreateAdvanceEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	GetYandexDecode(ctx *fasthttp.RequestCtx) (userToken string, err error)
	GetYandexEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	EditDecode(ctx *fasthttp.RequestCtx) (createUser models.User, err error)
	EditEncode(ctx *fasthttp.RequestCtx, user models.User) (err error)

	DeleteDecode(ctx *fasthttp.RequestCtx) (userID int, err error)
	DeleteEncode(ctx *fasthttp.RequestCtx) (err error)
}

type errorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error)
}
