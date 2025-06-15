package usecase

import (
	"errors"
	"log"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
)

type PurchaseHistory struct {
	PurchaseHistory repository.PurchaseHistoryRepo
}

func NewPurchaseHistoryUseCase(purchaseHistory repository.PurchaseHistoryRepo) *PurchaseHistory {
	return &PurchaseHistory{
		PurchaseHistory: purchaseHistory,
	}
}

func (u *PurchaseHistory) GetPurchaseHistoryByUserIdUseCAse(userID string) ([]*entity.PurchaseHistory, error) {

	purchaseHistory, err := u.PurchaseHistory.GetPurchaseHistoryByUserId(userID)

	if err != nil {
		log.Println("error occured with the database while getting the  purchase history  Error:", err.Error())
		return nil, errors.New("error occured with the database")
	}
	return purchaseHistory, nil
}
