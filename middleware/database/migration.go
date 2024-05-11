package database

import (
	"memoo-backend/model"
)

func migration() {
	_ = DB.AutoMigrate(&model.MemooProject{})
	_ = DB.AutoMigrate(&model.MemooProjectBanner{})
	_ = DB.AutoMigrate(&model.MemooProjectCreator{})
	_ = DB.AutoMigrate(&model.MemooProjectSocial{})
	_ = DB.AutoMigrate(&model.MemooToken{})
	_ = DB.AutoMigrate(&model.MemooUser{})
	_ = DB.AutoMigrate(&model.MemooUserBanner{})
	_ = DB.AutoMigrate(&model.MemooWhitelist{})
}
