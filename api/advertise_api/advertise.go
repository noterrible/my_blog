package advertise_api

import (
	"github.com/gin-gonic/gin"
	"my_blog/dao"
	"my_blog/global"
	"my_blog/models"
	"my_blog/response"
	"my_blog/services"
)

type CreateAdvertiseReq struct {
	Title  string `json:"title" binding:"required"`
	Href   string `json:"href" binding:"required,url"`
	IsShow bool   `json:"is_show" binding:"required"`
	Link   string `json:"link" binding:"required,url"`
}

func CreateAdvertise(c *gin.Context) {
	var createAdvertiseReq CreateAdvertiseReq
	err := c.ShouldBindJSON(&createAdvertiseReq)
	if err != nil {
		global.Log.Error("ShouldBindJSON" + err.Error())
		response.FailedWithMsg(c, "参数错误")
		return
	}
	err = dao.CreateAdvertise(models.Advertise{
		Title:  createAdvertiseReq.Title,
		Href:   createAdvertiseReq.Href,
		IsShow: createAdvertiseReq.IsShow,
		Link:   createAdvertiseReq.Link,
	})
	if err != nil {
		global.Log.Error("CreateAdvertise" + err.Error())
		response.FailedWithMsg(c, "创建失败")
		return
	}
	response.OK(c)
}

type ListAdvertiseReq struct {
	Title string `json:"title"`
	services.PageInfo
}

func ListAdvertise(c *gin.Context) {
	var listAdvertiseReq ListAdvertiseReq
	err := c.ShouldBind(&listAdvertiseReq)
	if err != nil {
		global.Log.Error("ShouldBindJSON" + err.Error())
		response.FailedWithMsg(c, "参数错误")
		return
	}
	req := services.Option{
		PageInfo:         listAdvertiseReq.PageInfo,
		LikeColumns:      []string{"title"},
		LikeColumnsValue: []string{listAdvertiseReq.Title},
	}
	var res []*models.Advertise
	resp, err := services.GetList(res, req)
	if err != nil {
		global.Log.Error("GetList" + err.Error())
		response.FailedWithMsg(c, "服务器内部错误")
		return
	}
	response.OKWithData(c, resp)
}
