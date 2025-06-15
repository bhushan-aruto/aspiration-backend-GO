package handlers

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type BlogsHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewBlogsHandler(bucketName, dirName, cloudFrontBaseUrl string, db *database.MongoDBdatabase, s3 *storage.S3Connection) *BlogsHandler {
	return &BlogsHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          db,
		s3:                s3,
	}
}

func (h *BlogsHandler) UploadBloghandler(ctx echo.Context) error {
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
	title := ctx.FormValue("title")
	description := ctx.FormValue("description")
	content := ctx.FormValue("content")
	date := ctx.FormValue("date")

	fileName := file.Filename

	blogUseCase := usecase.NewBlogsUseCase(h.database,
		storage.NewBlogsSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := blogUseCase.UploadBlog(title, description, content, date, fileName, fileSrc); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "blog uploaded  succesfully"})

}

func (h *BlogsHandler) GetAllBlogsHandler(ctx echo.Context) error {
	blogUseCase := usecase.NewBlogsUseCase(h.database,
		storage.NewBlogsSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	blogs, err := blogUseCase.GetAllBlogsUseCase()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"data": blogs,
	})

}

func (h *BlogsHandler) GetBlogHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	blogUseCase := usecase.NewBlogsUseCase(h.database,
		storage.NewBlogsSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	blog, err := blogUseCase.GetBlogByIdUseCase(id)
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(200, echo.Map{
		"data": blog,
	})
}

func (h *BlogsHandler) DeletBlogsHandler(ctx echo.Context) error {
	fileName := ctx.Param("fileName")
	blogUseCase := usecase.NewBlogsUseCase(h.database,
		storage.NewBlogsSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := blogUseCase.DeleteBlogUsecase(fileName); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "blogs deleted successfully",
	})
}
