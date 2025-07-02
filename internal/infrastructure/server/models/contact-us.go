package models

type EmailContactRequest struct {
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required,email"`
	Type    string `json:"type" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type WhatsAppContactRequest struct {
	Name    string `json:"name" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Type    string `json:"type" validate:"required"`
	Message string `json:"message" validate:"required"`
}

type EamilCourseEnquiryRequest struct {
	CourseName   string `json:"course_name" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Email        string `json:"email" validate:"required"`
	Phone        string `json:"phone" validate:"required"`
	CompanyName  string `json:"company_name" validate:"required"`
	AboutEnquiry string `json:"about_enquiry" validate:"required"`
}
