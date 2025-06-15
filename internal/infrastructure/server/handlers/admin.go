package handlers

import (
	"net/http"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AdminHandler struct {
	Db *database.MongoDBdatabase
}

func NewAdminHandler(db *database.MongoDBdatabase) *AdminHandler {
	
	return &AdminHandler{
		Db: db,
	}
}

func (h *AdminHandler) AdminSignupHandler(ctx echo.Context) error {
	var req models.AdminSignupRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error ": "invalid request body"})
	}

	useCase := usecase.NewAdminUseCase(h.Db)
	admin, err := useCase.AdminSignUp(req.AdminName, req.Email, req.PhoneNumber, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}
	response := models.AdimnSignupResponse{
		ID:          admin.Id,
		AdminName:   admin.AdminName,
		Email:       admin.Email,
		PhoneNumber: admin.PhoneNumber,
	}

	return ctx.JSON(http.StatusCreated, response)
}

func (h *AdminHandler) AdminLoginHandler(ctx echo.Context) error {
	var credential models.AdminLoginRequest

	if err := ctx.Bind(&credential); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid requets  body"})
	}
	useCase := usecase.NewAdminUseCase(h.Db)
	token, err := useCase.AdminLogin(credential.Email, credential.Password)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid credential "})
	}

	return ctx.JSON(http.StatusOK, echo.Map{
		"token": token,
	})

}
