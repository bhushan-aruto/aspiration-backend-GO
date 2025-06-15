package entity

import "time"

type Blog struct {
	Id         string    `json:"id" bson:"_id,omitempty"`
	Title      string    `json:"title" bson:"title"`
	Descrpiton string    `json:"description" bson:"description"`
	ImageUrl   string    `json:"image_url" bson:"image_url"`
	Content    string    `json:"content" bson:"content"`
	Date       string    `json:"date" bson:"date"`
	FileName   string    `json:"file_name" bson:"file_name"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

func NewBlog(id, title, description, imageurl, content, date, filename string) *Blog {
	return &Blog{
		Id:         id,
		Title:      title,
		Descrpiton: description,
		ImageUrl:   imageurl,
		Content:    content,
		Date:       date,
		FileName:   filename,
		CreatedAt:  time.Now(),
	}
}
