package handler

import (
	"go-todo/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(
		http.StatusOK, map[string]interface{}{
			"id": id,
		},
	)
}

func (h *Handler) signIn(c *gin.Context) {

}
