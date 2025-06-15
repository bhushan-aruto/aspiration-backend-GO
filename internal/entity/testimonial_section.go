package entity

import "time"

type Testimonial struct {
	Id        string    `json:"id" bson:"_id,omitempty"`
	Name      string    `json:"name" bson:"name"`
	Role      string    `json:"role" bson:"role"`
	Company   string    `json:"company" bson:"company"`
	ImageUrl  string    `json:"image_url" bson:"image_url"`
	Review    string    `json:"review" bson:"review"`
	Rating    string    `json:"rating" bson:"rating"`
	FileName  string    `json:"file_name" bson:"file_name"`
	Verified  bool      `json:"verified" bson:"verified"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}

func NewTestimonials(id, name, role, company, imageurl, review, fileName, rating string, verified bool) *Testimonial {
	return &Testimonial{
		Id:        id,
		Name:      name,
		Role:      role,
		Company:   company,
		ImageUrl:  imageurl,
		Review:    review,
		Rating:    rating,
		FileName:  fileName,
		Verified:  verified,
		CreatedAt: time.Now(),
	}

}
