package handler

import (
	"go-todo/internal/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary SignUp
// @Tags v1 — auth
// @Description create account
// @ID create-account
// @Accept json
// @Produce json
// @Param input body domain.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/v1/auth/sign-up [post]
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

type SignInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary Sign in
// @Tags v1 — auth
// @Description login
// @ID sign-in
// @Accept json
// @Produce json
// @Param input body SignInInput true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Failure default {object} ErrorResponse
// @Router /api/v1/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input SignInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(
		http.StatusOK, map[string]interface{}{
			"token": token,
		},
	)
}
