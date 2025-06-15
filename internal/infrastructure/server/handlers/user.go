package handlers

import (
	"log"
	"net/http"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/messaging"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	db       *database.MongoDBdatabase
	producer *messaging.RabbitMQProducer
}

func NewUserHandler(db *database.MongoDBdatabase, producer *messaging.RabbitMQProducer) *UserHandler {
	return &UserHandler{
		db:       db,
		producer: producer,
	}
}

func (h *UserHandler) UserSignupHandler(ctx echo.Context) error {

	var req models.UserSignupRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request body"})
	}
	usecase := usecase.NewUserUseCase(h.db)
	user, err := usecase.UserSignUp(req.UserName, req.Email, req.Password, req.PhoneNumber)
	if err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	go func() {
		if err := h.producer.SendWelcomeMail(user.Email, user.UserName); err != nil {
			ctx.Logger().Error("failed to send  the welcome mail", err)
		}

	}()

	resp := models.UserSignupResponse{
		ID:          user.Id,
		UserName:    user.UserName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	return ctx.JSON(200, resp)
}

func (h *UserHandler) SentOTPHandler(ctx echo.Context) error {
	var req models.UserOTPRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid  requeust body",
		})
	}

	useCase := usecase.NewUserUseCase(h.db)

	otp, err := useCase.SendOTP(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	if err := h.producer.SendOTP(req.Email, otp); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to send OTP",
		})
	}
	return ctx.JSON(200, echo.Map{
		"message": "OTP sent successfully",
	})
}

func (h *UserHandler) VerifyAndLoginHandler(ctx echo.Context) error {
	var req models.UserVerifyOTPRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid request body",
		})
	}

	useCase := usecase.NewUserUseCase(h.db)

	valid, err := useCase.VerifyOTP(req.Email, req.Otp)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to verify OTP",
		})
	}
	if !valid {
		return ctx.JSON(401, echo.Map{
			"error": "invalid or expired OTP",
		})
	}

	token, err := useCase.UserLogin(req.Email, req.Password)
	if err != nil {
		log.Printf("Login error for %q: %v\n", req.Email, err)
		return ctx.JSON(401, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"token": token,
	})
}

func (h *UserHandler) GetUserByIdHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	useCase := usecase.NewUserUseCase(h.db)
	user, err := useCase.GetUserById(id)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	resp := models.GetUserResponse{
		UserName: user.UserName,
		Email:    user.Email,
	}
	return ctx.JSON(200, echo.Map{
		"user": resp,
	})

}

func (h *UserHandler) ForgotPasswordSendOTPHandler(ctx echo.Context) error {
	var req models.ForgotPasswordRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	usecase := usecase.NewUserUseCase(h.db)

	otp, err := usecase.SendForgotPasswordOTP(req.Email)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	if err := h.producer.SendOTP(req.Email, otp); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"message": "OTP sent succesfully",
	})

}

func (h *UserHandler) ResetPasswordHandler(ctx echo.Context) error {
	var req models.ResetPasswordRequest

	if err := ctx.Bind(&req); err != nil {

		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	usecase := usecase.NewUserUseCase(h.db)

	if err := usecase.ResetPassword(req.Email, req.OTP, req.NewPassword); err != nil {
		return ctx.JSON(400, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(200, echo.Map{
		"message": "password reset succesfully",
	})

}
