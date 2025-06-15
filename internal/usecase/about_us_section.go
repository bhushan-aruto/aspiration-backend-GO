package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type AboutSectionUseCase struct {
	databaseRepo repository.AboutSectionDatabaseRepo
	storageRepo  repository.AboutSectionStorageRepo
}

func NewAboutSectionUseCase(dbRepo repository.AboutSectionDatabaseRepo, storageRepo repository.AboutSectionStorageRepo) *AboutSectionUseCase {
	return &AboutSectionUseCase{
		databaseRepo: dbRepo,
		storageRepo:  storageRepo,
	}
}

func (u *AboutSectionUseCase) GetAboutSectionUseCase() (*entity.AboutUsSection, error) {
	isAboutSectionExists, err := u.databaseRepo.CheckAboutSectionExists()
	if err != nil {
		log.Println("error occurred with database while checking about section exists, Error: ", err.Error())
		return nil, errors.New("error occurred with database")
	}

	if !isAboutSectionExists {
		return nil, nil
	}
	aboutSection, err := u.databaseRepo.GetAboutSection()

	if err != nil {
		log.Println("error occurred with database while getting the about section, Error: ", err.Error())
		return nil, errors.New("error occurred with database")
	}

	return aboutSection, nil
}

func (u *AboutSectionUseCase) CreateAboutSection() error {

	aboutSection := entity.NewAboutUsSection("pending", "pending")

	if err := u.databaseRepo.CreateAboutSection(aboutSection); err != nil {
		log.Println("error occurred with database while creating the about section, Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	return nil
}

func (u *AboutSectionUseCase) UpdateAboutSectionImage1(fileName string, fileData io.ReadSeeker) error {

	uploadedImageUrl, err := u.storageRepo.UploadImage(fileName, fileData)

	if err != nil {
		log.Println("error occurred while uploading the about section image to storage, Error: ", err.Error())
		return errors.New("error occurred with storage")
	}

	aboutSection, err := u.databaseRepo.GetAboutSection()

	if err != nil {
		log.Println("error occurred with database while getting the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	aboutSection.Image1Url = uploadedImageUrl

	if err := u.databaseRepo.UpdateAboutSection(aboutSection); err != nil {
		log.Println("error occurred with database while updating the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	return nil
}

func (u *AboutSectionUseCase) UpdateAboutSectionImage2(fileName string, fileData io.ReadSeeker) error {

	uploadedImageUrl, err := u.storageRepo.UploadImage(fileName, fileData)

	if err != nil {
		log.Println("error occurred while uploading the about section image to storage, Error: ", err.Error())
		return errors.New("error occurred with storage")
	}

	aboutSection, err := u.databaseRepo.GetAboutSection()

	if err != nil {
		log.Println("error occurred with database while getting the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	aboutSection.Image2Url = uploadedImageUrl

	if err := u.databaseRepo.UpdateAboutSection(aboutSection); err != nil {
		log.Println("error occurred with database while updating the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	return nil
}

func (u *AboutSectionUseCase) DeleteAboutSectionImage1(fileName string) error {
	if err := u.storageRepo.DeleteImage(fileName); err != nil {
		log.Println("error occurred with storage while deleting the about section image1, Error: ", err.Error())
		return errors.New("error occurred with storage")
	}

	aboutSection, err := u.databaseRepo.GetAboutSection()

	if err != nil {
		log.Println("error occurred with database while getting the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	aboutSection.Image1Url = "pending"

	if err := u.databaseRepo.UpdateAboutSection(aboutSection); err != nil {
		log.Println("error occurred with database while updating the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}
	return nil
}

func (u *AboutSectionUseCase) DeleteAboutSectionImage2(fileName string) error {
	if err := u.storageRepo.DeleteImage(fileName); err != nil {
		log.Println("error occurred with storage while deleting the about section image1, Error: ", err.Error())
		return errors.New("error occurred with storage")
	}

	aboutSection, err := u.databaseRepo.GetAboutSection()

	if err != nil {
		log.Println("error occurred with database while getting the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	aboutSection.Image2Url = "pending"

	if err := u.databaseRepo.UpdateAboutSection(aboutSection); err != nil {
		log.Println("error occurred with database while updating the about section Error: ", err.Error())
		return errors.New("error occurred with database")
	}

	return nil

}
