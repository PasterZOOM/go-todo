package handler

import (
	"go-todo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create todo list
// @Security ApiKeyAuth
// @Tags v1 — todo-lists
// @Description create todo list
// @ID create-todo-list
// @Accept  json
// @Produce  json
// @Param input body domain.TodoList true "list info"
// @Success 201 {string} string ""
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists [post]
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

// @Summary Get all todo lists
// @Security ApiKeyAuth
// @Tags v1 — todo-lists
// @Description get all todo lists
// @ID get-all-todo-lists
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllTodoListsResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists [get]
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

// @Summary Get todo list by id
// @Security ApiKeyAuth
// @Tags v1 — todo-lists
// @Description get todo list by id
// @ID get-todo-list-by-id
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Success 200 {object} GetTodoListByIdResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id [get]
func (h *Handler) getTodoListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid todo_list_id param")
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

// @Summary Update todo list
// @Security ApiKeyAuth
// @Tags v1 — todo-lists
// @Description update todo list
// @ID update-todo-list
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Param input body domain.UpdateTodoListInput true "list info"
// @Success 204
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id [put]
func (h *Handler) updateTodoList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid todo_list_id param")
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

// @Summary Delete todo list
// @Security ApiKeyAuth
// @Tags v1 — todo-lists
// @Description delete todo list
// @ID delete-todo-list
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Success 204
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id [delete]
func (h *Handler) deleteTodoList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid todo_list_id param")
		return
	}

	err = h.services.TodoList.Delete(userId, todoListId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
