package test

import (
	"log"
	"memoo-backend/initapp"
	"memoo-backend/oss"
	"testing"
	"time"
)

func TestUploadToS3(t *testing.T) {
	initapp.Init()
	url, err := oss.UploadToS3("C:\\Users\\leolj\\Pictures\\Feedback\\{0E72E60B-31DF-4887-91DE-3993B085AFC3}\\Capture001.png", "Capture001.png")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("url:", url)
}

func TestGetS3ObjectURL(t *testing.T) {
	initapp.Init()
	url := oss.GetS3ObjectURL("Capture001.png", 7*24*time.Hour)
	log.Println("url:", url)
}
