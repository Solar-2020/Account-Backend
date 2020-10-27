package accountStorage

import (
	"database/sql"
	"github.com/Solar-2020/Account-Backend/pkg/models"
)

type Storage interface {
	InsertUser(user models.User) (userID int, err error)

	UpdateUser(user models.User) (err error)
	SelectUserByID(userID int) (user models.User, err error)
	SelectUserByEmail(email string) (user models.User, err error)

	DeleteUser(userID int) (err error)
}

type storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) Storage {
	return &storage{
		db: db,
	}
}

func (s *storage) InsertUser(user models.User) (userID int, err error) {
	const sqlQuery = `
	INSERT INTO users(email, name, surname, avatar_url)
	VALUES ($1, $2, $3, $4)
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, user.Email, user.Name, user.Surname, user.AvatarURL).Scan(&userID)

	return
}

func (s *storage) UpdateUser(user models.User) (err error) {
	panic("implement me")
}

func (s *storage) SelectUserByID(userID int) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE u.id = $1
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, userID).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) SelectUserByEmail(email string) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE UPPER(u.email) = UPPER($1)
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) DeleteUser(userID int) (err error) {
	panic("implement me")
}
