package service

import (
	todos "BE23TODO/features/Todos"
	"errors"
)

type todoService struct {
	todoData todos.DataTodosInterface
}

func New(pr todos.DataTodosInterface) todos.ServiceTodosInterface {
	return &todoService{
		todoData: pr,
	}

}

// Create implements todo.ServiceTodoInterface.
func (pr *todoService) Create(todo todos.TodosEntity) error {
	if todo.TodoName == "" {
		return errors.New("todo name cannot be empty")
	}
	err := pr.todoData.Insert(todo)
	if err != nil {
		return err
	}
	return nil
}

// Delete implements todo.ServiceTodoInterface.
func (pr *todoService) Delete(id uint, userid uint) error {
	if id <= 0 {
		return errors.New("invalid todo ID")
	}
	cekuserid, err := pr.todoData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot delete todo")
	}

	return pr.todoData.Delete(id)
}

// GetById implements todo.ServiceTodoInterface.
func (pr *todoService) GetById(id uint) (todo *todos.TodosEntity, err error) {
	if id <= 0 {
		return nil, errors.New("id not valid")
	}
	return pr.todoData.SelectById(id)
}

// Update implements todo.ServiceTodoInterface.
func (pr *todoService) Update(id uint, userid uint, todo todos.TodosEntity) error {
	if id == 0 {
		return errors.New("invalid todo ID")
	}
	if todo.TodoName == "" {
		return errors.New("todo name cannot be empty")
	}

	cekuserid, err := pr.todoData.SelectById(id)
	if err != nil {
		return err
	}

	if cekuserid.UserID != userid {
		return errors.New("user id not match, cannot update todo")
	}

	return pr.todoData.Update(id, todo)
}

// GetByUserId implements todo.ServiceTodoInterface.
func (pr *todoService) GetByUserId(id uint) ([]todos.TodosEntity, error) {
	return pr.todoData.SelectByUserId(id)
}
