package main

import (
	"fmt"
	"orse/app/api"
	"orse/router"
)

func main() {
	api.Router()
	if err := router.R.Run(":8080"); err != nil {
		fmt.Println("startup service failed, err:%v\n", err)
	}
}
