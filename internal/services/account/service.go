package account

import (
	"database/sql"
	"errors"
	"github.com/Solar-2020/Account-Backend/internal/clients/yandex"
	models2 "github.com/Solar-2020/Account-Backend/internal/models"
	"github.com/Solar-2020/Account-Backend/pkg/models"
)

var (
	ErrorUserNotFound = errors.New("Пользователь не найден")
	ErrorInternalServer = errors.New("Внутренняя ошибка сервера")
)

type Service interface {
	GetByID(userID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Create(createUser models.User) (user models.User, err error)
	GetYandex(userToken string) (user models.User, err error)
	Edit(editUser models.User) (user models.User, err error)
	Delete(userID int) (err error)
}

type service struct {
	accountStorage accountStorage
	yandexClient   yandex.Client
}

func NewService(accountStorage accountStorage, yandexClient yandex.Client) Service {
	return &service{
		accountStorage: accountStorage,
		yandexClient:   yandexClient,
	}
}

func (s *service) GetByID(userID int) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrorUserNotFound
		}
		return user, ErrorInternalServer
	}
	return
}

func (s *service) GetByEmail(email string) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, ErrorUserNotFound
		}
		return user, ErrorInternalServer
	}
	return
}

func (s *service) Create(createUser models.User) (user models.User, err error) {
	err = s.validateUser(createUser)
	if err != nil {
		return
	}

	err = s.checkUniqueEmail(createUser.Email)
	if err != nil {
		return
	}

	createUser.ID, err = s.accountStorage.InsertUser(createUser)
	if err != nil {
		return
	}

	return createUser, nil
}

func (s *service) GetYandex(userToken string) (user models.User, err error) {
	var yandexUser models2.YandexUser
	yandexUser, err = s.yandexClient.GetUserInfo(userToken)
	if err != nil {
		return
	}

	userID, err := s.accountStorage.SelectUserIDByYandexID(yandexUser.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			user.Name = yandexUser.FirstName
			user.Surname = yandexUser.LastName
			user.Email = yandexUser.DefaultEmail
			if yandexUser.IsAvatarEmpty == false {
				user.AvatarURL = yandexUser.DefaultAvatarID
			}

			user, err = s.Create(user)
			if err != nil {
				return
			}

			err = s.accountStorage.InsertYandexUser(user.ID, yandexUser.ID)

			return
		}
		return
	}

	user, err = s.accountStorage.SelectUserByID(userID)

	return
}

func (s *service) Edit(editUser models.User) (user models.User, err error) {
	err = s.validateUser(editUser)
	if err != nil {
		return
	}

	err = s.checkUniqueEmail(editUser.Email)
	if err != nil {
		return
	}

	err = s.accountStorage.UpdateUser(editUser)

	return
}

func (s *service) Delete(userID int) (err error) {
	err = s.accountStorage.DeleteUser(userID)
	return
}

func (s *service) validateUser(user models.User) (err error) {
	//TODO VALIDATION
	return
}

func (s *service) checkUniqueEmail(email string) (err error) {
	_, err = s.accountStorage.SelectUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	return errors.New("Аккаунт с указанным email уже существует")
}
