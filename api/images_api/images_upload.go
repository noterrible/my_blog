package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"my_blog/global"
	"my_blog/response"
	"strings"
)

type UploadReq struct {
}
type UploadResp struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

var list = []string{
	"png",
	"jpg",
	"gif",
	"jpeg",
}

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.FailedWithMsg(c, err.Error())
		return
	}
	fileList, ok := form.File["images"]
	if !ok {
		response.FailedWithMsg(c, "没有文件")
		return
	}
	respL := make([]UploadResp, 0)
	for _, header := range fileList {

		//限制文件大小
		if header.Size > global.Config.Image.Size {
			respL = append(respL, UploadResp{
				FileName:  header.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("more than size : %d M", int(global.Config.Image.Size/1024/1024)),
			})
			continue
		} else if !isImage(header.Filename) {
			respL = append(respL, UploadResp{
				FileName:  header.Filename,
				IsSuccess: false,
				Msg:       "not supported file",
			})
			continue
		}
		//保存到本地
		err := c.SaveUploadedFile(header, global.Config.Image.Path+header.Filename)
		if err != nil {
			respL = append(respL, UploadResp{
				FileName:  header.Filename,
				IsSuccess: false,
				Msg:       "failed",
			})
			response.FailedWithMsg(c, err.Error())
			return
		}
		respL = append(respL, UploadResp{
			FileName:  header.Filename,
			IsSuccess: true,
			Msg:       "success",
		})
	}
	response.OKWithData(c, respL)
}
func isImage(fileName string) bool {
	sList := strings.Split(fileName, ".")
	suffix := sList[len(sList)-1]
	for _, s := range list {
		if s == suffix {
			return true
		}
	}
	return false
}
