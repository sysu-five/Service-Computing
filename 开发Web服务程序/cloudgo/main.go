package main

import (
	"os"
	flag "github.com/spf13/pflag"
	"github.com/sysu-five/cloudgo/service"
)

// 设置默认端口
const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	// 若没有监听端口则设为默认端口
	if len(port) == 0 {
		port = PORT
	}
	// 用户可通过-p参数设置端口
	pPort := flag.StringP("port","p",PORT,"PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	// 启动server
	service.NewServer(port)
}