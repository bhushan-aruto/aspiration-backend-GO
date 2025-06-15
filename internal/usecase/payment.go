package usecase

import (
	"errors"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type PaymentGateWayUseCase struct {
	paymentGatewayRepo repository.PaymentGateWayRepo
	dbRepo             repository.PaymentDatabaseRepo
}

func NewPaymentGateWayUseCase(paymentGateWayRepo repository.PaymentGateWayRepo, dbRepo repository.PaymentDatabaseRepo) *PaymentGateWayUseCase {
	return &PaymentGateWayUseCase{
		paymentGatewayRepo: paymentGateWayRepo,
		dbRepo:             dbRepo,
	}
}

func (u *PaymentGateWayUseCase) CreateOrder(coursesId []string) ([]byte, error) {

	totalPrice, err := u.dbRepo.GetCourseAmountByIds(coursesId)

	if err != nil {
		log.Println("error occurred with database: ", err.Error())
		return nil, errors.New("error occurred with database while fetching the coruses price")
	}

	resp, err := u.paymentGatewayRepo.CreateOrder(
		totalPrice,
	)
	if err != nil {
		log.Println("error  occurred while creating the order ,Error :", err.Error())
		return nil, errors.New("error occurred with razorpay order creation")
	}
	return resp, nil
}

func (u *PaymentGateWayUseCase) VerifyOrderUseCase(orderID, paymentID, razorpaySignature, secret string) bool {
	valid := u.paymentGatewayRepo.VerifyOrder(orderID, paymentID, razorpaySignature, secret)
	return valid
}

func (u *PaymentGateWayUseCase) DeleteCartUseCase(userID, courseID string) error {
	if err := u.dbRepo.DeleteCartCourseafterPayment(userID, courseID); err != nil {
		log.Println("error occurred with database while deletingng cart ", err.Error())
		return errors.New("error occured  with database")
	}
	return nil
}
