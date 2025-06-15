package entity

type MyLearning struct {
	UserId    string   `json:"user_id" bson:"user_id"`
	CourseIds []string `json:"cours_ids" bson:"course_ids"`
}

func NewMylearning(userID string, courseIds []string) *MyLearning {
	return &MyLearning{
		UserId:    userID,
		CourseIds: courseIds,
	}
}
