package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type PhilosopyUseCase struct {
	databaseRepo repository.PhilosopySectionDatabaseRepo
	storageRepo  repository.PhilosopySectionStorageRepo
}

func NewPhilosopyUseCase(dbaseRepo repository.PhilosopySectionDatabaseRepo, storagerepo repository.PhilosopySectionStorageRepo) *PhilosopyUseCase {
	return &PhilosopyUseCase{
		databaseRepo: dbaseRepo,
		storageRepo:  storagerepo,
	}
}

func (u *PhilosopyUseCase) GetPhilosopySectionUseCase() (*entity.PhilosopySection, error) {
	isAboutSectionExists, err := u.databaseRepo.CheckPhilosopySectionExists()
	if err != nil {
		log.Println("error occurred with database while checking about section exists, Error: ", err.Error())
		return nil, errors.New("error occurred with database")
	}

	if !isAboutSectionExists {
		return nil, nil
	}
	aboutSection, err := u.databaseRepo.GetPhilospySection()

	if err != nil {
		log.Println("error occurred with database while getting the philosopy section, Error: ", err.Error())
		return nil, errors.New("error occurred with database")
	}
	return aboutSection, nil
}

func (u *PhilosopyUseCase) CreatePhilosopySectionUseCase() error {
	philosopySection := entity.NewPhilosopySection("pending", "pending")
	if err := u.databaseRepo.CreatePhilosopySection(philosopySection); err != nil {
		log.Println("error occured with database while creating the philosopy section ,Error:", err.Error())
		return errors.New("error occurred with database")
	}
	return nil
}

func (u *PhilosopyUseCase) UpdatePhilospySectionImage1(fileName string, fileData io.ReadSeeker) error {
	uploadImageUrl, err := u.storageRepo.UploadImage(fileName, fileData)
	if err != nil {
		log.Println("error occured  while uploading teh philosopy section image to the storage")
		return errors.New("error occured with the storage")
	}

	philosopySection, err := u.databaseRepo.GetPhilospySection()
	if err != nil {
		log.Println("error occured  with database while getting the philosopy section ,Error:", err.Error())
		return errors.New("error occured with the database")
	}

	philosopySection.Image1Url = uploadImageUrl

	if err = u.databaseRepo.UpdatePhilosopySection(philosopySection); err != nil {
		log.Println("error occured with the database while updating the  philospysection ,Error :", err.Error())
		return errors.New("error cocured with the database")
	}

	return nil

}

func (u *PhilosopyUseCase) UpdatePhilospySectionImage2(fileName string, fileData io.ReadSeeker) error {
	uploadImageUrl, err := u.storageRepo.UploadImage(fileName, fileData)
	if err != nil {
		log.Println("error occured  while uploading teh philosopy section image to the storage")
		return errors.New("error occured with the storage")
	}

	philosopySection, err := u.databaseRepo.GetPhilospySection()
	if err != nil {
		log.Println("error occured  with database while getting the philosopy section ,Error:", err.Error())
		return errors.New("error occured with the database")
	}

	philosopySection.Image2Url = uploadImageUrl

	if err = u.databaseRepo.UpdatePhilosopySection(philosopySection); err != nil {
		log.Println("error occured with the database while updating the  philospysection ,Error :", err.Error())
		return errors.New("error cocured with the database")
	}

	return nil

}

func (u *PhilosopyUseCase) DeletePhilosopySectionImage1(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with storage while deleting philosopy section image ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	philosopySection, err := u.databaseRepo.GetPhilospySection()
	if err != nil {
		log.Println("error occured with the database while getting the database philosopy section")
		return errors.New("error occured with database")
	}

	philosopySection.Image1Url = "pending"

	if err := u.databaseRepo.UpdatePhilosopySection(philosopySection); err != nil {
		log.Println("error occured with database while updating the philosopy section")
		return errors.New("error occured with the database")
	}
	return nil
}
func (u *PhilosopyUseCase) DeletePhilosopySectionImage2(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with storage while deleting philosopy section image ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	philosopySection, err := u.databaseRepo.GetPhilospySection()
	if err != nil {
		log.Println("error occured with the database while getting the database philosopy section")
		return errors.New("error occured with database")
	}

	philosopySection.Image2Url = "pending"

	if err := u.databaseRepo.UpdatePhilosopySection(philosopySection); err != nil {
		log.Println("error occured with database while updating the philosopy section")
		return errors.New("error occured with the database")
	}
	return nil
}
