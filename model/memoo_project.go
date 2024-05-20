package model

import "gorm.io/gorm"

type MemooProject struct {
	gorm.Model
	Ticker         string `json:"ticker" gorm:"type:varchar(100);uniqueIndex:uidx_project_ticker"`
	ProjectName    string `json:"projectName" gorm:"type:varchar(100);"`
	Website        string `json:"website" gorm:"type:varchar(100);"`
	Introduction   string `gorm:"type:text"`
	Description    string `json:"description" gorm:"type:varchar(2048);"`
	Twitter        string `gorm:"type:varchar(100);"`
	CreatorAddress string `json:"creatorAddress" gorm:"type:varchar(100);"`
	Telegram       string `json:"telegram" gorm:"type:varchar(100);"`
	PinnedTwitter  string `json:"pinnedTwitter" gorm:"type:varchar(100);"`
}
