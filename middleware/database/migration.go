package database

import (
	"memoo-backend/model"
)

func migration() {
	_ = DB.AutoMigrate(&model.MemooCoin{})
}
