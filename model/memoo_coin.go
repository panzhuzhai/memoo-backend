package model

import "gorm.io/gorm"

type MemooCoin struct {
	gorm.Model
	TokenName         string `json:"tokenName" gorm:"type:varchar(100);uniqueIndex:uidx_tokenName"`
	Ticker            string `json:"ticker" gorm:"type:varchar(100);uniqueIndex:uidx_ticker"`
	ContractAddress   string `json:"contractAddress" gorm:"type:varchar(100);"`
	LPContractAddress string `json:"lpContractAddress" gorm:"type:varchar(100);"`
	ImageUrl          string `json:"imageUrl" gorm:"type:varchar(200);"`
	Description       string `json:"description" gorm:"type:varchar(2048);"`
}
