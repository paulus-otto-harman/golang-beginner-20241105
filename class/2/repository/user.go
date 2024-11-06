package repository

import (
	"20241105/class/2/model"
	"database/sql"
)

type User struct {
	Db *sql.DB
}

func InitUserRepo(db *sql.DB) *User {
	return &User{Db: db}
}

func (repo *User) Create(user *model.User) (*model.Session, error) {
	query := `INSERT INTO users (name, username, password) VALUES ($1, $2, $3) RETURNING id`

	tx, err := repo.Db.Begin()
	if err != nil {
		return nil, err
	}

	if err = repo.Db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&user.Id); err != nil {
		return nil, err
	}

	var session model.Session

	query = `INSERT INTO sessions (user_id) VALUES ($1) RETURNING id`
	if err = repo.Db.QueryRow(query, user.Id).Scan(&session.Id); err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	return &session, nil
}

func (repo *User) All(session model.Session) ([]model.User, error) {
	query := `SELECT users.id, users.name, users.username, users.is_active
				FROM users
				JOIN sessions ON users.id=sessions.user_id
				WHERE sessions.id != $1 AND sessions.expired_at IS NULL `

	rows, err := repo.Db.Query(query, session.Id)
	var users []model.User
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Username, &user.IsActive); err != nil {
			return []model.User{}, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return []model.User{}, err
	}
	return users, nil
}

func (repo *User) Get(session *model.Session) error {
	query := `SELECT sessions.id, users.name, users.username, users.password
				FROM sessions
				JOIN users ON sessions.user_id = users.id WHERE sessions.user_id = $1`

	if err := repo.Db.QueryRow(query, session.User.Id).Scan(&session.Id, &session.User.Name, &session.User.Username, &session.User.Password); err != nil {
		return err
	}
	return nil
}
