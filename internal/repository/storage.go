package repository

import "io"

type AboutSectionStorageRepo interface {
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
	DeleteImage(fileName string) error
}

type PhilosopySectionStorageRepo interface {
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
	DeleteImage(fileName string) error
}

type StorySectionStorageRepo interface {
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
	DeleteImage(fileName string) error
}

type EventGallaryStorageRepo interface {
	DeleteImage(fileName string) error
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
}
type BlogStorageRepo interface {
	DeleteImage(fileName string) error
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
}

type TestimonialStorageRepo interface {
	DeleteImage(fileName string) error
	UploadImage(fileName string, file io.ReadSeeker) (string, error)
}

type CourseStorageRepo interface {
	UploadFile(fileName string, file io.ReadSeeker) (string, error)
	DeleteFile(fileName string) error
}
