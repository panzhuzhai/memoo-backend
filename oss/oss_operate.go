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
	"time"
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
	url, err := UploadToS3(filePath, file.Filename)
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

func UploadToS3(filePath, fileName string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	key := fileName

	params := &s3.PutObjectInput{
		Bucket:      aws.String(config.AppConfig.AwsAttribute.AwsBucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	}

	_, err = AWScon.PutObject(params)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func GetS3ObjectURL(objectKey string, timeOut time.Duration) string {
	req, _ := AWScon.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(config.AppConfig.AwsAttribute.AwsBucketName),
		Key:    aws.String(objectKey),
	})

	url, err := req.Presign(timeOut)
	if err != nil {
		fmt.Println("Failed to generate pre-signed URL:", err)
		return ""
	}

	return url
}
