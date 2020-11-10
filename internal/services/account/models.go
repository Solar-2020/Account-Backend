package account

import (
	"github.com/Solar-2020/Account-Backend/pkg/models"
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
