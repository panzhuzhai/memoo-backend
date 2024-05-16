package oss

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudfront/sign"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
	"log"
	"memoo-backend/localconfig"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var AWScon *s3.S3
var sess *session.Session
var imageExts []string

func InitAws() error {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(localconfig.AppConfig.AwsAttribute.AwsRegion),
		Credentials: credentials.NewStaticCredentials(localconfig.AppConfig.AwsAttribute.AwsAccessKeyId,
			localconfig.AppConfig.AwsAttribute.AwsSecretAccessKey, ""),
	})
	if err != nil {
		return err
	}

	AWScon = s3.New(sess)
	imageExts = []string{".png", ".jpg", ".jpeg", ".gif", ".bmp", ".svg"}
	return nil
}

func BatchUploadFile(files []*multipart.FileHeader) ([]string, error) {
	urls := make([]string, 0)
	for _, file := range files {
		bucketFileKey := (uuid.New()).String()
		fileOpen, err := file.Open()
		defer fileOpen.Close()
		err = DirectUploadToS3(bucketFileKey, fileOpen, file.Filename)
		if err != nil {
			return nil, err
		}
		urls = append(urls, bucketFileKey)
	}
	return urls, nil
}

func isImage(fileExt string) bool {
	for _, ext := range imageExts {
		if fileExt == ext {
			return true
		}
	}
	return false
}

func DirectUploadToS3(bucketFileKey string, file multipart.File, srcFileName string) error {
	_, err := AWScon.PutObject(BuildPutObjectInput(bucketFileKey, file, srcFileName))
	return err
}

func BuildPutObjectInput(bucketFileKey string, file multipart.File, srcFileName string) *s3.PutObjectInput {
	ext := strings.ToLower(filepath.Ext(srcFileName))
	if isImage(ext) {
		return &s3.PutObjectInput{
			Bucket:      aws.String(localconfig.AppConfig.AwsAttribute.AwsBucketName),
			Key:         aws.String(bucketFileKey),
			Body:        file,
			ContentType: aws.String("image/jpeg"),
		}
	}
	return &s3.PutObjectInput{
		Bucket: aws.String(localconfig.AppConfig.AwsAttribute.AwsBucketName),
		Key:    aws.String(bucketFileKey),
		Body:   file,
	}
}

func UploadToS3(filePath, fileName string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	key := fileName

	params := &s3.PutObjectInput{
		Bucket:      aws.String(localconfig.AppConfig.AwsAttribute.AwsBucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String("image/jpeg"),
	}

	_, err = AWScon.PutObject(params)
	if err != nil {
		return err
	}
	return nil
}

func GetS3ObjectURL(objectKey string) string {
	req, _ := AWScon.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(localconfig.AppConfig.AwsAttribute.AwsBucketName),
		Key:    aws.String(objectKey),
	})
	url, err := req.Presign(time.Duration(localconfig.AppConfig.AwsAttribute.AwsExpirationTime) * time.Hour)
	if err != nil {
		fmt.Println("Failed to generate pre-signed URL:", err)
		return ""
	}
	return url
}

func GetCdnURL(objectKey string, timeOut time.Duration) string {
	signer := sign.NewURLSigner("", nil)
	url := GetS3ObjectURL(objectKey)
	expirationTime := time.Now().Add(24 * time.Hour)
	signedURL, err := signer.Sign(url, expirationTime)
	if err != nil {
		log.Println(err)
	}

	return signedURL
}
