package entity

type PhilosopySection struct {
	Image1Url string `json:"image1_url" bson:"image1_url"`
	Image2Url string `json:"image2_url" bson:"image2_url"`
}

func NewPhilosopySection(image1Url, image2Url string) *PhilosopySection {
	return &PhilosopySection{
		Image1Url: image1Url,
		Image2Url: image2Url,
	}
}
