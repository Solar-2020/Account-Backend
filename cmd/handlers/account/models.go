package accountHandler

import (
	"github.com/Solar-2020/Account-Backend/internal/models"
	"github.com/valyala/fasthttp"
)

type accountService interface {
	GetByID(userID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Create(user models.User) (createdUser models.User, err error)
	Edit(user models.User) (updatedUser models.User, err error)
	Delete(userID int) (err error)
}

type accountTransport interface {
	AuthorizationDecode(ctx *fasthttp.RequestCtx) (request models.Authorization, err error)
	AuthorizationEncode(ctx *fasthttp.RequestCtx, cookie models.Cookie) (err error)

	RegistrationDecode(ctx *fasthttp.RequestCtx) (request models.Registration, err error)
	RegistrationEncode(ctx *fasthttp.RequestCtx, cookie models.Cookie) (err error)

	GetUserIdByCookieDecode(ctx *fasthttp.RequestCtx) (cookieValue string, err error)
	GetUserIdByCookieEncode(ctx *fasthttp.RequestCtx, userID int) (err error)
}

type errorWorker interface {
	ServeJSONError(ctx *fasthttp.RequestCtx, serveError error) (err error)
	ServeFatalError(ctx *fasthttp.RequestCtx)
}
