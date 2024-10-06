package handler

import "github.com/gin-gonic/gin"

type Handler struct{}

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
			todoLists.GET("/:id", h.getTodoListById)
			todoLists.PUT("/:id", h.updateTodoList)
			todoLists.DELETE("/:id", h.deleteTodoList)

			tasks := todoLists.Group(":id/task")
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
