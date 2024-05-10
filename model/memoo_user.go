package model

import "gorm.io/gorm"

type MemooUser struct {
	gorm.Model
	UserName        string `json:"userName" gorm:"type:varchar(100)"`
	UserBio         string `json:"userBio" gorm:"type:text;"`
	Address         string `json:"address" gorm:"type:varchar(100);"`
	Twitter         string `json:"twitter" gorm:"type:varchar(200)"`
	WebsiteUrl      string `json:"websiteUrl" gorm:"type:varchar(200);"`
	ProfileImageUrl string `json:"profileImageUrl" gorm:"type:varchar(200);"`
}
