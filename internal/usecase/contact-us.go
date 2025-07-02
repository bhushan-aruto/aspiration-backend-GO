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
	SendEnquiryCourseUsecase(req models.EamilCourseEnquiryRequest, cfg *config.Config) error
}

type ContactUs struct {
}

func NewContactUs() *ContactUs {
	return &ContactUs{}
}

func (u *ContactUs) SendEmailUseCse(req models.EmailContactRequest, cfg *config.Config) error {
	body := fmt.Sprintf(`
You have received a new contact message:

----------------------------------------
Name     : %s
Email    : %s
Type     : %s
----------------------------------------

Message:
%s
`, req.Name, req.Email, req.Type, req.Message)

	msg := []byte("From: " + cfg.FromEmail + "\r\n" +
		"To: " + cfg.ToEmail + "\r\n" +
		"Subject: New Contact Message\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", "apikey", cfg.EmailAppPassword, cfg.SmtpHost)

	return smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.FromEmail, []string{cfg.ToEmail}, msg)
}


func (u *ContactUs) GenerateWhatsAppURL(req models.WhatsAppContactRequest, cfg *config.Config) (string, error) {
	msg := fmt.Sprintf(
		`You have received a new WhatsApp inquiry:

Name    : %s
Phone   : %s
Type    : %s

Message :
%s`, req.Name, req.Phone, req.Type, req.Message)

	return fmt.Sprintf("https://wa.me/%s?text=%s", cfg.WhatspNumber, url.QueryEscape(msg)), nil
}

func (u *ContactUs) SendEnquiryCourseUsecase(req models.EamilCourseEnquiryRequest, cfg *config.Config) error {

	body := fmt.Sprintf(
		"📘 Course Enquiry Details 📘\n\n"+
			"Course Name   : %s\n"+
			"Full Name     : %s\n"+
			"Email Address : %s\n"+
			"Phone Number  : %s\n"+
			"Company Name  : %s\n"+
			"Enquiry       : %s\n",
		req.CourseName,
		req.Name,
		req.Email,
		req.Phone,
		req.CompanyName,
		req.AboutEnquiry,
	)

	msg := []byte("From: " + cfg.FromEmail + "\r\n" +
		"To: " + cfg.ToEmail + "\r\n" +
		"Reply-To: " + req.Email + "\r\n" +
		"Subject: 📩 New Course Enquiry - " + req.CourseName + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
		"\r\n" +
		body + "\r\n")

	auth := smtp.PlainAuth("", "apikey", cfg.EmailAppPassword, cfg.SmtpHost)

	return smtp.SendMail(cfg.SmtpHost+":"+cfg.SmtpPort, auth, cfg.FromEmail, []string{cfg.ToEmail}, msg)
}
