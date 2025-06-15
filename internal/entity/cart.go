package entity

type CartItem struct {
	UserID   string `json:"user_id" bson:"user_id"`
	CourseID string `json:"course_id" bson:"course_id"`
}


