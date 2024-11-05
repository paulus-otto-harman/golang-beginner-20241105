package repository

import (
	"20241105/class/2/model"
	"database/sql"
)

type Todo struct {
	Db *sql.DB
}

func InitTodoRepo(db *sql.DB) *Todo {
	return &Todo{Db: db}
}

func (repo *Todo) Get(session model.Session) ([]model.Todo, error) {
	query := `SELECT todos.id, description,completed
				FROM todos
				JOIN sessions ON todos.user_id=sessions.user_id
				WHERE sessions.id=$1 AND sessions.expired_at IS NULL AND todos.deleted_at IS NULL`

	rows, err := repo.Db.Query(query, session.Id)
	var todos []model.Todo
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&todo.Id, &todo.Description, &todo.Completed); err != nil {
			return todos, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return todos, err
	}
	return todos, nil
}

func (repo *Todo) Create(todo *model.Todo, session model.Session) error {
	query := `INSERT INTO todos (description, user_id) 
				SELECT $1,user_id
				FROM sessions
				WHERE id=$2 AND sessions.expired_at IS NULL
    			RETURNING id`

	if err := repo.Db.QueryRow(query, todo.Description, session.Id).Scan(&todo.Id); err != nil {
		return err
	}
	return nil
}

func (repo *Todo) Update(todo *model.Todo, session model.Session) error {
	query := `UPDATE todos
				SET completed=NOT completed, updated_at=NOW()
				FROM sessions
				WHERE todos.user_id=sessions.user_id AND todos.id=$1 AND sessions.id=$2 AND sessions.expired_at IS NULL
				RETURNING completed`

	if err := repo.Db.QueryRow(query, todo.Id, session.Id).Scan(&todo.Completed); err != nil {
		return err
	}
	return nil
}

func (repo *Todo) Delete(todo *model.Todo, session model.Session) error {
	query := `UPDATE todos
				SET deleted_at=NOW()
				FROM sessions
				WHERE todos.user_id=sessions.user_id AND todos.id=$1 AND sessions.id=$2 AND sessions.expired_at IS NULL AND deleted_at IS NULL
				RETURNING todos.id, todos.description`

	if err := repo.Db.QueryRow(query, todo.Id, session.Id).Scan(&todo.Id, &todo.Description); err != nil {
		return err
	}
	return nil
}
