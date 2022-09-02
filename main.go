/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 16:51:22
 * @FilePath: /quake_go/main.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"quake/src/apis"
	"strings"
)

func main() {
	test()

}

func Init() {
	fmt.Println("Starting Quake Cli...")
	var model, token, query, start, size string
	flag.StringVar(&size, "size", "10", "size String value")
	flag.StringVar(&start, "start", "0", "start String value")
	flag.StringVar(&query, "query", "", "query String value")
	flag.StringVar(&model, "model", "", "model String value,example: server,info")
	flag.StringVar(&token, "token", "", "token Sting value")

	flag.Parse()

	if len(os.Args) == 1 {
		fmt.Println(`example:
	./quake -token=<token> -model=<model> -query=<query>`)
		fmt.Println(`
Usage of ./quake_go:
  -model string
        model String value (default "info")
  -query string
        query String value
  -size string
        size String value (default "10")
  -start string
        start String value (default "0")
  -token string
        token Sting value`)
		return
	}
	if token == "" {
		fmt.Println("!!!!token is empty!!!!")
		return
	}

	switch strings.ToLower(model) {
	case "info":
		apis.InfoGet(token)
	case "server":
		apis.SearchServicePost(query, start, size, token)
	case "host":
		fmt.Println("主机数据接口待完成。。。")
	case "favicon":
		fmt.Println("favicon相似度接口待完成。。。")
	default:
		fmt.Println("model参数错误")
	}
}

type Server struct {
	ServerName string // 首字母大写，首字母不敏感，其他字母敏感
	ServerIP   string // 首字母大写
}

type Serverslice struct {
	Servers []Server
}

func test() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Local_Web","serverIP":"127.0.0.1"},{"serverName":"Local_DB","serverIP":"127.0.0.1"}]}`
	json.Unmarshal([]byte(str), &s)

	// 输出转换后第一条
	fmt.Println(s.Servers[0])
}
