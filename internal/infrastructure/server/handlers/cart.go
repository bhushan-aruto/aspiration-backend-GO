package handlers

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type CartHandler struct {
	Db *database.MongoDBdatabase
}

func NewCartHandler(db *database.MongoDBdatabase) *CartHandler {
	return &CartHandler{
		Db: db,
	}
}

func (h *CartHandler) AddToCartHandler(ctx echo.Context) error {
	var item entity.CartItem

	if err := ctx.Bind(&item); err != nil {
		return ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}

	usecase := usecase.NewCartUseCase(h.Db)

	if err := usecase.AddCartItemUseCase(&item); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "added to cart",
	})

}

func (h *CartHandler) GetAllCartCourseHandler(ctx echo.Context) error {
	userID := ctx.QueryParam("user_id")

	if userID == "" {
		return ctx.JSON(400, echo.Map{
			"error": "user_id is required",
		})
	}

	usecase := usecase.NewCartUseCase(h.Db)
	course, err := usecase.GetCArtItemSUseCase(userID)

	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": course,
	})

}

func (h *CartHandler) DeleteCartCourseHandler(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	courseID := ctx.Param("course_id")

	usecase := usecase.NewCartUseCase(h.Db)
	if err := usecase.DeletecartItem(userID, courseID); err != nil {
		return ctx.JSON(500,
			echo.Map{
				"error": err.Error(),
			})

	}

	return ctx.JSON(200, echo.Map{
		"message": "Deleted successfully",
	})

}
