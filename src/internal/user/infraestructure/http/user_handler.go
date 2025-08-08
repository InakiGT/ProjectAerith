package http

import (
	"net/http"
	"rapi-pedidos/src/internal/user/application"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	createUser   application.CreateUserCommand
	findAllUsers application.FindAllUsersCommand
}

func NewUserHandler(createUser application.CreateUserCommand, findAllUsers application.FindAllUsersCommand) *UserHandler {
	return &UserHandler{
		createUser:   createUser,
		findAllUsers: findAllUsers,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if _, err := h.createUser.Execute(ctx, input.Username, input.Email, input.Password); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "user created"})
}

func (h *UserHandler) FindAllUsers(ctx *gin.Context) {
	users, err := h.findAllUsers.Execute(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}
