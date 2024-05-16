package routers

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"my_blog/api/advertise_api"
	"my_blog/api/images_api"
	"my_blog/api/system_setting_api"
	docs "my_blog/docs"
	"my_blog/global"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	r := gin.Default()
	r.GET("ping",
		ping)
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	system_setting_api.InitRouter(r)
	images_api.InitRouter(r)
	advertise_api.InitRouter(r)
	return r
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} HelloWorld
// @Router /ping [get]
func ping(context *gin.Context) {
	context.JSON(http.StatusOK, "hello word")
}
