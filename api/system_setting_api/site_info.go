package system_setting_api

import (
	"github.com/gin-gonic/gin"
	"my_blog/global"
	"my_blog/response"
)

type GetSiteInfoResp struct {
	CreateAt string `json:"create_at"`
	Title    string `json:"title"`
	Job      string `json:"job"`
	Addr     string `json:"addr"`
	Github   string `json:"github"`
}

func GetSiteInfo(c *gin.Context) {
	response.OKWithData(c, GetSiteInfoResp{
		CreateAt: global.Config.SiteInfo.CreateAt,
		Title:    global.Config.SiteInfo.Title,
		Job:      global.Config.SiteInfo.Job,
		Addr:     global.Config.SiteInfo.Addr,
		Github:   global.Config.SiteInfo.Github,
	})
}
