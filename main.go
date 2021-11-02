package main

import (
	"fmt"
	"ginjwt/pkg/setting"
	"ginjwt/routers"

	"github.com/fvbock/endless"
)

func main() {

	// r := gin.Default()
	// r.Run()

	var err error
	endless.DefaultReadTimeOut = setting.ReadTimeOut
	endless.DefaultWriteTimeOut = setting.WriteTimeOut
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HttpPort)

	fmt.Printf("运行端口:%v\n", endPoint)

	server := endless.NewServer(endPoint, routers.InitRouter())

	err = server.ListenAndServe()
	if err != nil {
		//panic(err)
		fmt.Printf("%v\n", err.Error())
		return
	}

}
