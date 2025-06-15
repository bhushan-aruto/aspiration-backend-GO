package handlers

import (
	"fmt"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type StorySectionHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewStorySectionHandler(bucketName, dirName, CloudFrontBaseUrl string, db *database.MongoDBdatabase, s3 *storage.S3Connection) *StorySectionHandler {
	return &StorySectionHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: CloudFrontBaseUrl,
		database:          db,
		s3:                s3,
	}
}

func (h *StorySectionHandler) GetStorySectionHandler(ctx echo.Context) error {
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	storySection, err := storySectionUseCase.GetStorySEctionUseCase()
	if err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
		return err
	}

	ctx.JSON(200, echo.Map{
		"data": storySection,
	})
	return nil
}

func (h *StorySectionHandler) CreateStorySectionHandler(ctx echo.Context) error {
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.CreateStorySectionUseCase(); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "story section created successfully",
	})
}

func (h *StorySectionHandler) UpdateStorySectionImage1Handler(ctx echo.Context) error {
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

	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.UpdateStorySEctionImage1(fileName, fileSrc); err != nil {
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

func (h *StorySectionHandler) UpdateStorySectionImage2Handler(ctx echo.Context) error {
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

	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.UpdateStorySEctionImage2(fileName, fileSrc); err != nil {
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

func (h *StorySectionHandler) UpdateStorySectionImage3Handler(ctx echo.Context) error {
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

	fileName := fmt.Sprintf("image3.%v", inputFileType)

	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.UpdateStorySEctionImage3(fileName, fileSrc); err != nil {
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
func (h *StorySectionHandler) UpdateStorySectionImage4Handler(ctx echo.Context) error {
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

	fileName := fmt.Sprintf("image4.%v", inputFileType)

	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.UpdateStorySEctionImage4(fileName, fileSrc); err != nil {
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

func (h *StorySectionHandler) DeleteStorySectionImage1Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.DeleteStorySectionImage1(fileName); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}
func (h *StorySectionHandler) DeleteStorySectionImage2Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.DeleteStorySectionImage2(fileName); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}

func (h *StorySectionHandler) DeleteStorySectionImage3Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.DeleteStorySectionImage3(fileName); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}

func (h *StorySectionHandler) DeleteStorySectionImage4Handler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")
	storySectionUseCase := usecase.NewStorySectionUseCase(h.database,
		storage.NewStorySectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := storySectionUseCase.DeleteStorySectionImage4(fileName); err != nil {
		ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(
		200,
		echo.Map{
			"message": "image deleted successfully",
		},
	)

}
