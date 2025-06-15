package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress              string
	DatabaseUrl                string
	DatabaseName               string
	RabbitMQUrl                string
	S3Region                   string
	AwsAccessKeyId             string
	AwsSecretAccessKey         string
	S3BucketName               string
	CourseS3BucketName         string
	AboutSectionDirName        string
	PhilosopysectionDirName    string
	StorysectionDirName        string
	CloudFrontUrl              string
	CourseCloufontUrl          string
	EventGallerysectionDirName string
	BlogsSectionDirName        string
	TestimonialSectionDirName  string
	CourseSectionDirName       string
	FromEmail                  string
	EmailAppPassword           string
	ToEmail                    string
	SmtpHost                   string
	SmtpPort                   string
	WhatspNumber               string
	PaymentGatewayUrl          string
	PaymentGateWayKeyId        string
	PaymentGateWayKeySecrete   string
}

func LoadConfig() *Config {

	serverMode := os.Getenv("SERVER_MODE")

	if serverMode == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Fatalln("missing .env file. Error: ", err.Error())
		}
	} else if serverMode != "prod" {
		log.Println("Invalid SERVER_MODE. Please set SERVER_MODE to 'dev' or 'prod'.")

	}

	return &Config{
		ServerAddress:              getEnv("SERVER_ADDRESS"),
		DatabaseUrl:                getEnv("DATABASE_URL"),
		DatabaseName:               getEnv("DATABASE_NAME"),
		RabbitMQUrl:                getEnv("RABBITMQ_URL"),
		S3Region:                   getEnv("AWS_S3_REGION"),
		AwsAccessKeyId:             getEnv("AWS_ACCESS_KEY_ID"),
		AwsSecretAccessKey:         getEnv("AWS_SECRETE_ACCESS_KEY"),
		S3BucketName:               getEnv("AWS_S3_BUCKET_NAME"),
		CourseS3BucketName:         getEnv("AWS_S3_COURSE_BUCKET_NAME"),
		AboutSectionDirName:        getEnv("AWS_S3_ABOUT_SECTION_DIR_NAME"),
		CloudFrontUrl:              getEnv("CLOUD_FRONT_URL"),
		CourseCloufontUrl:          getEnv("COURSE_CLOUD_FRONT_URL"),
		PhilosopysectionDirName:    getEnv("AWS_S3_PHILOSOPY_SECTION_DIR_NAME"),
		StorysectionDirName:        getEnv("AWS_S3_STORY_SECTION_DIR_NAME"),
		EventGallerysectionDirName: getEnv("AWS_S3_EVENT_GALLERY_DIR_NAME"),
		BlogsSectionDirName:        getEnv("AWS_S3_BLOG_SECTION_DIR_NAME"),
		TestimonialSectionDirName:  getEnv("AWS_S3_TESTIMONIAL_SECTION_DIR_NAME"),
		CourseSectionDirName:       getEnv("AWS_S3_COURSE_SECTION_DIR_NAME"),
		FromEmail:                  getEnv("FROM_EMAIL"),
		ToEmail:                    getEnv("TO_EMAIL"),
		EmailAppPassword:           getEnv("EMAIL_APP_PASSWORD"),
		SmtpHost:                   getEnv("SMTP_HOST"),
		SmtpPort:                   getEnv("SMTP_PORT"),
		WhatspNumber:               getEnv("WHATASAPP_NUMBER"),
		PaymentGatewayUrl:          getEnv("PAYMENT_GWATEWAY_URL"),
		PaymentGateWayKeyId:        getEnv("PAYMENT_GATEWAY_KEYID"),
		PaymentGateWayKeySecrete:   getEnv("PAYMENT_GATEWAY_KEYSECRETE"),
	}

}

func getEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("missing or empty %s env variable", key)
	}
	return value
}
