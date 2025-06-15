package repository

type PaymentGateWayRepo interface {
	CreateOrder(amount int32) ([]byte, error)
	VerifyOrder(orderID, paymentID, razorpaySignature, secret string) bool
}
