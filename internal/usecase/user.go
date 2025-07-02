package usecase

import (
	"errors"
	"log"
	"strings"

	"github.com/bhushan-aruto/aspiration-matters-backend/internal/entity"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/repository"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepo repository.UserDatabaseRepo
}

func NewUserUseCase(userRepo repository.UserDatabaseRepo) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (u *UserUseCase) UserSignUp(userName, email, password, phoneNumber string) (*entity.User, error) {

	if strings.TrimSpace(userName) == "" || strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" || strings.TrimSpace(phoneNumber) == "" {
		return nil, errors.New("all fields are required")
	}
	if !utils.IsValidEmail(email) {
		return nil, errors.New("invalid email format")
	}
	existUser, err := u.userRepo.CheckUserByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if existUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	id := utils.NewId()

	user := entity.NewUser(id, userName, email, hashedPassword, phoneNumber)

	err = u.userRepo.SaveUser(user)

	if err != nil {
		log.Println("error occureed with database while save teh user Error:", err.Error())
		return nil, err
	}
	learning := entity.NewMylearning(id, []string{})
	if err := u.userRepo.MyLearningCollection(learning); err != nil {
		log.Println("error occureed with database while creating the my leaning Error:", err.Error())
		return nil, errors.New("error occured with the database")

	}
	return user, nil
}

func (u *UserUseCase) SendOTP(email, password string) (string, error) {

	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return "", errors.New("all fields are required")
	}

	user, err := u.userRepo.CheckUserByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("password does not match")
	}

	otp, expiry := utils.GenerateOTPWithExpiry()
	err = u.userRepo.StoreOTP(email, otp, expiry)
	if err != nil {
		return "", err
	}

	return otp, nil
}

func (u *UserUseCase) VerifyOTP(email, otp string) (bool, error) {
	return u.userRepo.VerifyOTP(email, otp)
}

func (u *UserUseCase) UserLogin(email, password string) (string, error) {
	if strings.TrimSpace(email) == "" || strings.TrimSpace(password) == "" {
		return "", errors.New("all the fields are required")
	}
	user, err := u.userRepo.CheckUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		log.Println("Password mismatch error:", err)
		return "", errors.New("password does not match")
	}

	token, err := utils.GenerateJWT(user.Id)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (u *UserUseCase) GetUserById(id string) (*entity.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		log.Println("error occured with database while getting the user ,Error :", err.Error())
		return nil, errors.New("error occured with the database")
	}
	return user, nil
}

func (u *UserUseCase) SendForgotPasswordOTP(email string) (string, error) {
	if strings.TrimSpace(email) == "" {
		return "", errors.New("email is required")
	}

	_, err := u.userRepo.CheckUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	otp, expiry := utils.GenerateOTPWithExpiry()
	err = u.userRepo.StoreOTP(email, otp, expiry)
	if err != nil {
		return "", err
	}
	return otp, nil
}

func (u *UserUseCase) ResetPassword(email, otp, Password string) error {

	valid, err := u.userRepo.VerifyOTP(email, otp)

	if err != nil || !valid {
		return errors.New("invalid or expired OTP")
	}
	if err := u.userRepo.UpdatePassword(email, Password); err != nil {
		return err
	}
	return nil
}
