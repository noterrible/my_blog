package images_api

import "github.com/gin-gonic/gin"

func InitRouter(r *gin.Engine) {
	g := r.Group("/images")
	g.GET("/upload", Upload)
}
