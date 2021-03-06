package accountStorage

import (
	"database/sql"
	"github.com/Solar-2020/Account-Backend/pkg/models"
)

type Storage interface {
	InsertUser(user models.User) (userID int, err error)
	InsertUserAdvance(user models.UserAdvance) (userID int, err error)
	InsertYandexUser(userID int, yandexID string) (err error)

	UpdateUser(user models.User) (err error)
	UpdateUserAdvance(user models.User) (userID int, err error)
	SelectUserByID(userID int) (user models.User, err error)
	SelectUserByEmail(email string) (user models.User, err error)
	SelectUserAdvanceByEmail(email string) (user models.User, err error)
	SelectCreatedUserByEmail(email string) (user models.User, err error)
	SelectUserIDByYandexID(yandexID string) (userID int, err error)

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
	RETURNING id`

	err = s.db.QueryRow(sqlQuery, user.Email, user.Name, user.Surname, user.AvatarURL).Scan(&userID)

	return
}

func (s *storage) InsertUserAdvance(user models.UserAdvance) (userID int, err error) {
	const sqlQuery = `
	INSERT INTO users(email, name, surname, avatar_url, status)
	VALUES ($1, $2, $3, $4, 3)
	RETURNING id`

	err = s.db.QueryRow(sqlQuery, user.Email, "", "", "").Scan(&userID)

	return
}

func (s *storage) InsertYandexUser(userID int, yandexID string) (err error) {
	const sqlQuery = `
	INSERT INTO users_yandex(user_id, yandex_id)
	VALUES ($1, $2);`

	_, err = s.db.Exec(sqlQuery, userID, yandexID)

	return
}

func (s *storage) UpdateUser(user models.User) (err error) {
	const sqlQuery = `
	UPDATE users
	SET email = $1,
		name = $2,
		surname = $3,
		avatar_url = $4
	WHERE id = $5 AND status = 1;`

	_, err = s.db.Exec(sqlQuery, user.Email, user.Name, user.Surname, user.AvatarURL, user.ID)

	return
}

func (s *storage) UpdateUserAdvance(user models.User) (userID int, err error) {
	const sqlQuery = `
	UPDATE users
	SET name = $1,
		surname = $2,
		avatar_url = $3,
		status = 1
	WHERE email = $4 AND status = 3
	RETURNING id;`

	err = s.db.QueryRow(sqlQuery, user.Name, user.Surname, user.AvatarURL, user.Email).Scan(&userID)

	return
}

func (s *storage) SelectUserByID(userID int) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE u.id = $1 AND status IN (1,3);`

	err = s.db.QueryRow(sqlQuery, userID).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) SelectUserByEmail(email string) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE UPPER(u.email) = UPPER($1) AND status IN (1,3);`

	err = s.db.QueryRow(sqlQuery, email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) SelectCreatedUserByEmail(email string) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE UPPER(u.email) = UPPER($1) AND status = 1;`

	err = s.db.QueryRow(sqlQuery, email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) SelectUserAdvanceByEmail(email string) (user models.User, err error) {
	const sqlQuery = `
	SELECT u.id, u.email, u.name, u.surname, u.avatar_url
	FROM users as u
	WHERE UPPER(u.email) = UPPER($1) AND status = 3;`

	err = s.db.QueryRow(sqlQuery, email).Scan(&user.ID, &user.Email, &user.Name, &user.Surname, &user.AvatarURL)

	return
}

func (s *storage) SelectUserIDByYandexID(yandexID string) (userID int, err error) {
	const sqlQuery = `
	SELECT uy.user_id
	FROM users_yandex AS uy
	WHERE uy.yandex_id = $1;`
	err = s.db.QueryRow(sqlQuery, yandexID).Scan(&userID)

	return
}

func (s *storage) DeleteUser(userID int) (err error) {
	const sqlQuery = `
	UPDATE users
	SET status = 2
	WHERE id = $1;`

	_, err = s.db.Exec(sqlQuery, userID)

	return
}
