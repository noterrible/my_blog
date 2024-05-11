package system_setting_api

import (
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	g := r.Group("/system_setting")
	g.GET("/site_info", GetSiteInfo)
}
