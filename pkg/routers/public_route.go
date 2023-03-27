package routers

import (
	"Template/pkg/controllers"
	"Template/pkg/controllers/healthchecks"
	"Template/pkg/utils/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealth)

	// Sample Endpoints for dashboard testing
	dashboardRoutes := v1Endpoint.Group("/dashboards")
	dashboardRoutes.Post("/login-authentication", controllers.ReportsLoginAuth)
	dashboardRoutes.Post("/change-password", controllers.ChangePassword)
	dashboardRoutes.Post("/create-account", controllers.CreateReportsAccount)
	dashboardRoutes.Post("/change-password", controllers.ChangePassword)
	dashboardRoutes.Get("/test", controllers.HelloWorld)

	// protected route
	protectedRoutes := dashboardRoutes.Group("/protected", middleware.AuthMiddleware)
	protectedRoutes.Get("/app-list", controllers.ListApps)
}

func SetupPublicRoutesB(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealthB)
}
