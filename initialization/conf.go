package initialization

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"my_blog/config"
	"my_blog/global"
)

func InitConfig(path string) {
	c := config.Config{}
	yamlConf, err := ioutil.ReadFile(path)
	if err != nil {
		panic("读取配置文件失败")
	}
	err = yaml.Unmarshal(yamlConf, &c)
	if err != nil {
		panic("yaml解析配置文件失败")
	}
	fmt.Println("config:", c)
	global.Config = &c
}
