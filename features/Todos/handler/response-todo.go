package handler

type TodoResponse struct {
	ID          uint   `json:"id"`
	UserID      uint   `json:"user_id"`
	TodoName    string `json:"todo_name"`
	Description string `json:"Description"`
}
