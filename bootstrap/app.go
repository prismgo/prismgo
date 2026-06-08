package bootstrap

import (
	appcmd "prismgo/app/cmd"
	apphttp "prismgo/app/http"
	appmiddleware "prismgo/app/http/middleware"
	appschedule "prismgo/app/schedule"
	_ "prismgo/database/migrations"

	"github.com/gin-gonic/gin"
	"github.com/prismgo/framework/foundation"
	httpmiddleware "github.com/prismgo/framework/http/middleware"
)

// NewApplication creates the project application instance.
func NewApplication(basePath ...string) *foundation.Application {
	return foundation.Configure(basePath...).
		WithProviders(Providers()...).
		WithRouting(func(r *foundation.Routing) {
			r.Commands(appcmd.CommandFactories()...)
			r.Schedules(appschedule.Register)
			r.Routes(apphttp.RegisterRoutes)
		}).
		WithMiddleware(func(m *foundation.Middleware) {
			m.Prepend(func(engine *gin.Engine) {
				engine.Use(httpmiddleware.RequestID())
				engine.Use(appmiddleware.Example())
			})
		}).
		Create()
}
