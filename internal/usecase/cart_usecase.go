package usecase

import (
	"errors"
	"log"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type CartUsecase struct {
	database repository.CartDatabaseRepo
}

func NewCartUseCase(db repository.CartDatabaseRepo) *CartUsecase {
	return &CartUsecase{
		database: db,
	}
}

func (u *CartUsecase) AddCartItemUseCase(item *entity.CartItem) error {
	if strings.TrimSpace(item.UserID) == "" || strings.TrimSpace(item.CourseID) == "" {
		return errors.New("all the fields are required")
	}

	if err := u.database.AddToCart(item); err != nil {
		log.Println("error occured with database while adding cart  ,Error :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}

func (u *CartUsecase) GetCArtItemSUseCase(UserID string) ([]*entity.Course, error) {
	courses, err := u.database.GetCartCourse(UserID)
	if err != nil {
		log.Println("error occured with database while getting carted items  ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}

	return courses, nil
}

func (u *CartUsecase) DeletecartItem(userID, courseID string) error {
	if err := u.database.DeleteCartCourse(userID, courseID); err != nil {
		log.Println("error occured with database while deleteing the cart item ,Error :", err.Error())
		return errors.New("error occured with the database")
	}
	return nil
}
