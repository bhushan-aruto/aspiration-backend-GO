package storage

import (
	"fmt"
	"io"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type StorySectionStorageRepo struct {
	bucketName        string
	dirName           string
	cloudFrontBaseUrl string
	s3                *S3Connection
}

func NewStorySectionRepo(bucketName, dirName, cloudFrontBaseUrl string, s3 *S3Connection) *StorySectionStorageRepo {
	return &StorySectionStorageRepo{
		bucketName:        bucketName,
		dirName:           dirName,
		cloudFrontBaseUrl: cloudFrontBaseUrl,
		s3:                s3,
	}
}

func (repo *StorySectionStorageRepo) UploadImage(fileName string, file io.ReadSeeker) (string, error) {
	filePath := fmt.Sprintf("/%v/%v", repo.dirName, fileName)

	log.Println(repo.bucketName, filePath)

	_, err := repo.s3.S3Client.PutObject(
		&s3.PutObjectInput{
			Bucket: aws.String(repo.bucketName),
			Key:    aws.String(filePath),
			Body:   file,
		},
	)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v/%v/%v", repo.cloudFrontBaseUrl, repo.dirName, fileName), nil

}

func (repo *StorySectionStorageRepo) DeleteImage(fileName string) error {
	filePath := fmt.Sprintf("%v/%v", repo.dirName, fileName)

	_, err := repo.s3.S3Client.DeleteObject(
		&s3.DeleteObjectInput{
			Bucket: &repo.bucketName,
			Key:    aws.String(filePath),
		},
	)

	return err
}
