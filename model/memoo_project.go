package model

import "gorm.io/gorm"

type MemooProject struct {
	gorm.Model
	ProjectName  string `json:"projectName" gorm:"type:varchar(100);uniqueIndex:uidx_projectName"`
	Introduction string `gorm:"type:text"`
}
