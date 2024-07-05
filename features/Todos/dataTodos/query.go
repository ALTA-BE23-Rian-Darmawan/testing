package datatodos

import (
	todos "BE23TODO/features/Todos"

	"gorm.io/gorm"
)

type todoQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) todos.DataTodosInterface {
	return &todoQuery{
		db: db,
	}
}

// Delete implements todo.DataTodosInterface.
func (t *todoQuery) Delete(id uint) error {
	tx := t.db.Delete(&Todos{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Insert implements todo.DataTodosInterface.
func (t *todoQuery) Insert(todo todos.TodosEntity) error {
	//var todoGorm Project
	todoGorm := Todos{
		UserID:      todo.UserID,
		TodoName:    todo.TodoName,
		Description: todo.Description,
	}
	tx := t.db.Create(&todoGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Update implements todo.DataTodosInterface.
func (t *todoQuery) Update(id uint, todo todos.TodosEntity) error {
	var todoGorm Todos
	tx := t.db.First(&todoGorm, id)
	if tx.Error != nil {
		return tx.Error
	}

	todoGorm.TodoName = todo.TodoName
	todoGorm.Description = todo.Description

	tx = t.db.Save(&todoGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (t *todoQuery) SelectByUserId(id uint) ([]todos.TodosEntity, error) {
	var allTodos []Todos // var penampung data yg dibaca dari db
	tx := t.db.Where("user_id =?", id).Find(&allTodos)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allTodosCore []todos.TodosEntity
	for _, v := range allTodos {
		allTodosCore = append(allTodosCore, todos.TodosEntity{
			ID:          v.ID,
			UserID:      v.UserID,
			TodoName:    v.TodoName,
			Description: v.Description,
		})
	}

	return allTodosCore, nil
}

// SelectById implements todo.DataTodosInterface.
func (t *todoQuery) SelectById(id uint) (*todos.TodosEntity, error) {
	var todoGorm Todos
	tx := t.db.First(&todoGorm, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping
	var projectcore = todos.TodosEntity{
		ID:          todoGorm.ID,
		UserID:      todoGorm.UserID,
		TodoName:    todoGorm.TodoName,
		Description: todoGorm.Description,
	}

	return &projectcore, nil
}
