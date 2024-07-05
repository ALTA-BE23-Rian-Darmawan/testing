package routes

import (
	"BE23TODO/app/middlewares"
	"BE23TODO/utils/encrypts"

	_todoData "BE23TODO/features/Todos/dataTodos"
	_todoHandler "BE23TODO/features/Todos/handler"
	_todoService "BE23TODO/features/Todos/service"
	_userData "BE23TODO/features/Users/dataUsers"
	_userHandler "BE23TODO/features/Users/handler"
	_userService "BE23TODO/features/Users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(e *echo.Echo, db *gorm.DB) {
	hashService := encrypts.NewHashService()
	userData := _userData.New(db)
	userService := _userService.New(userData, hashService)
	userHandlerAPI := _userHandler.New(userService)

	todoData := _todoData.New(db)
	todoService := _todoService.New(todoData)
	todoHandlerAPI := _todoHandler.New(todoService)

	//userHandler
	e.POST("/users", userHandlerAPI.Register)
	e.POST("/login", userHandlerAPI.Login)

	//projectHandler
	e.GET("/todos", todoHandlerAPI.GetAllTodo, middlewares.JWTMiddleware())
	e.POST("/todos", todoHandlerAPI.CreateTodo, middlewares.JWTMiddleware())
	e.PUT("/todos/:id", todoHandlerAPI.UpdateTodo, middlewares.JWTMiddleware())
	e.DELETE("/todos/:id", todoHandlerAPI.DeleteTodo, middlewares.JWTMiddleware())
}
