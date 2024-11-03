package handler

import (
	"go-todo/cmd/docs"
	"go-todo/internal/service"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRouter() *gin.Engine {
	route := gin.New()

	docs.SwaggerInfo.BasePath = "/"

	v1 := route.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("sign-up", h.signUp)
			auth.POST("sign-in", h.signIn)
		}

		todoLists := v1.Group("todo-lists", h.userIdentity)
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

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return route
}
