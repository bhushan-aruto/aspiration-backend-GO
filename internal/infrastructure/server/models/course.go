package models

type CourseResponse struct {
	Id            string   `json:"id" bson:"_id,omitempty"`
	Title         string   `json:"title" bson:"title"`
	Instructor    string   `json:"instructor" bson:"instructor"`
	Thumbnail     string   `json:"thumbnail" bson:"thumbnail"`
	Price         int      `json:"price" bson:"price"`
	OriginalPrice int      `json:"originalPrice" bson:"original_price"`
	Duration      string   `json:"duration" bson:"duration"`
	Description   string   `json:"description" bson:"description"`
	Tags          []string `json:"tags" bson:"tags"`
	Purchased     bool     `json:"purchased" bson:"purchased"`
}
