package storage

import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type CourseStorageRepo struct {
	bucketName        string
	dirName           string
	cloudFrontBaseURL string
	s3                *S3Connection
}

func NewCourseStorageRepo(bucket, dir, base string, s3 *S3Connection) *CourseStorageRepo {
	return &CourseStorageRepo{
		bucketName:        bucket,
		dirName:           dir,
		cloudFrontBaseURL: base,
		s3:                s3,
	}
}

func (repo *CourseStorageRepo) UploadFile(fileName string, file io.ReadSeeker) (string, error) {
	filePath := fmt.Sprintf("/%v/%v", repo.dirName, fileName)

	_, err := repo.s3.S3Client.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(repo.bucketName),
			Key:    aws.String(filePath),
			Body:   file,
		},
	)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%v/%v/%v", repo.cloudFrontBaseURL, repo.dirName, fileName), nil

}

// func (repo *CourseStorageRepo) UploadFile(fileName string, file io.ReadSeeker) (string, error) {
// 	filePath := fmt.Sprintf("%v/%v", repo.dirName, fileName)

// 	uploader := s3manager.NewUploaderWithClient(repo.s3.S3Client, func(u *s3manager.Uploader) {
// 		u.PartSize = 5 * 1024 * 1024
// 		u.Concurrency = 3
// 	})

// 	_, err := uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(repo.bucketName),
// 		Key:    aws.String(filePath),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		return "", err
// 	}

// 	return fmt.Sprintf("%v/%v/%v", repo.cloudFrontBaseURL, repo.dirName, fileName), nil
// }

func (repo *CourseStorageRepo) DeleteFile(fileName string) error {
	filePath := fmt.Sprintf("%v/%v", repo.dirName, fileName)

	_, err := repo.s3.S3Client.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: &repo.bucketName,
			Key:    aws.String(filePath),
		},
	)

	return err

}
