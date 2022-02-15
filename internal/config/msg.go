package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/ini.v1"
	"os"
)

func GetError(err validator.ValidationErrors) string {
	cfg, iniErr := ini.Load("config/msg/error_message.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	for _, val := range err {
		return cfg.Section(val.Tag()).Key(val.Field()).String()
	}
	return "参数错误"
}