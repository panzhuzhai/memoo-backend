package model

import "gorm.io/gorm"

type MemooToken struct {
	gorm.Model
	TokenName            string `json:"tokenName" gorm:"type:varchar(100);uniqueIndex:uidx_tokenName"`
	Ticker               string `json:"ticker" gorm:"type:varchar(100);uniqueIndex:uidx_ticker"`
	ContractAddress      string `json:"contractAddress" gorm:"type:varchar(100);"`
	LPContractAddress    string `json:"lpContractAddress" gorm:"type:varchar(100);"`
	ImageUrl             string `json:"imageUrl" gorm:"type:varchar(100);"`
	Description          string `json:"description" gorm:"type:varchar(2048);"`
	Status               string `json:"status" gorm:"type:varchar(20);"`
	PreLaunchDuration    string `json:"preLaunchDuration" gorm:"type:varchar(30);"`
	PreMarketAcquisition string `json:"preMarketAcquisition" gorm:"type:varchar(30);"`
}
