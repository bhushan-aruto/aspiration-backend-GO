package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
)

type TestimonialSectionUseCase struct {
	databaseRepo repository.TestimonialDatabaseRepo
	storageRepo  repository.TestimonialStorageRepo
}

func NewTestimonialUseCase(dbRepo repository.TestimonialDatabaseRepo, storageRepo repository.TestimonialStorageRepo) *TestimonialSectionUseCase {
	return &TestimonialSectionUseCase{
		databaseRepo: dbRepo,
		storageRepo:  storageRepo,
	}
}

func (u *TestimonialSectionUseCase) AddTestimonial(name, role, company, review string, rating, fileName string, file io.ReadSeeker) error {
	id := utils.NewId()
	imageUrl, err := u.storageRepo.UploadImage(fileName, file)
	if err != nil {
		log.Println("error occured with storage while uploding iamge to s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	testimonial := entity.NewTestimonials(id, name, role, company, imageUrl, review, fileName, rating, false)

	if err := u.databaseRepo.Addtestimonials(testimonial); err != nil {
		log.Println("error occured with database while adding testimonials ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}

func (u *TestimonialSectionUseCase) GetVerifiedTestimonialsUseCase() ([]*entity.Testimonial, error) {
	testimonials, err := u.databaseRepo.GetVerifiedTestimonials()
	if err != nil {
		log.Println("error occured with dataabse while getting verified testimonials ,Error :", err.Error())
		return nil, errors.New("error occured with the database")

	}
	return testimonials, nil
}

func (u *TestimonialSectionUseCase) GetUnVerifiedTestimonialsUseCase() ([]*entity.Testimonial, error) {
	testimonials, err := u.databaseRepo.GetUnverifiedTestimonials()
	if err != nil {
		log.Println("error occured with dataabse while getting unverified testimonials ,Error :", err.Error())
		return nil, errors.New("error occured with the database")

	}
	return testimonials, nil
}

func (u *TestimonialSectionUseCase) VerifyTestimonialUseCase(id string) error {
	if err := u.databaseRepo.VerifyTestimonial(id); err != nil {
		log.Println("error occured with dataabse while verifieing testimonials ,Error :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *TestimonialSectionUseCase) DeleteTestimonialByFileNameUseCase(fileName string) error {
	if err := u.storageRepo.DeleteImage(fileName); err != nil {
		log.Println("error occured with storage while deleteing image from s3,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	if err := u.databaseRepo.DeleteTestimonialByFileName(fileName); err != nil {
		log.Println("error occured with dataabse whiledeleteing the testimonial ,Error :", err.Error())
		return errors.New("error occured with the database")

	}

	return nil
}
