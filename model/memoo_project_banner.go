package model

import "gorm.io/gorm"

type MemooProjectBanner struct {
	gorm.Model
	ProjectId uint   `json:"projectId" gorm:"type:bigint"`
	BannerUrl string `json:"bannerurl" gorm:"type:varchar(100);"`
}
