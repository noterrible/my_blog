package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"
	"my_blog/global"
	"my_blog/response"
	"strings"
)

type UploadReq struct {
}

var list = []string{
	"png",
	"jpg",
	"gif",
	"jpeg",
}

type UploadResp struct {
	FileName  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"`
	Https     string `json:"https"`
	Msg       string `json:"msg"`
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
		header.Open()
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
		ioReader, err := header.Open()
		if err != nil {
			respL = append(respL, UploadResp{
				FileName:  header.Filename,
				IsSuccess: false,
				Msg:       "server error",
			})
			continue
		}
		//
		ID := ksuid.New()
		var saveSuffix string
		sp := strings.Split(header.Filename, ".")
		if len(sp) > 0 {
			saveSuffix = sp[len(sp)-1]
		}
		saveName := ID.String() + "." + saveSuffix
		//保存到oss
		err = global.Bucket.PutObject(saveName, ioReader)
		if err != nil {
			respL = append(respL, UploadResp{
				FileName:  header.Filename,
				IsSuccess: false,
				Msg:       "server error :bucket",
			})
			continue
		}
		respL = append(respL, UploadResp{
			FileName:  header.Filename,
			IsSuccess: true,
			Https:     global.Config.Image.GetDomain + saveName,
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
