package handler

import (
	"net/http"

	"github.com/app/gin-postgres-api/dto"
	"github.com/app/gin-postgres-api/internal/service"
	"github.com/app/gin-postgres-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(c *gin.Context)
}

type AuthHandlerImpl struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
	return &AuthHandlerImpl{authService: authService}
}

func (h *AuthHandlerImpl) Login(c *gin.Context) {
	var body dto.AuthLoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, http.StatusBadRequest, "username and password are required")
		return
	}

	token, err := h.authService.Login(body.Username, body.Password)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, http.StatusOK, gin.H{"token": token})
}
