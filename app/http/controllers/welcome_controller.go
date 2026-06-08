package controllers

import (
	"net/http"
	"prismgo/config"

	"github.com/gin-gonic/gin"
)

// WelcomeController handles the default API landing endpoint.
type WelcomeController struct{}

// NewWelcomeController creates the default welcome controller.
func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

// Show returns a small JSON payload proving that the application is running.
func (c *WelcomeController) Show(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"framework": "prismgo",
		"name":      config.GetString("app.name", "Prismgo"),
		"message":   "Welcome to Prismgo.",
	})
}
