package account

import (
	"database/sql"
	"fmt"
	models2 "github.com/Solar-2020/Account-Backend/internal/models"
	"github.com/Solar-2020/Account-Backend/pkg/models"
	"github.com/pkg/errors"
	"github.com/valyala/fasthttp"
)

var (
	ErrorInternalServer = errors.New("Внутренняя ошибка сервера")
)

type Service interface {
	GetByID(userID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Create(createUser models.User) (user models.User, err error)
	CreateAdvance(createUser models.UserAdvance) (user models.User, err error)
	GetYandex(userToken string) (user models.User, err error)
	Edit(editUser models.User) (user models.User, err error)
	Delete(userID int) (err error)
}

type service struct {
	accountStorage accountStorage
	yandexClient   yandexClient
	errorWorker    errorWorker
}

func NewService(accountStorage accountStorage, yandexClient yandexClient, errorWorker errorWorker) Service {
	return &service{
		accountStorage: accountStorage,
		yandexClient:   yandexClient,
		errorWorker:    errorWorker,
	}
}

func (s *service) GetByID(userID int) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.errorWorker.NewError(fasthttp.StatusBadRequest, ErrorUserNotFound, errors.Wrap(ErrorUserNotFound, fmt.Sprintf("User ID: %v", userID)))
			return
		}
		return user, s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
	}
	return
}

func (s *service) GetByEmail(email string) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			err = s.errorWorker.NewError(fasthttp.StatusBadRequest, ErrorUserNotFoundByEmail, errors.Wrap(ErrorUserNotFoundByEmail, email))
			return
		}
		return user, s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
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

	_, err = s.accountStorage.SelectUserAdvanceByEmail(createUser.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			createUser.ID, err = s.accountStorage.InsertUser(createUser)
			if err != nil {
				err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
				return
			}
			return createUser, nil
		}
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}

	createUser.ID, err = s.accountStorage.UpdateUserAdvance(createUser)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}

	return createUser, nil
}

func (s *service) CreateAdvance(createUser models.UserAdvance) (user models.User, err error) {
	//err = s.checkUniqueEmail(createUser.Email)
	//if err != nil {
	//	return
	//}

	user, err = s.accountStorage.SelectUserByEmail(createUser.Email)
	if err == nil {
		return
	}
	user.Email = createUser.Email
	user.ID, err = s.accountStorage.InsertUserAdvance(createUser)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}

	return user, nil
}

func (s *service) GetYandex(userToken string) (user models.User, err error) {
	var yandexUser models2.YandexUser
	yandexUser, err = s.yandexClient.GetUserInfo(userToken)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
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
				err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
				return
			}

			err = s.accountStorage.InsertYandexUser(user.ID, yandexUser.ID)
			if err != nil {
				err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
				return
			}
			return
		}
		return
	}

	user, err = s.accountStorage.SelectUserByID(userID)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}
	return
}

func (s *service) Edit(editUser models.User) (user models.User, err error) {
	err = s.validateUser(editUser)
	if err != nil {
		return
	}

	existUser, err := s.accountStorage.SelectUserByEmail(editUser.Email)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}

	if existUser.ID != editUser.ID {
		return user, s.errorWorker.NewError(fasthttp.StatusBadRequest, ErrorNoUniqueEmail, errors.Wrap(ErrorNoUniqueEmail, editUser.Email))
	}

	err = s.accountStorage.UpdateUser(editUser)
	if err != nil {
		err = s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
		return
	}

	return editUser, nil
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
	_, err = s.accountStorage.SelectCreatedUserByEmail(email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return s.errorWorker.NewError(fasthttp.StatusInternalServerError, nil, err)
	}

	return s.errorWorker.NewError(fasthttp.StatusBadRequest, ErrorNoUniqueEmail, errors.Wrap(ErrorNoUniqueEmail, email))
}
