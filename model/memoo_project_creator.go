package model

import "gorm.io/gorm"

type MemooProjectCreator struct {
	gorm.Model
	ProjectId      uint   `json:"projectId" gorm:"type:bigint"`
	CreatorAddress string `json:"creatorAddress" gorm:"type:varchar(100);"`
}
