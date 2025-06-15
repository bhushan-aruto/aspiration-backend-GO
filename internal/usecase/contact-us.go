package usecase

import (
	"fmt"
	"net/smtp"
	"net/url"

	"github.com/bhushan-aruto/aspiration-matters-backend/config"
	"github.com/bhushan-aruto/aspiration-matters-backend/internal/infrastructure/server/models"
)

type ContactUsecase interface {
	SendEmailUseCse(req models.EmailContactRequest, cfg *config.Config) error
	GenerateWhatsAppURL(req models.WhatsAppContactRequest, cfg *config.Config) (string, error)
}

type ContactUs struct{
}

func NewContactUs() *ContactUs {
	return &ContactUs{}
}

func (u *ContactUs) SendEmailUseCse(req models.EmailContactRequest, cfg *config.Config) error {

	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", req.Name, req.Email, req.Message)

	msg := []byte("To: " + cfg.ToEmail + "\r\n" +
		"Subject: New Contact Message\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", cfg.FromEmail, cfg.EmailAppPassword, cfg.SmtpHost)

	return smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.FromEmail, []string{cfg.ToEmail}, msg)
}

func (u *ContactUs) GenerateWhatsAppURL(req models.WhatsAppContactRequest, cfg *config.Config) (string, error) {

	msg := fmt.Sprintf("Name: %s\nPhone: %s\nMessage: %s", req.Name, req.Phone, req.Message)
	return fmt.Sprintf("https://wa.me/%s?text=%s", cfg.WhatspNumber, url.QueryEscape(msg)), nil
}
