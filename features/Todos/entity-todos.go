package todos

type TodosEntity struct {
	ID          uint
	UserID      uint
	TodoName    string
	Description string
}

type DataTodosInterface interface {
	Insert(todo TodosEntity) error
	Delete(id uint) error
	Update(id uint, todo TodosEntity) error
	SelectByUserId(id uint) ([]TodosEntity, error)
	SelectById(id uint) (*TodosEntity, error)
}

type ServiceTodosInterface interface {
	Create(todo TodosEntity) error
	Delete(id uint, userid uint) error
	Update(id uint, userid uint, todo TodosEntity) error
	GetById(id uint) (todo *TodosEntity, err error)
	GetByUserId(id uint) ([]TodosEntity, error)
}
