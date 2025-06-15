package entity

type EventGallary struct {
	ID       string `json:"id" bson:"_id,omitempty"`
	FileName string `json:"file_name" bson:"file_name"`
	ImageUrl string `json:"image_url" bson:"image_url"`
}

func NewEventGallery(id, fileName string, imageUrl string) *EventGallary {
	return &EventGallary{
		ID:       id,
		FileName: fileName,
		ImageUrl: imageUrl,
	}
}
