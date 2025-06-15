package usecase

import (
	"errors"
	"io"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
)

type BlogsUseCase struct {
	databaserepo repository.BlogsSectionDatabaseRepo
	storage      repository.BlogStorageRepo
}

func NewBlogsUseCase(db repository.BlogsSectionDatabaseRepo, st repository.BlogStorageRepo) *BlogsUseCase {
	return &BlogsUseCase{
		databaserepo: db,
		storage:      st,
	}
}

func (u *BlogsUseCase) UploadBlog(title, description, content, date, fileName string, file io.ReadSeeker) error {

	newId := utils.NewId()

	imageUrl, err := u.storage.UploadImage(fileName, file)
	if err != nil {
		log.Println("error occured with storage while uploding iamge to s3 ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}

	blog := entity.NewBlog(newId, title, description, imageUrl, content, date, fileName)
	if err := u.databaserepo.AddBlog(blog); err != nil {
		log.Println("error occured with database while adding new blog ,Error :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *BlogsUseCase) GetAllBlogsUseCase() ([]*entity.Blog, error) {
	blogs, err := u.databaserepo.GetAllBlogs()
	if err != nil {
		log.Println("error occured with database while getting a blogs ,Error :", err.Error())
		return nil, errors.New("error occured with the database")

	}

	return blogs, nil

}

func (u *BlogsUseCase) DeleteBlogUsecase(fileName string) error {
	if err := u.storage.DeleteImage(fileName); err != nil {
		log.Println("error occured with storage  while deleting  the  event section image to the s3  ,Error :", err.Error())
		return errors.New("error occured with the storage")
	}
	if err := u.databaserepo.DeleteBlogByFileName(fileName); err != nil {
		log.Println("error occured with the databse while deleting the the blog ,Error :", err.Error())
		return errors.New("error occured with the databse")

	}

	return nil
}

func (u *BlogsUseCase) GetBlogByIdUseCase(id string) (*entity.Blog, error) {
	blog, err := u.databaserepo.GetBlogById(id)
	if err != nil {
		log.Println("error occured with the databse while getting the the blog ,Error :", err.Error())
		return nil, errors.New("error occured with the databse")

	}
	return blog, err
}
