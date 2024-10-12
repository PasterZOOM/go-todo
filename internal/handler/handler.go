package handler

import (
	"go-todo/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
	route := gin.New()

	auth := route.Group("/auth")
	{
		auth.POST("sign-up", h.signUp)
		auth.POST("sign-in", h.signIn)
	}

	api := route.Group("/api")
	{
		todoLists := api.Group("todo-lists")
		{
			todoLists.POST("/", h.createTodoList)
			todoLists.GET("/", h.getAllTodoLists)
			todoLists.GET("/:todo_list_id", h.getTodoListById)
			todoLists.PUT("/:todo_list_id", h.updateTodoList)
			todoLists.DELETE("/:todo_list_id", h.deleteTodoList)

			tasks := todoLists.Group(":todo_list_id/tasks")
			{
				tasks.POST("/", h.createTask)
				tasks.GET("/", h.getAllTasks)
				tasks.GET("/:task_id", h.getTaskById)
				tasks.PUT("/:task_id", h.updateTask)
				tasks.DELETE("/:task_id", h.deleteTask)
			}
		}
	}

	return route
}
