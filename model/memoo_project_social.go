package model

import "gorm.io/gorm"

type MemooProjectSocial struct {
	gorm.Model
	ProjectId  uint `json:"projectId" gorm:"type:bigint"`
	SocialUrl  uint `json:"socialUrl" gorm:"type:varchar(200);"`
	SocialType uint `json:"socialType" gorm:"type:varchar(20);"`
}
