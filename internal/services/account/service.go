package account

import (
	"database/sql"
	"errors"
	"github.com/Solar-2020/Account-Backend/pkg/models"
)

type Service interface {
	GetByID(userID int) (user models.User, err error)
	GetByEmail(email string) (user models.User, err error)
	Create(createUser models.User) (user models.User, err error)
	Edit(editUser models.User) (user models.User, err error)
	Delete(userID int) (err error)
}

type service struct {
	accountStorage accountStorage
}

func NewService(accountStorage accountStorage) Service {
	return &service{
		accountStorage: accountStorage,
	}
}

func (s *service) GetByID(userID int) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByID(userID)
	return
}

func (s *service) GetByEmail(email string) (user models.User, err error) {
	user, err = s.accountStorage.SelectUserByEmail(email)
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
