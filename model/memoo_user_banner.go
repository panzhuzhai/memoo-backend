package model

import "gorm.io/gorm"

type MemooUserBanner struct {
	gorm.Model
	UserAddress      string `json:"userAddress" gorm:"type:varchar(100)"`
	ProfileBannerUrl string `json:"profileBannerUrl" gorm:"type:varchar(200);"`
}
