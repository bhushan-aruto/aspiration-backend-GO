package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type StorySectionUseCase struct {
	databaseRepo repository.StorySectionDatabaseRepo
	storageRepo  repository.StorySectionStorageRepo
}

func NewStorySectionUseCase(db repository.StorySectionDatabaseRepo, st repository.StorySectionStorageRepo) *StorySectionUseCase {
	return &StorySectionUseCase{
		databaseRepo: db,
		storageRepo:  st,
	}

}

func (u *StorySectionUseCase) GetStorySEctionUseCase() (*entity.StorySection, error) {
	userExist, err := u.databaseRepo.CheckStorySectionExists()
	if err != nil {
		log.Println("error occured with the database while checking the story section ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}
	if !userExist {
		return nil, nil
	}
	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database while getting the story section ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}

	return storySection, err
}

func (u *StorySectionUseCase) CreateStorySectionUseCase() error {
	storySection := entity.NewStorySection("pending", "pending", "pending", "pending")
	if err := u.databaseRepo.CreateStorySection(storySection); err != nil {
		log.Println("error occured with database while creating the story section .Error :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *StorySectionUseCase) UpdateStorySEctionImage1(filename string, fileData io.ReadSeeker) error {
	uploadedImageurl, err := u.storageRepo.UploadImage(filename, fileData)
	if err != nil {
		log.Println("error occured with storage  while uploding the  story section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage ")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database  while getting story section  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	storySection.Image1Url = uploadedImageurl

	if err = u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the  story section ,Erorr :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *StorySectionUseCase) UpdateStorySEctionImage2(filename string, fileData io.ReadSeeker) error {
	uploadedImageurl, err := u.storageRepo.UploadImage(filename, fileData)
	if err != nil {
		log.Println("error occured with storage  while uploding the  story section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage ")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database  while getting story section  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	storySection.Image2Url = uploadedImageurl

	if err = u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the  story section ,Erorr :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *StorySectionUseCase) UpdateStorySEctionImage3(filename string, fileData io.ReadSeeker) error {
	uploadedImageurl, err := u.storageRepo.UploadImage(filename, fileData)
	if err != nil {
		log.Println("error occured with storage  while uploding the  story section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage ")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database  while getting story section  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	storySection.Image3url = uploadedImageurl

	if err = u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the  story section ,Erorr :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *StorySectionUseCase) UpdateStorySEctionImage4(filename string, fileData io.ReadSeeker) error {
	uploadedImageurl, err := u.storageRepo.UploadImage(filename, fileData)
	if err != nil {
		log.Println("error occured with storage  while uploding the  story section image to the s3,Error :", err.Error())
		return errors.New("error occured with the storage ")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database  while getting story section  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	storySection.Image4url = uploadedImageurl

	if err = u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the  story section ,Erorr :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *StorySectionUseCase) DeleteStorySectionImage1(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with the storage while deleting the story section image 1 ,Error :", err.Error())
		return errors.New("error occurred with the storage")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database while getting story section  ,Error :", err.Error())
		return errors.New("error occured with database")
	}
	storySection.Image1Url = "pending"

	if err := u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the story section ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}

func (u *StorySectionUseCase) DeleteStorySectionImage2(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with the storage while deleting the story section image 2 ,Error :", err.Error())
		return errors.New("error occurred with the storage")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database while getting story section  ,Error :", err.Error())
		return errors.New("error occured with database")
	}
	storySection.Image2Url = "pending"

	if err := u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the story section ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}

func (u *StorySectionUseCase) DeleteStorySectionImage3(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with the storage while deleting the story section image 3 ,Error :", err.Error())
		return errors.New("error occurred with the storage")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database while getting story section  ,Error :", err.Error())
		return errors.New("error occured with database")
	}
	storySection.Image3url = "pending"

	if err := u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the story section ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}

func (u *StorySectionUseCase) DeleteStorySectionImage4(filename string) error {
	if err := u.storageRepo.DeleteImage(filename); err != nil {
		log.Println("error occured with the storage while deleting the story section image 4 ,Error :", err.Error())
		return errors.New("error occurred with the storage")
	}

	storySection, err := u.databaseRepo.GetstorySection()
	if err != nil {
		log.Println("error occured with the database while getting story section  ,Error :", err.Error())
		return errors.New("error occured with database")
	}
	storySection.Image4url = "pending"

	if err := u.databaseRepo.UpdateStorySection(storySection); err != nil {
		log.Println("error occured with the database while updating the story section ,Error :", err.Error())
		return errors.New("error occured with the database")
	}

	return nil

}
