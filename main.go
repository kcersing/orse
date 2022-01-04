package main

import (
	"fmt"
	"orse/app/api"
	"orse/router"

)
func main()  {
	r := router.Router()
	api.Api(r)
	if err:=r.Run(":8080");err!=nil{
		fmt.Println("startup service failed, err:%v\n", err)
	}
}

