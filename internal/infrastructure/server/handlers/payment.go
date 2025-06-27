package handlers

import (
	"os"
	"time"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PaymentHandler struct {
	GateWay repository.PaymentGateWayRepo
	DbRepo  repository.PaymentDatabaseRepo
}

func NewPaymentHandler(GateWay repository.PaymentGateWayRepo, dbRepo repository.PaymentDatabaseRepo) *PaymentHandler {
	return &PaymentHandler{
		GateWay: GateWay,
		DbRepo:  dbRepo,
	}
}

func (h *PaymentHandler) CreateOrderHandler(ctx echo.Context) error {

	req := new(models.CreatePaymentOrderRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.JSON(400, echo.Map{
			"message": "invalid json request body format",
		})
	}

	paymentUseCase := usecase.NewPaymentGateWayUseCase(h.GateWay, h.DbRepo)

	bytes, err := paymentUseCase.CreateOrder(req.CoursesId)

	if err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSONBlob(200, bytes)

}

func (h *PaymentHandler) VerifyPaymentHandler(ctx echo.Context) error {
	var req models.VerifyPaymentRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid request: " + err.Error(),
		})
	}

	secret := os.Getenv("PAYMENT_GATEWAY_KEYSECRETE")
	paymentUseCase := usecase.NewPaymentGateWayUseCase(h.GateWay, h.DbRepo)

	valid := paymentUseCase.VerifyOrderUseCase(req.RazorpayOrderID, req.RazorpayPaymentID, req.RazorpaySignature, secret)
	if !valid {
		return ctx.JSON(401, echo.Map{
			"success": false,
			"error":   "Signature mismatch",
		})
	}

	if err := h.DbRepo.AppendPurchasedCoursesIds(req.UserID, req.CourseIDs); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to update purchased courses: " + err.Error(),
		})
	}

	for _, courseID := range req.CourseIDs {
		price, err := h.DbRepo.GetSingleCoursePriceById(courseID)
		if err != nil {
			return ctx.JSON(500, echo.Map{
				"error": "failed to get course price: " + err.Error(),
			})
		}

		purchase := &entity.PurchaseHistory{
			ID:            uuid.New().String(),
			UserID:        req.UserID,
			CourseID:      courseID,
			Date:          time.Now().Format("2006-01-02 15:04:05"),
			Amount:        price,
			PaymentMethod: "razorpay",
		}

		if err := h.DbRepo.SavePurchaseHistory(purchase); err != nil {
			return ctx.JSON(500, echo.Map{
				"error": "failed to save purchase history: " + err.Error(),
			})
		}
		
		if err := paymentUseCase.DeleteCartUseCase(req.UserID, courseID); err != nil {
			return ctx.JSON(500, echo.Map{
				"error": "failed to deleted carted course : " + err.Error(),
			})
		}
	}

	return ctx.JSON(200, echo.Map{
		"success": true,
		"message": "Payment verified and courses updated",
	})
}
