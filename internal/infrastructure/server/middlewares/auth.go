package middlewares

import (
	"net/http"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		authHeader := ctx.Request().Header.Get("Authorization")
		if authHeader == "" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid token format"})
		}

		tokenStr := parts[1]
		claims, err := utils.ValidateJWT(tokenStr)
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid or expired token"})
		}

		ctx.Set("userId", claims.UserId)

		return next(ctx)
	}
}
