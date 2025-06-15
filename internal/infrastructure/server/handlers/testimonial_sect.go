package handlers

import (
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type TestimonialHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewTestimonialsHandler(bucketName, dirName, cloudFrontBaseUrl string, db *database.MongoDBdatabase, s3 *storage.S3Connection) *TestimonialHandler {
	return &TestimonialHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          db,
		s3:                s3,
	}
}

func (h *TestimonialHandler) AddTestimonialsHandler(ctx echo.Context) error {
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

	name := ctx.FormValue("name")
	role := ctx.FormValue("role")
	company := ctx.FormValue("company")
	review := ctx.FormValue("review")
	rating := ctx.FormValue("rating")
	fileName := file.Filename

	TestimonialsUsecase := usecase.NewTestimonialUseCase(
		h.database,
		storage.NewTestimonialSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	if err := TestimonialsUsecase.AddTestimonial(name, role, company, review, rating, fileName, fileSrc); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "testimonial uploaded  succesfully"})
}

func (h *TestimonialHandler) GetVerifiedtestimonials(ctx echo.Context) error {
	TestimonialsUsecase := usecase.NewTestimonialUseCase(
		h.database,
		storage.NewTestimonialSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	testimonials, err := TestimonialsUsecase.GetVerifiedTestimonialsUseCase()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})

	}

	return ctx.JSON(200, echo.Map{
		"data": testimonials,
	})

}

func (h *TestimonialHandler) GetUnVerifiedtestimonials(ctx echo.Context) error {
	TestimonialsUsecase := usecase.NewTestimonialUseCase(
		h.database,
		storage.NewTestimonialSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	testimonials, err := TestimonialsUsecase.GetUnVerifiedTestimonialsUseCase()
	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})

	}

	return ctx.JSON(200, echo.Map{
		"data": testimonials,
	})

}

func (h *TestimonialHandler) ApproveTestimonials(ctx echo.Context) error {
	id := ctx.Param("id")
	TestimonialsUsecase := usecase.NewTestimonialUseCase(
		h.database,
		storage.NewTestimonialSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := TestimonialsUsecase.VerifyTestimonialUseCase(id); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})

	}

	return ctx.JSON(200, echo.Map{
		"message": "testimonial aproved succesfully"})

}

func (h *TestimonialHandler) DeleteTestimonial(ctx echo.Context) error {
	fileName := ctx.Param("fileName")

	TestimonialsUsecase := usecase.NewTestimonialUseCase(
		h.database,
		storage.NewTestimonialSectionRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := TestimonialsUsecase.DeleteTestimonialByFileNameUseCase(fileName); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "testimonial deleted  succesfully"})
}
