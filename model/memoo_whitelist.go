package model

import "gorm.io/gorm"

type MemooWhitelist struct {
	gorm.Model
	ProjectName string `json:"projectName" gorm:"type:varchar(100);uniqueIndex:uidx_ProjectName_BtcAddress"`
	EthAddress  string `json:"ethAddress" gorm:"type:varchar(100);"`
	BtcAddress  string `json:"btcAddress" gorm:"type:varchar(100);uniqueIndex:uidx_ProjectName_BtcAddress"`
	PublicKey   string `json:"publicKey" gorm:"type:varchar(100);"`
	Count       int64  `json:"count" gorm:"type:bigint;"`
	Status      bool   `json:"status" gorm:"type:boolean;default:false;"`
}
