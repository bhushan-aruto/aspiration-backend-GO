package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
)

type CourseUsecase struct {
	database repository.CourseDatabaseRepo
	storage  repository.CourseStorageRepo
}

func NewCourseUsecase(db repository.CourseDatabaseRepo, storage repository.CourseStorageRepo) *CourseUsecase {
	return &CourseUsecase{database: db, storage: storage}
}

func (u *CourseUsecase) UploadCourse(title, instructor, desc, duration string, price, originalPrice int, tags []string, thumbFile, videoFile io.ReadSeeker, thumbName, videoName string) error {
	newId := utils.NewId()

	videoURL, err := u.storage.UploadFile(videoName, videoFile)
	if err != nil {
		log.Println("error occured with storage while uploding video to s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	thumbURL, err := u.storage.UploadFile(thumbName, thumbFile)
	if err != nil {
		log.Println("error occured with storage while uploding thumbnail to s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	course := entity.NewCourse(newId, title, instructor, thumbURL, videoURL, originalPrice, price, duration, desc, tags)
	if err := u.database.AddCourse(course); err != nil {
		log.Println("error occured with database while adding new course  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil
}

func (u *CourseUsecase) GetAllCoursesUseCase() ([]*entity.Course, error) {
	course, err := u.database.GetAllTheCourses()
	if err != nil {
		log.Println("error occured with database while getting a courses ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}
	return course, nil
}

func (u *CourseUsecase) DeleteCourseByIDUseCase(id string) error {
	course, err := u.database.GetCourseByID(id)
	if err != nil {
		log.Println("error occured with database while getting the course ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	videoFile := utils.ExtractFileNameFromURL(course.VideoURL)
	thumbFile := utils.ExtractFileNameFromURL(course.Thumbnail)

	if err := u.storage.DeleteFile(videoFile); err != nil {
		log.Println("error occured with storage while deleting video from s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	if err := u.storage.DeleteFile(thumbFile); err != nil {
		log.Println("error occured with storage while deleting thumbnail from s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	if err := u.database.DeleteCourseById(id); err != nil {
		log.Println("error occured with database while deleting the course ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}

func (u *CourseUsecase) GetCourseById(id string) (*entity.Course, error) {
	course, err := u.database.GetCourseByID(id)
	if err != nil {
		log.Println("error occured with database while getting the course ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}

	return course, nil
}

func (u *CourseUsecase) GetPurchasedCoursesByUserIdUseCase(userID string) ([]*entity.Course, error) {
	courses, err := u.database.GetPurchasedCoursesByUserId(userID)
	if err != nil {

		if err.Error() == "mongo: no documents in result" {
			return []*entity.Course{}, nil
		}
		log.Println("Database error:", err.Error())
		return nil, errors.New("internal server error")
	}
	return courses, nil
}

func (u *CourseUsecase) GetCoursesNotPurchasedByUser(userID string) ([]*entity.Course, error) {
	courses, err := u.database.GetCoursesExcludingUserPurchased(userID)
	if err != nil {
		log.Println("error occured with database while getting a courses ,Error :", err.Error())
		return nil, errors.New("failed to get filtered courses in database")
	}
	return courses, nil
}
