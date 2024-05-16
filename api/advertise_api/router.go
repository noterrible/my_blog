package advertise_api

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	g := r.Group("/advertise")
	g.POST("/create", CreateAdvertise)
	g.POST("/list", ListAdvertise)
}
