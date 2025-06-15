package handlers

import (
	"fmt"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type PhilosopySectionHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewPhilosopySectionHandler(bucketName, dirName, cloudFrontBaseUrl string, database *database.MongoDBdatabase, s3 *storage.S3Connection) *PhilosopySectionHandler {
	return &PhilosopySectionHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          database,
		s3:                s3,
	}
}

func (h *PhilosopySectionHandler) GetPhilosopysectionHandler(ctx echo.Context) error {
	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	philosopySection, err := philosopySectionUseCase.GetPhilosopySectionUseCase()
	if err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
		return err
	}

	return ctx.JSON(200, echo.Map{
		"data": philosopySection,
	})
}

func (h *PhilosopySectionHandler) CreatePhilosopySectionHandler(ctx echo.Context) error {
	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := philosopySectionUseCase.CreatePhilosopySectionUseCase(); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "Philospy section section created successfully",
	})
}

func (h *PhilosopySectionHandler) UpdatePhilosopySectionImage1Handler(ctx echo.Context) error {
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

	fileNameArr := strings.Split(file.Filename, ".")

	inputFileType := fileNameArr[len(fileNameArr)-1]

	fileName := fmt.Sprintf("philosopyImage1.%v", inputFileType)

	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := philosopySectionUseCase.UpdatePhilospySectionImage1(fileName, fileSrc); err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"error": err.Error(),
			},
		)
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image updated successfully",
		},
	)

}

func (h *PhilosopySectionHandler) UpdatePhilosopySectionImage2Handler(ctx echo.Context) error {
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

	fileNameArr := strings.Split(file.Filename, ".")

	inputFileType := fileNameArr[len(fileNameArr)-1]

	fileName := fmt.Sprintf("philosopyImage2.%v", inputFileType)

	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := philosopySectionUseCase.UpdatePhilospySectionImage2(fileName, fileSrc); err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"error": err.Error(),
			},
		)
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image updated successfully",
		},
	)

}

func (h *PhilosopySectionHandler) DeletePhilosopySectionImage1Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := philosopySectionUseCase.DeletePhilosopySectionImage1(fileName); err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"error": err.Error(),
			},
		)
	}
	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}

func (h *PhilosopySectionHandler) DeletePhilosopySectionImage2Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	philosopySectionUseCase := usecase.NewPhilosopyUseCase(
		h.database,
		storage.NewPhilosopySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := philosopySectionUseCase.DeletePhilosopySectionImage2(fileName); err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"error": err.Error(),
			},
		)
	}
	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}
