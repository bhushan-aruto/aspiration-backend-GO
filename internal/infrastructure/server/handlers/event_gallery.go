package handlers

import (
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type EventGalleryHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewEventGalleryHandler(bucketName, dirName, cloudFrontBaseUrl string, db *database.MongoDBdatabase, s3 *storage.S3Connection) *EventGalleryHandler {
	return &EventGalleryHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          db,
		s3:                s3,
	}

}

func (h *EventGalleryHandler) EventGalleryUploadImageHandler(ctx echo.Context) error {
	file, err := ctx.FormFile("image")

	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while getting the image",
			},
		)
	}
	if file == nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "file was empty",
			},
		)
	}
	fileSrc, err := file.Open()

	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while opening the image",
			},
		)
	}
	defer fileSrc.Close()

	eventGalleryUseCase := usecase.NewEventGalleryUseCase(
		h.database,
		storage.NewEvenyGallleryRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	fileName := file.Filename
	if err := eventGalleryUseCase.UploadEventImage(fileName, fileSrc); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "image uploaded successfully",
	})
}

func (h *EventGalleryHandler) GetAllEventGalleryImagesHandler(ctx echo.Context) error {
	eventGalleryUseCase := usecase.NewEventGalleryUseCase(
		h.database,
		storage.NewEvenyGallleryRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	images, err := eventGalleryUseCase.GetEventSectionAllImages()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": images,
	})

}

func (h *EventGalleryHandler) DeleteImageHandler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	if strings.TrimSpace(fileName) == "" {
		return ctx.JSON(400, echo.Map{"error": "filename is required"})
	}

	eventGalleryUseCase := usecase.NewEventGalleryUseCase(
		h.database,
		storage.NewEvenyGallleryRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := eventGalleryUseCase.DeleteImageByFileName(fileName); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "image deleted successfully",
	})
}
