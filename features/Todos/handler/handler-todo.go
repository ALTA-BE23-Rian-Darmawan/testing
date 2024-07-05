package handler

import (
	"BE23TODO/app/middlewares"
	todos "BE23TODO/features/Todos"
	"BE23TODO/utils/responses"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoService todos.ServiceTodosInterface
}

func New(td todos.ServiceTodosInterface) *TodoHandler {
	return &TodoHandler{
		todoService: td,
	}
}

func (th *TodoHandler) CreateTodo(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	// membaca data dari request body
	newTodo := TodoRequest{}
	errBind := c.Bind(&newTodo)
	if errBind != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error bind todo: " + errBind.Error(),
		})
	}

	// mapping  dari request ke project
	inputTodo := todos.TodosEntity{
		UserID:      uint(userID),
		TodoName:    newTodo.TodoName,
		Description: newTodo.Description,
	}

	if errInsert := th.todoService.Create(inputTodo); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "Failed Create Todo: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "Failed Create Todo: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "Successfully Create Todo", nil))
}

func (th *TodoHandler) GetAllTodo(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	result, err := th.todoService.GetByUserId(uint(userID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"status":  "failed",
			"message": "error read todo",
		})
	}
	var allTodoResponse []TodoResponse
	for _, value := range result {
		allTodoResponse = append(allTodoResponse, TodoResponse{
			UserID:      value.UserID,
			ID:          value.ID,
			TodoName:    value.TodoName,
			Description: value.Description,
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"status":  "success",
		"message": "success read todo",
		"results": allTodoResponse,
	})
}

func (th *TodoHandler) DeleteTodo(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"status":  "failed",
			"message": "error convert id: " + errConv.Error(),
		})
	}
	if errInsert := th.todoService.Delete(uint(idConv), uint(userID)); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error delete todo: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "error delete todo: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "success delete todo", nil))
}

func (th *TodoHandler) UpdateTodo(c echo.Context) error {
	// Extract user ID from authentication context
	userID := middlewares.ExtractTokenUserId(c)
	if userID == 0 {
		return errors.New("user ID not found in context")
	}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error converting id: " + errConv.Error(),
		})
	}

	var updateData TodoRequest
	if err := c.Bind(&updateData); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"status":  "failed",
			"message": "error binding todo: " + err.Error(),
		})
	}

	inputProject := todos.TodosEntity{
		TodoName:    updateData.TodoName,
		Description: updateData.Description,
	}

	if errInsert := th.todoService.Update(uint(idConv), uint(userID), inputProject); errInsert != nil {
		if strings.Contains(errInsert.Error(), "validation") {
			return c.JSON(http.StatusBadRequest, responses.JSONWebResponse(http.StatusBadRequest, "error", "error updating project: "+errInsert.Error(), nil))
		}
		return c.JSON(http.StatusInternalServerError, responses.JSONWebResponse(http.StatusInternalServerError, "error", "error updating todo: "+errInsert.Error(), nil))
	}

	return c.JSON(http.StatusCreated, responses.JSONWebResponse(http.StatusCreated, "success", "successfully updated todo", nil))
}
