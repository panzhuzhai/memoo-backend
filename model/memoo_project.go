package model

import "gorm.io/gorm"

type MemooProject struct {
	gorm.Model
	Ticker       string `json:"ticker" gorm:"type:varchar(100);"`
	ProjectName  string `json:"projectName" gorm:"type:varchar(100);uniqueIndex:uidx_projectName"`
	Introduction string `gorm:"type:text"`
	Description  string `json:"description" gorm:"type:varchar(2048);"`
	Twitter      string `gorm:"type:varchar(200);"`
}
