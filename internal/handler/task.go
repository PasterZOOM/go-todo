package handler

import (
	"go-todo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid todo_list_id param")
		return
	}

	var input domain.Task
	if err := c.ShouldBindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Task.Create(userId, todoListId, input)
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

type GetAllTasksResponse struct {
	Data []domain.Task `json:"data"`
}

func (h *Handler) getAllTasks(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	todoListId, err := strconv.Atoi(c.Param("todo_list_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid todo_list_id param")
		return
	}

	tasks, err := h.services.Task.GetAllTasks(userId, todoListId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, GetAllTasksResponse{
			Data: tasks,
		},
	)
}

type GetTaskByIdResponse struct {
	Data domain.Task `json:"data"`
}

func (h *Handler) getTaskById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	task, err := h.services.Task.GetById(userId, taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(
		http.StatusOK, GetTaskByIdResponse{
			Data: task,
		},
	)
}

func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	var input domain.UpdateTaskInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.Task.Update(userId, taskId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.Task.Delete(userId, taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
