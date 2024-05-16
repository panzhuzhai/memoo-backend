package model

import "gorm.io/gorm"

type MemooProjectSocial struct {
	gorm.Model
	ProjectId       uint   `json:"projectId" gorm:"type:bigint"`
	SocialUrl       string `json:"socialUrl" gorm:"type:varchar(100);"`
	SocialType      string `json:"socialType" gorm:"type:varchar(20);"`
	MajorCategories string `json:"majorCategories" gorm:"type:varchar(20);"`
}
