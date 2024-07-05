package handler

type TodoRequest struct {
	TodoName    string `json:"todo_name"`
	Description string `json:"description"`
}
