package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prismgo/framework/foundation"
	"prismgo/app/http/controllers"
	"prismgo/routes"
)

// RegisterRoutes registers application HTTP routes with the Prismgo router.
func RegisterRoutes(app *foundation.Application, _ *gin.Engine) error {
	routes.Register(routes.Dependencies{
		WelcomeController: welcomeController(),
	})
	return nil
}

// welcomeController builds the default welcome controller for a fresh application.
func welcomeController() *controllers.WelcomeController {
	return controllers.NewWelcomeController()
}
