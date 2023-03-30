package middleware

import (
	"Template/pkg/models"
	"Template/pkg/models/response"
	"Template/pkg/utils/go-utils/database"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware(c *fiber.Ctx) error {
	// Get token from Authorization header
	header := c.Get("Authorization")
	if header == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Missing authorization header",
		})
	}
	token := strings.TrimPrefix(header, "Bearer ")

	// token := c.Cookies("token")

	// Parse token
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Check if token is valid
	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok && parsedToken.Valid {
		username, ok := claims["username"]
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}
		var account models.User_Accounts
		if err := database.DBConn.Find(&account).Where("username = ?", username).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "User not found",
			})
		}
		userid := claims["userid"]
		c.Locals("userid", userid)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": "Invalid token",
	})
}

// verify if user is authorized to acces the app specific dashboards
func VerifyAuth1stLayer(c *fiber.Ctx) error {
	app_id := c.Params("id")
	// user_id := c.Cookies("userid")

	// var user_apps []string

	// get user apps
	var checker bool
	err := database.DBConn.Raw("SELECT EXISTS(SELECT 1 FROM access_account_apps WHERE user_id = ? AND app_id = ?)", 1, app_id).Row().Scan(&checker)
	if err != nil {
		return c.JSON(response.ResponseModel{
			RetCode: "203",
			Message: "query error",
			Data:    err.Error(),
		})
	}

	// check if access exist
	if !checker {
		return c.JSON(response.ResponseModel{
			RetCode: "400",
			Message: "fail",
			Data:    "access denied",
		})
	}

	return c.Next()
}
