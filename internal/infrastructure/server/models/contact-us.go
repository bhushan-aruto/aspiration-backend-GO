package models

type EmailContactRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Message string `json:"message" validate:"required"`
}

type WhatsAppContactRequest struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Message string `json:"message" validate:"required"`
}
