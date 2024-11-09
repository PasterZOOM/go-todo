package handler

import (
	"go-todo/internal/domain"
	"go-todo/internal/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// @Summary Register a new user
// @Tags v1 — auth
// @DescriptionRegister a new user with name, username, password, and email
// @ID create-account
// @Accept json
// @Produce json
// @Param input body domain.UserInput true "User info"
// @Success 201 {object} domain.UserResponse "Created user data with ID, timestamps"
// @Failure 400 {object} utils.ValidationErrorsResponse "Invalid input"
// @Failure 409 {object} ErrorResponse "Username or email already in use"
// @Failure 500 {object} ErrorResponse "Failed to create user"
// @Router /api/v1/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input domain.UserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		validationResponse := utils.HandleValidationErrors(err)
		c.JSON(http.StatusBadRequest, validationResponse)
		return
	}

	createdUser, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "username"), strings.Contains(err.Error(), "email"):
			newErrorResponse(c, http.StatusConflict, err.Error())
		default:
			newErrorResponse(c, http.StatusInternalServerError, "Failed to create user: "+err.Error())
		}
		return
	}

	c.JSON(http.StatusCreated, createdUser)
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
