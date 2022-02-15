package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)
//https://mp.weixin.qq.com/s?__biz=Mzg2ODU1MTI0OA==&mid=2247484889&idx=1&sn=a72ed7c901ad88d38b8dd8d2e026d04b&chksm=ceabdae6f9dc53f09a71ee5975ea70f8f86f30e58571640dff5048eb469381d895591a5f80a1&scene=21#wechat_redirect
type Config struct {
	Name string `yaml:"Name"`
	Host string `yaml:"Host"`
	Port string `yaml:"Port"`
	Mysql Mysql `yaml:"Mysql"`
	Auth Auth `yaml:"Auth"`
	Salt string `yaml:"Salt"`
}
type Mysql struct{
	DataSource string `yaml:"DataSource"`
}
type Auth struct{
	AccessSecret string `yaml:"AccessSecret"`
	AccessExpire int `yaml:"AccessExpire"`
}


func GetConfig(filePath string)  Config{
	c := Config{}
	buffer,err:=ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("解析config.yaml读取错误: %v", err)
	}
	if yaml.Unmarshal(buffer, &c) != nil {
		log.Fatalf("解析config.yaml出错: %v", err)
	}
	return c
}

func NewConfig() Config {
	return GetConfig("config/user/user.yaml")
}