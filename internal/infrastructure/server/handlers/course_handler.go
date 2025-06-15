package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/database"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/storage"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/usecase"
	"github.com/labstack/echo/v4"
)

type CourseHandler struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	database          *database.MongoDBdatabase
	s3                *storage.S3Connection
}

func NewCourseHandler(bucketName, dirName, cloudFrontBaseUrl string, db *database.MongoDBdatabase, s3 *storage.S3Connection) *CourseHandler {
	return &CourseHandler{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		database:          db,
		s3:                s3,
	}
}

func (h *CourseHandler) UploadCourseHandler(ctx echo.Context) error {
	videoFile, err := ctx.FormFile("video")

	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while getting the video",
			},
		)
	}

	if videoFile == nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "video file is empty",
			},
		)
	}

	videoSrc, err := videoFile.Open()
	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while opening the video",
			},
		)
	}
	defer videoSrc.Close()

	thumbnailFile, err := ctx.FormFile("thumbnail")
	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while getting the thumbnail",
			},
		)
	}

	if thumbnailFile == nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "thumbnail file is empty",
			},
		)
	}

	thumbnailSrc, err := thumbnailFile.Open()
	if err != nil {
		return ctx.JSON(
			400,
			echo.Map{
				"error": "error occurred while opening the thumbnail",
			},
		)
	}
	defer thumbnailSrc.Close()

	title := ctx.FormValue("title")
	instructor := ctx.FormValue("instructor")
	description := ctx.FormValue("description")
	price, err := strconv.Atoi(ctx.FormValue("price"))
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid price",
		})
	}
	originalPrice, err := strconv.Atoi(ctx.FormValue("originalPrice"))
	if err != nil {
		return ctx.JSON(400, echo.Map{
			"error": "invalid original price",
		})
	}
	duration := ctx.FormValue("duration")
	tags := ctx.FormValue("tags")
	tagList := strings.Split(tags, ",")
	videoFileName := videoFile.Filename
	thumbnailFileName := thumbnailFile.Filename

	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := courseUseCase.UploadCourse(title,
		instructor,
		description,
		duration,
		price,
		originalPrice,
		tagList,
		thumbnailSrc,
		videoSrc,
		thumbnailFileName,
		videoFileName,
	); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to upload course: " + err.Error(),
		})
	}
	return ctx.JSON(201, echo.Map{
		"message": "course uploaded successfully",
	})

}

func (h *CourseHandler) GetAllCourseUsecase(ctx echo.Context) error {

	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	course, err := courseUseCase.GetAllCoursesUseCase()

	if err != nil {

		return ctx.JSON(500, echo.Map{
			"error": "failed to get the course: " + err.Error(),
		})
	}

	var response []models.CourseResponse
	for _, c := range course {

		courseResp := models.CourseResponse{
			Id:            c.Id,
			Title:         c.Title,
			Instructor:    c.Instructor,
			Thumbnail:     c.Thumbnail,
			Price:         c.Price,
			OriginalPrice: c.OriginalPrice,
			Duration:      c.Duration,
			Description:   c.Description,
			Tags:          c.Tags,
		}
		response = append(response, courseResp)

	}

	return ctx.JSON(200, echo.Map{
		"data": response,
	})

}

func (h *CourseHandler) DeleteCourseByIdHandler(ctx echo.Context) error {
	id := ctx.Param("id")
	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	if err := courseUseCase.DeleteCourseByIDUseCase(id); err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to delete the course: " + err.Error(),
		})
	}

	return ctx.JSON(200, echo.Map{
		"message": "course deleted succeccfully",
	})
}

func (h *CourseHandler) GetCourseByIdHandler(ctx echo.Context) error {

	id := ctx.Param("id")
	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)
	course, err := courseUseCase.GetCourseById(id)

	if err != nil {
		return ctx.JSON(500, echo.Map{
			"error": "failed to get the course: " + err.Error(),
		})
	}

	resp := models.PurchasedCourse{
		Id:            course.Id,
		Title:         course.Title,
		Instructor:    course.Instructor,
		Thumbnail:     course.Thumbnail,
		VideoURL:      course.VideoURL,
		Price:         course.Price,
		OriginalPrice: course.OriginalPrice,
		Duration:      course.Duration,
		Description:   course.Description,
		Tags:          course.Tags,
		Purchased:     true,
	}

	return ctx.JSON(200, echo.Map{
		"data": resp,
	})

}

func (h *CourseHandler) GetPurchasedCoursesByUserIdHandler(ctx echo.Context) error {
	userId := ctx.Param("user_id")
	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	course, err := courseUseCase.GetPurchasedCoursesByUserIdUseCase(userId)
	if err != nil {
		ctx.JSON(400, echo.Map{
			"error": err.Error(),
		})

	}

	var response []models.PurchasedCourse
	for _, c := range course {

		courseResp := models.PurchasedCourse{
			Id:            c.Id,
			Title:         c.Title,
			Instructor:    c.Instructor,
			Thumbnail:     c.Thumbnail,
			VideoURL:      c.VideoURL,
			Price:         c.Price,
			OriginalPrice: c.OriginalPrice,
			Duration:      c.Duration,
			Description:   c.Description,
			Tags:          c.Tags,
			Purchased:     true,
		}
		response = append(response, courseResp)

	}

	return ctx.JSON(200, echo.Map{
		"data": response,
	})

}

func (h *CourseHandler) GetAllCourseUsecaseExpectPurchasedHandler(ctx echo.Context) error {
	userID := ctx.Param("user_id")
	if userID == "" {
		return ctx.JSON(http.StatusBadRequest, echo.Map{
			"error": "user_id is required",
		})
	}
	courseUseCase := usecase.NewCourseUsecase(
		h.database,
		storage.NewCourseStorageRepo(
			h.bucketName,
			h.dirName,
			h.cloudFrontBaseUrl,
			h.s3,
		),
	)

	courses, err := courseUseCase.GetCoursesNotPurchasedByUser(userID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to get the course: " + err.Error(),
		})
	}

	if len(courses) == 0 {
		return ctx.JSON(200, echo.Map{
			"data": nil,
		})
	}

	var response []models.CourseResponse
	for _, c := range courses {
		courseResp := models.CourseResponse{
			Id:            c.Id,
			Title:         c.Title,
			Instructor:    c.Instructor,
			Thumbnail:     c.Thumbnail,
			Price:         c.Price,
			OriginalPrice: c.OriginalPrice,
			Duration:      c.Duration,
			Description:   c.Description,
			Tags:          c.Tags,
			Purchased:     false,
		}
		response = append(response, courseResp)
	}

	return ctx.JSON(200, echo.Map{
		"data": response,
	})
}
