package handler

import (
	"go-todo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTodoList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input domain.TodoList
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.TodoList.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusCreated, map[string]interface{}{
			"id": id,
		},
	)
}

type GetAllTodoListsResponse struct {
	Data []domain.TodoList `json:"data"`
}

func (h *Handler) getAllTodoLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoLists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, GetAllTodoListsResponse{
			Data: todoLists,
		},
	)
}

type GetTodoListByIdResponse struct {
	Data domain.TodoList `json:"data"`
}

func (h *Handler) getTodoListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	todoList, err := h.services.TodoList.GetById(userId, todoListId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, GetTodoListByIdResponse{
			Data: todoList,
		},
	)
}

func (h *Handler) updateTodoList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input domain.UpdateTodoListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.TodoList.Update(userId, todoListId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteTodoList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.TodoList.Delete(userId, todoListId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
