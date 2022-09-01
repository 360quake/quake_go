/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: ph4nt0mer rootphantomy@hotmail.com
 * @LastEditTime: 2022-09-01 22:50:21
 * @FilePath: /quake_go/quake_go.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package main

import (
	"flag"
	"fmt"
	"os"
	"quake/src/apis"
	"strings"
)

func main() {
	Init()
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
        model String value
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
		fmt.Println("!!!!token is empty")
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

func Init() {
	fmt.Println("Starting Quake Cli...")
}
