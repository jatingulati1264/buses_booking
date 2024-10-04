package controllers

import (
    "net/http"
    "bus_booking/services"
    "github.com/gin-gonic/gin"
)

type AuthController struct {
    service services.AuthService
}

func NewAuthController(authService services.AuthService) *AuthController {
    return &AuthController{service: service}
}

func (ctrl *AuthController) Login(c *gin.Context) {
    var credentials struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&credentials); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := ctrl.service.Login(credentials.Username, credentials.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
