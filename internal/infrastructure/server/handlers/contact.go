package handlers

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/config"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type ContactHandler struct {
	contactUC usecase.ContactUsecase
	cfg       *config.Config
}

func NewContactHandler(contactUC usecase.ContactUsecase, cfg *config.Config) *ContactHandler {
	return &ContactHandler{contactUC: contactUC,
		cfg: cfg}
}

func (h *ContactHandler) HandleEmailContact(ctx echo.Context) error {
	var req models.EmailContactRequest

	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(400,
			echo.Map{"error": err.Error()})
	}

	if err := h.contactUC.SendEmailUseCse(req, h.cfg); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})

	}
	return ctx.JSON(200, echo.Map{"message": "email sent successfully"})
}

func (h *ContactHandler) HandleWhatsAppContact(c echo.Context) error {
	var req models.WhatsAppContactRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error()})
	}

	link, err := h.contactUC.GenerateWhatsAppURL(req, h.cfg)
	if err != nil {
		return c.JSON(500, echo.Map{
			"error": "failed to generate link"})
	}

	return c.JSON(200, echo.Map{"redirect_url": link})
}

func (h *ContactHandler) CourseEnqiryHandler(c echo.Context) error {
	var req models.EamilCourseEnquiryRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(400, echo.Map{
			"error": err.Error(),
		})
	}
	if err := h.contactUC.SendEnquiryCourseUsecase(req, h.cfg); err != nil {
		return c.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(200, echo.Map{
		"message": "enquiry sent successfully",
	})

}
