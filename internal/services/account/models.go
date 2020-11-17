package account

import (
	yandex "github.com/Solar-2020/Account-Backend/internal/models"
	"github.com/Solar-2020/Account-Backend/pkg/models"
	"github.com/pkg/errors"
)

var (
	ErrorUserNotFound        = errors.New("Пользователь не найден!")
	ErrorUserNotFoundByEmail = errors.New("Пользователь с указанным email не найден!")
	ErrorNoUniqueEmail       = errors.New("Аккаунт с указанным email уже существует!")
)

type accountStorage interface {
	InsertUser(user models.User) (userID int, err error)
	InsertYandexUser(userID int, yandexID string) (err error)

	UpdateUser(user models.User) (err error)
	SelectUserByID(userID int) (user models.User, err error)
	SelectUserByEmail(email string) (user models.User, err error)
	SelectUserIDByYandexID(yandexID string) (userID int, err error)

	DeleteUser(userID int) (err error)
}

type yandexClient interface {
	GetUserInfo(userToken string) (user yandex.YandexUser, err error)
}

type errorWorker interface {
	NewError(httpCode int, responseError error, fullError error) (err error)
}
