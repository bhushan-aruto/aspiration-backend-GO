package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
)

type EventGallaryUseCase struct {
	databaserepo repository.EventGallerySectionDatabaseRepo
	storage      repository.EventGallaryStorageRepo
}

func NewEventGalleryUseCase(db repository.EventGallerySectionDatabaseRepo, st repository.EventGallaryStorageRepo) *EventGallaryUseCase {
	return &EventGallaryUseCase{
		databaserepo: db,
		storage:      st,
	}
}

func (u *EventGallaryUseCase) UploadEventImage(fileName string, file io.ReadSeeker) error {
	imageId := utils.NewId()

	uploadedurl, err := u.storage.UploadImage(fileName, file)
	if err != nil {
		log.Println("error occured with storage  while uploding the  event section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage ")
	}

	image := entity.NewEventGallery(imageId, fileName, uploadedurl)
	if err = u.databaserepo.AddEventGalleryImage(image); err != nil {
		log.Println("error occured with database  while adding  event section image to the s3 ,Error :", err.Error())
		return errors.New("error occured with the databse ")

	}
	return nil
}

func (u *EventGallaryUseCase) GetEventSectionAllImages() ([]*entity.EventGallary, error) {
	images, err := u.databaserepo.GetAllEventGalleryImages()
	if err != nil {
		log.Println("error occured with database while getting the  event gallery images ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}
	return images, nil
}

func (u *EventGallaryUseCase) DeleteImageByFileName(fileName string) error {
	if err := u.storage.DeleteImage(fileName); err != nil {
		log.Println("error occured with storage  while deleting  the  event section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	if err := u.databaserepo.DeleteEventImagebyFileName(fileName); err != nil {
		log.Println("error occured with the databse while deleteing the the event   section image to the s3 ,Error :", err.Error())
		return errors.New("error occured with the databse")

	}
	return nil
}
