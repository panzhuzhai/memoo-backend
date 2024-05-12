package test

import (
	"github.com/google/uuid"
	"log"
	"memoo-backend/initapp"
	"memoo-backend/oss"
	"testing"
	"time"
)

func TestUploadToS3(t *testing.T) {
	initapp.Init()
	fileName := (uuid.New()).String()
	err := oss.UploadToS3("C:\\Users\\leolj\\Downloads\\Telegram Desktop\\image_2024-05-07_15-03-09.png", fileName)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("fileName:", fileName)
}

func TestGetS3ObjectURL(t *testing.T) {
	initapp.Init()
	url := oss.GetS3ObjectURL("b548ad8d-44c8-4990-973a-9a0082f34398")
	log.Println("url:", url)
}

func TestGetCdnURL(t *testing.T) {
	initapp.Init()
	//url := oss.GetS3ObjectURL("Capture001.png", 24*time.Hour)
	url := oss.GetCdnURL("image_2024-05-08_14-52-28.png", 24*time.Hour)
	log.Println("url:", url)
}
