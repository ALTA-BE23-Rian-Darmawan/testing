package service_test

import (
	todos "BE23TODO/features/Todos"
	"BE23TODO/features/Todos/service"
	"BE23TODO/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {
	mockTodoData := new(mocks.DataTodosInterface)
	service := service.New(mockTodoData)

	t.Run("success", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName:    "Learn Golang",
			Description: "Learn Golang basic functionality",
		}

		mockTodoData.On("Insert", todo).Return(nil).Once()

		err := service.Create(todo)
		assert.NoError(t, err)
		mockTodoData.AssertExpectations(t)
	})

	t.Run("failed", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName:    "",
			Description: "",
		}

		err := service.Create(todo)
		assert.Error(t, err)
		assert.Equal(t, "todo name cannot be empty", err.Error())
	})
}

func TestGetTodosByID(t *testing.T) {
	mockTodoData := new(mocks.DataTodosInterface)
	service := service.New(mockTodoData)

	t.Run("success", func(t *testing.T) {
		todo := &todos.TodosEntity{
			TodoName: "Learn Golang",
		}
		mockTodoData.On("SelectById", uint(1)).Return(todo, nil).Once()

		result, err := service.GetById(1)
		assert.NoError(t, err)
		assert.Equal(t, todo, result)
		mockTodoData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		_, err := service.GetById(0)
		assert.Error(t, err)
		assert.Equal(t, "id not valid", err.Error())
	})
}

func TestDeleteTodos(t *testing.T) {
	mockTodoData := new(mocks.DataTodosInterface)
	service := service.New(mockTodoData)

	t.Run("success", func(t *testing.T) {
		mockTodoData.On("SelectById", uint(1)).Return(&todos.TodosEntity{UserID: 1}, nil).Once()
		mockTodoData.On("Delete", uint(1)).Return(nil).Once()

		err := service.Delete(1, 1)
		assert.NoError(t, err)
		mockTodoData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		err := service.Delete(0, 1)
		assert.Error(t, err)
		assert.Equal(t, "invalid todo ID", err.Error())
	})

	t.Run("user ID mismatch", func(t *testing.T) {
		mockTodoData.On("SelectById", uint(1)).Return(&todos.TodosEntity{UserID: 2}, nil).Once()

		err := service.Delete(1, 1)
		assert.Error(t, err)
		assert.Equal(t, "user id not match, cannot delete todo", err.Error())
	})
}

func TestUpdateTodo(t *testing.T) {
	mockTodoData := new(mocks.DataTodosInterface)
	service := service.New(mockTodoData)

	t.Run("success", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName:    "Learn Golang",
			Description: "Learn Golang basic",
		}
		mockTodoData.On("SelectById", uint(1)).Return(&todos.TodosEntity{UserID: 1}, nil).Once()
		mockTodoData.On("Update", uint(1), todo).Return(nil).Once()

		err := service.Update(1, 1, todo)
		assert.NoError(t, err)
		mockTodoData.AssertExpectations(t)
	})

	t.Run("invalid ID", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName:    "Learn Golang",
			Description: "Learn Golang basic",
		}
		err := service.Update(0, 1, todo)
		assert.Error(t, err)
		assert.Equal(t, "invalid todo ID", err.Error())
	})

	t.Run("empty todo name", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName:    "",
			Description: "",
		}
		err := service.Update(1, 1, todo)
		assert.Error(t, err)
		assert.Equal(t, "todo name cannot be empty", err.Error())
	})

	t.Run("user ID mismatch", func(t *testing.T) {
		todo := todos.TodosEntity{
			TodoName: "Learn Golang",
		}
		mockTodoData.On("SelectById", uint(1)).Return(&todos.TodosEntity{UserID: 2}, nil).Once()

		err := service.Update(1, 1, todo)
		assert.Error(t, err)
		assert.Equal(t, "user id not match, cannot update todo", err.Error())
	})
}

func TestGetUserTodosByUserID(t *testing.T) {
	mockTodoData := new(mocks.DataTodosInterface)
	service := service.New(mockTodoData)

	t.Run("success", func(t *testing.T) {
		todosList := []todos.TodosEntity{
			{TodoName: "Learn Golang"},
			{TodoName: "Learn Docker"},
		}
		mockTodoData.On("SelectByUserId", uint(1)).Return(todosList, nil).Once()

		result, err := service.GetByUserId(1)
		assert.NoError(t, err)
		assert.Equal(t, todosList, result)
		mockTodoData.AssertExpectations(t)
	})
}
