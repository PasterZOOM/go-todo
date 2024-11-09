package handler

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get user by ID
// @Description Retrieve user details by user ID
// @Tags v1 — users
// @Produce json
// @Param user_id path int true "User ID"
// @Success 200 {object} domain.UserResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/v1/users/:user_id [get]
func (h *Handler) getUserData(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid user_id param")
		return
	}

	user, err := h.services.User.GetUserData(userId)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			newErrorResponse(c, http.StatusNotFound, fmt.Sprintf("user with id %d not found", userId))
		default:
			newErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("unable to fetch user: %s", err))
		}
		return
	}

	c.JSON(
		http.StatusOK, user,
	)
}
