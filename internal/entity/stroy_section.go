package entity

type StorySection struct {
	Image1Url string `json:"image1_url" bson:"image1_url"`
	Image2Url string `json:"image2_url" bson:"image2_url"`
	Image3url string `json:"image3_url" bson:"image3_url"`
	Image4url string `json:"image4_url" bson:"image4_url"`
}

func NewStorySection(image1url, image2url, image3url, image4url string) *StorySection {
	return &StorySection{
		Image1Url: image1url,
		Image2Url: image2url,
		Image3url: image3url,
		Image4url: image4url,
	}
}
