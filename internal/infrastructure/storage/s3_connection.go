package storage

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Connection struct {
	S3Client *s3.S3
}

func NewS3Connection(awsRegion, awsAccessKeyId, awsSecretAccessKey string) *S3Connection {

	session, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awsRegion),
			Credentials: credentials.NewStaticCredentials(
				awsAccessKeyId,
				awsSecretAccessKey,
				"",
			),
		},
	)

	if err != nil {
		log.Fatalln("error occured with s3 ,error :", err.Error())
	}
	return &S3Connection{
		S3Client: s3.New(session),
	}

}
