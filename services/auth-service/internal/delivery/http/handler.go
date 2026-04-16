package http

import (
	"auth-service/internal/usecase"
	"errors"
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	authUsecase *usecase.AuthUsecase
}

func NewHandler(u *usecase.AuthUsecase) *Handler {
	return &Handler{u}
}

type registerRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (h *Handler) Register(c *gin.Context) {
	var req registerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authUsecase.Register(req.Email, req.Password)
    if err != nil {
        // TAMBAHKAN LOG INI:
        fmt.Printf("[DEBUG] Register Error: %v\n", err) 

        if errors.Is(err, usecase.ErrEmailAlreadyExists) {
            c.JSON(http.StatusConflict, gin.H{"error": "email already registered"})
            return
        }
        // Kirim error aslinya ke JSON agar bisa dibaca di Postman (Hanya untuk Debug!)
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	c.JSON(http.StatusCreated, gin.H{"message": "registration successful"})
}

func (h *Handler) Login(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidCredentials) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to login"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": token})
}