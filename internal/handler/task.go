package handler

import (
	"go-todo/internal/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create task
// @Security ApiKeyAuth
// @Tags v1 — tasks
// @Description create task
// @ID create-task
// @Accept  json
// @Produce  json
// @Param input body domain.Task true "task info"
// @Success 201 {string} string ""
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id/tasks [post]
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

// @Summary Get all tasks
// @Security ApiKeyAuth
// @Tags v1 — tasks
// @Description get all tasks for todo list
// @ID get-all-tasks
// @Accept  json
// @Produce  json
// @Success 200 {object} GetAllTasksResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id/tasks [get]
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

// @Summary Get task by id
// @Security ApiKeyAuth
// @Tags v1 — tasks
// @Description get task by id
// @ID get-task-by-id
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Param task_id path int true "Task ID"
// @Success 200 {object} GetTaskByIdResponse
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id/tasks/:task_id [get]
func (h *Handler) getTaskById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid task_id param")
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

// @Summary Update task
// @Security ApiKeyAuth
// @Tags v1 — tasks
// @Description update task
// @ID update-task
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Param task_id path int true "Task ID"
// @Param input body domain.UpdateTaskInput true "task info"
// @Success 204
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id/tasks/:task_id [put]
func (h *Handler) updateTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid task_id param")
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

// @Summary Delete task
// @Security ApiKeyAuth
// @Tags v1 — tasks
// @Description delete task
// @ID delete-task
// @Accept json
// @Produce json
// @Param todo_list_id path int true "Todo list ID"
// @Param task_id path int true "Task ID"
// @Success 204
// @Failure 400,404,500 {object} ErrorResponse
// @Router /api/v1/todo-lists/:todo_list_id/tasks/:task_id [delete]
func (h *Handler) deleteTask(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	taskId, err := strconv.Atoi(c.Param("task_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid task_id param")
		return
	}

	err = h.services.Task.Delete(userId, taskId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}
