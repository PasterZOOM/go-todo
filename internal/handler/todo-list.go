package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodoList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	c.JSON(
		http.StatusOK, map[string]interface{}{
			userCtx: id,
		},
	)
}
func (h *Handler) getAllTodoLists(c *gin.Context) {}
func (h *Handler) getTodoListById(c *gin.Context) {}
func (h *Handler) updateTodoList(c *gin.Context)  {}
func (h *Handler) deleteTodoList(c *gin.Context)  {}
