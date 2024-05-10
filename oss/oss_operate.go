package oss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"memoo-backend/config"
	"mime/multipart"
	"os"
)

var AWScon *s3.S3

func InitAws() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(config.AppConfig.AwsAttribute.AwsRegion),
		Credentials: credentials.NewStaticCredentials(config.AppConfig.AwsAttribute.AwsAccessKeyId,
			config.AppConfig.AwsAttribute.AwsSecretAccessKey, ""),
	})
	if err != nil {
		return err
	}

	AWScon = s3.New(sess)
	return nil
}

func BatchUploadFile(files []*multipart.FileHeader) ([]string, error) {
	urls := make([]string, 0)
	for _, bannerFile := range files {
		fileUrl, err := UploadFile(bannerFile)
		if err != nil {
			return nil, err
		}
		urls = append(urls, fileUrl)
	}
	return urls, nil
}

func UploadFile(file *multipart.FileHeader) (string, error) {
	filePath := "./uploads/" + file.Filename
	err := SaveFile(file, filePath)
	if err != nil {
		return "", err
	}
	url, err := UploadToS3(AWScon, "your-s3-bucket", filePath, file.Filename)
	if err != nil {
		return "", err
	}

	return url, nil
}

func SaveFile(file *multipart.FileHeader, filePath string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	return nil
}

func UploadToS3(s3Client *s3.S3, bucketName, filePath, fileName string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := fileName

	params := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   file,
	}

	_, err = s3Client.PutObject(params)
	if err != nil {
		return "", err
	}
	url := getS3ObjectURL(AWScon, bucketName, key)
	return url, nil
}

func getS3ObjectURL(s3Client *s3.S3, bucketName, objectKey string) string {
	req, _ := s3Client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	url, err := req.Presign(0)
	if err != nil {
		fmt.Println("Failed to generate pre-signed URL:", err)
		return ""
	}

	return url
}
