package global

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"my_blog/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
	Bucket *oss.Bucket
)
