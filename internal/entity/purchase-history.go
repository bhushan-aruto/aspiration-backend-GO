package entity

type PurchaseHistory struct {
	ID            string      `json:"id" bson:"_id"`
	UserID        string      `json:"user_id" bson:"user_id"`
	CourseID      string      `json:"course_id" bson:"course_id"`
	Date          string      `json:"date" bson:"date"`
	Amount        int32       `json:"amount" bson:"amount"`
	PaymentMethod string      `json:"payment_method" bson:"payment_method"`
	Course        interface{} `json:"course,omitempty" bson:"course,omitempty"`
}
