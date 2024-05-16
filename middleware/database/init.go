package database

import (
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"memoo-backend/constants"
	"memoo-backend/localconfig"
	"memoo-backend/logger"
	"os"
	"time"
)

var DB *gorm.DB

func InitDatabase() {
	logger.ZapLogger.Infof("Try database connection")
	gormConfig := &gorm.Config{Logger: logger.GormLogger}
	//fmt.Println(gormConfig.)
	db, err := gorm.Open(postgres.Open(localconfig.AppConfig.Database.Dsn), gormConfig)
	if err != nil {
		logger.ZapLogger.Errorf("Failure database connection")
		os.Exit(constants.ErrExitStatus)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.ZapLogger.Error("db lost: %v", zap.Error(err))
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Hour)
	logger.ZapLogger.Infof("Success database connection")
	DB = db
	migration()
}
