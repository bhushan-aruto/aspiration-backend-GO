package usecase

import (
	"errors"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase struct {
	AdminRepo repository.AdminDataBaseRepo
}

func NewAdminUseCase(adminRepo repository.AdminDataBaseRepo) *AdminUseCase {
	return &AdminUseCase{
		AdminRepo: adminRepo,
	}
}

func (a *AdminUseCase) AdminSignUp(adminName, email, phonNumber, password string) (*entity.Admin, error) {
	if strings.TrimSpace(adminName) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(phonNumber) == "" || strings.TrimSpace(password) == "" {
		return nil, errors.New("all feilds are required ")
	}

	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	existUser := a.AdminRepo.CheckAdminByEmail(email)
	if existUser != nil {
		return nil, errors.New("email already registered")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}
	id := utils.NewId()

	if !utils.IsValidPhonenumber(phonNumber) {
		return nil, errors.New("invalid phone number.please include a valid country code ")
	}
	admin := entity.NewAdmin(id, adminName, email, phonNumber, hashedPassword)

	if err = a.AdminRepo.SaveAdmin(admin); err != nil {
		return nil, err
	}

	return admin, nil
}

func (a *AdminUseCase) AdminLogin(email, password string) (string, error) {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return "", errors.New("all the fields are required")
	}

	admin := a.AdminRepo.CheckAdminByEmail(email)
	if admin == nil {
		return "", errors.New("no admin found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return "", errors.New(" password does not match")
	}

	token, err := utils.GenerateJWT(admin.Id)
	if err != nil {
		return "", err
	}

	return token, nil

}
