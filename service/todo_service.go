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

func (s *Todo) Update(id int, input model.UpdateTodo) (*model.Todo, error) {
	var todo model.Todo
	rows, err := s.DB.Queryx(
		"UPDATE todos SET"+
			" text = COALESCE($1, text)"+
			", parent_id = COALESCE($2, parent_id)"+
			", priority = COALESCE($3, priority)"+
			", done = COALESCE($4, done)"+
			" WHERE id = $5 RETURNING *", input.Text, input.ParentID, input.Priority, input.Done, id)

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
	rows, err := s.DB.NamedQuery("INSERT INTO todos (text) VALUES (:text) RETURNING *", input)
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
