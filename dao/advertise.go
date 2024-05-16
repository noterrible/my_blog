package dao

import (
	"my_blog/global"
	"my_blog/models"
)

func CreateAdvertise(advertise models.Advertise) error {
	err := global.DB.Create(&advertise).Error
	if err != nil {
		global.Log.Error("创建广告失败", err)
		return err
	}
	return nil
}
