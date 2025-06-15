package handlers

import (
	"fmt"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type AboutSectionHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewAboutSectionHandler(bucketName, dirName, cloudFrontBaseUrl string, database *database.MongoDBdatabase, s3 *storage.S3Connection) *AboutSectionHandler {
	return &AboutSectionHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          database,
		s3:                s3,
	}
}

func (h *AboutSectionHandler) GetAboutSectionHandler(ctx echo.Context) error {
	abooutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	aboutSection, err := abooutSectionUseCase.GetAboutSectionUseCase()

	if err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(200, echo.Map{
		"data": aboutSection,
	})

	return nil
}

func (h *AboutSectionHandler) CreateAboutSectionHandler(ctx echo.Context) error {
	abooutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := abooutSectionUseCase.CreateAboutSection(); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "about section created successfully",
	})

}

func (h *AboutSectionHandler) UpdateAbouSectionImage1Handler(ctx echo.Context) error {
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

	fileName := fmt.Sprintf("image1.%v", inputFileType)
	

	aboutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := aboutSectionUseCase.UpdateAboutSectionImage1(fileName, fileSrc); err != nil {
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

func (h *AboutSectionHandler) UpdateAboutSectionImage2Handler(ctx echo.Context) error {
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

	fileName := fmt.Sprintf("image2.%v", inputFileType)

	aboutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := aboutSectionUseCase.UpdateAboutSectionImage2(fileName, fileSrc); err != nil {
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

func (h *AboutSectionHandler) DeleteAboutSectionImage1Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	aboutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := aboutSectionUseCase.DeleteAboutSectionImage1(fileName); err != nil {
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

func (h *AboutSectionHandler) DeleteAboutSectionImage2Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	aboutSectionUseCase := usecase.NewAboutSectionUseCase(
		h.database,
		storage.NewAboutSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := aboutSectionUseCase.DeleteAboutSectionImage2(fileName); err != nil {
		return ctx.JSON(
			500,
			echo.Map{
				"error": err.Error(),
			},
		)
	}

	return ctx.JSON(200,
		echo.Map{
			"message": "image deleted succefully",
		})
}
