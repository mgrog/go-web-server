package service

import (
	"fmt"
	"go_server/graph/model"

	"github.com/jmoiron/sqlx"
)

type Todo struct {
	DB *sqlx.DB
}

func (s *Todo) GetAll() ([]*model.Todo, error) {
	var todos []*model.Todo
	err := s.DB.Select(&todos, "SELECT * FROM todos")
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (s *Todo) Get(id int) (*model.Todo, error) {
	var todo model.Todo
	err := s.DB.Get(&todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (s *Todo) Update(input model.UpdateTodo) (*model.Todo, error) {
	var todo model.Todo
	rows, err := s.DB.NamedQuery(
		"UPDATE todos SET"+
			" text = COALESCE(:text, text)"+
			", parent_id = COALESCE(:parent_id, parent_id)"+
			", priority = COALESCE(:priority, priority)"+
			", done = COALESCE(:done, done)"+
			" WHERE id = :id RETURNING *", input)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&todo); err != nil {
			return nil, err
		}
		return &todo, nil
	}

	return nil, fmt.Errorf("failed to update todo")
}

func (s *Todo) RemoveParent(id int) (*model.Todo, error) {
	var todo model.Todo
	rows, err := s.DB.Queryx(
		"UPDATE todos SET"+
			" parent_id = NULL"+
			" WHERE id = $1 RETURNING *", id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(&todo); err != nil {
			return nil, err
		}
		return &todo, nil
	}
	return nil, fmt.Errorf("failed to update todo")
}

func (s *Todo) Create(input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{}
	rows, err := s.DB.NamedQuery("INSERT INTO todos (text, parent_id) VALUES (:text, :parent_id) RETURNING *", input)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	if rows.Next() {
		if err := rows.StructScan(todo); err != nil {
			return nil, err
		}
		return todo, nil
	}
	return nil, fmt.Errorf("failed to insert todo")
}

func (s *Todo) Delete(id int) (string, error) {
	result, err := s.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return "error deleting todo", err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "error checking delete result", err
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("todo with id %d does not exist", id)
	}
	return fmt.Sprintf("todo at %d deleted successfully", id), nil
}
