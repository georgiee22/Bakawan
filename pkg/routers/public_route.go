package routers

import (
	"Template/pkg/controllers"
	"Template/pkg/controllers/healthchecks"

	"github.com/gofiber/fiber/v2"
)

func SetupPublicRoutes(app *fiber.App) {
	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealth)

	// Sample Endpoints for data lake testing
	datalakeRoutes := v1Endpoint.Group("/data-lake")
	datalakeRoutes.Post("/login-authentication", controllers.ReportsLoginAuth)
	datalakeRoutes.Post("/change-password", controllers.ChangePassword)
}

func SetupPublicRoutesB(app *fiber.App) {

	// Endpoints
	apiEndpoint := app.Group("/api")
	publicEndpoint := apiEndpoint.Group("/public")
	v1Endpoint := publicEndpoint.Group("/v1")

	// Service health check
	v1Endpoint.Get("/", healthchecks.CheckServiceHealthB)
}
