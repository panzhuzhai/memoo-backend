package server

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"memoo-backend/api"
	"memoo-backend/config"
	"memoo-backend/docs"
	"memoo-backend/middleware"
	"memoo-backend/middleware/jwt"
	"memoo-backend/serializer"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io:8080
// @BasePath /v2
func NewRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	initJwt := config.InitJwt()

	r.NoRoute(initJwt.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, serializer.Fail(http.StatusNotFound, "Page not found"))
	})
	versionPath := "/backend/api/v1"
	docs.SwaggerInfo.BasePath = versionPath
	//url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router
	v1 := r.Group(versionPath)
	//ping
	v1.POST("ping", api.Ping)
	// user login
	v1.POST("auth/login", initJwt.LoginHandler)

	//authGroup
	{
		// Refresh time can be longer than token timeout
		authGroup := v1.Group("")
		authGroup.Use(initJwt.MiddlewareFunc())
		authGroup.POST("auth/refresh-token", initJwt.RefreshHandler)
		authGroup.POST("auth/logout", initJwt.LogoutHandler)
	}

	unauthorized := v1.Group("web-unauthorized/")
	{
		//Creator Ranking
		unauthorized.GET("trending-creators", api.TrendingCreators)
		unauthorized.GET("top-creators", api.TopCreators)

		//token Ranking
		unauthorized.GET("trending-tokens", api.TrendingTokens)
		unauthorized.GET("top-tokens", api.TopTokens)
	}

	webOrientedGroup := v1.Group("web-oriented/")
	{
		//user
		webOrientedGroup.Use(initJwt.MiddlewareFunc())
		webOrientedGroup.GET("user-view-others", api.UserViewOthers)
		webOrientedGroup.GET("user-view-others-list", api.UserViewOthersList)
		webOrientedGroup.POST("user", api.UserNewOrEdit)

		//ido
		webOrientedGroup.GET("ido-active", api.IdoActive)
		webOrientedGroup.GET("ido-completed", api.IdoCompleted)
		webOrientedGroup.GET("ido-upcoming", api.IdoUpcoming)

		//ido-detail
		webOrientedGroup.GET("ido-upcoming-detail", api.IdoUpcomingDetail)
		webOrientedGroup.GET("ido-active-detail", api.IdoActiveDetail)
		webOrientedGroup.GET("ido-launched-detail", api.IdoLaunchedDetail)
		webOrientedGroup.GET("ido-launched-detail-top10", api.IdoLaunchedDetailTop10List)

		//project
		webOrientedGroup.POST("project", api.ProjectNewOrEdit)

		//coin
		webOrientedGroup.POST("token", api.TokenNewOrEdit)
		webOrientedGroup.GET("token", api.TokenList)

	}
	return r
}
