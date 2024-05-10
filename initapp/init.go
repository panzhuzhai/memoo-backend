package initapp

import (
	"github.com/gin-gonic/gin"
	"memoo-backend/config"
	"memoo-backend/logger"
	"memoo-backend/middleware/database"
	"memoo-backend/oss"
)

func Init() {

	// 读取翻译文件
	if err := config.LoadLocales("config/locales/zh-cn.yaml"); err != nil {
		//util.Log().Panic("翻译文件加载失败", err)
	}

	config.LoadAppConfig()
	gin.SetMode(config.AppConfig.Extension.GinMode)

	logger.InitLogger(gin.Mode())
	config.InitBtcNetwork(gin.Mode())

	database.InitDatabase()
	oss.InitAws()
}
