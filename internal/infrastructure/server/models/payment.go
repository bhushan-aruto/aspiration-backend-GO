package models

type CreatePaymentOrderRequest struct {
	CoursesId []string `json:"courses_id"`
}

type VerifyPaymentRequest struct {
	RazorpayOrderID   string   `json:"razorpay_order_id"`
	RazorpayPaymentID string   `json:"razorpay_payment_id"`
	RazorpaySignature string   `json:"razorpay_signature"`
	UserID            string   `json:"user_id"`
	CourseIDs         []string `json:"course_ids"`
}
