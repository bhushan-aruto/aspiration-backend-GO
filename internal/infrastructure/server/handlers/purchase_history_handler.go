package handlers

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/labstack/echo/v4"
)

type PurchaseHistoryHandler struct {
	Db *database.MongoDBdatabase
}

func NewPurchaseHistoryHandler(db *database.MongoDBdatabase) *PurchaseHistoryHandler {
	return &PurchaseHistoryHandler{
		Db: db,
	}
}

func (h *PurchaseHistoryHandler) GetPurchaseHistoryByUserIdHandler(ctx echo.Context) error {

	userId := ctx.Param("user_id")

	purchaseHistory, err := h.Db.GetPurchaseHistoryByUserId(userId)
	if err != nil {
		ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"data": purchaseHistory,
	})

}
