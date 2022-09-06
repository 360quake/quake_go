/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 16:39:06
 * @FilePath: /quake_go/main.go
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
	"quake/src/utils"
	"strings"
)

func main() {
	Init()
}

func Init() {
	fmt.Println("Starting Quake Cli...")
	num := len(os.Args)
	if num < 2 {
		fmt.Println(`
Usage of ./quake:
  init <token>
		token initialization
  info 
        person infomation
  query <string>
        query string value
  -t <string,string>
        field string value(example:./quake query port:8088 -t ip,port)
  -size string
        size String value (default "10")
  -start string
        start String value (default "0")`)
		return
	}
	path := "./config.yaml"
	if strings.ToLower(os.Args[1]) == "init" {
		if num < 3 {
			fmt.Println("!!!!token is empty !!!!")
			return
		}
		utils.WriteYaml(path, os.Args[2])
		return
	}
	token, status := utils.ReadYaml(path)
	if status {
		fmt.Println("!!!!please ./quake init token!!!!")
		return
	}
	start, size, field := flaginit()
	switch strings.ToLower(os.Args[1]) {
	case "info":
		apis.InfoGet(token.Token)
	case "query":
		if num < 3 {
			fmt.Println("!!!!query is empty !!!!")
			return
		}
		apis.SearchServicePost(os.Args[2], start, size, token.Token, field)
	case "host":
		fmt.Println("主机数据接口待完成。。。")
	case "favicon":
		fmt.Println("favicon相似度接口待完成。。。")
	case "domain":
		fmt.Println("domain ")
	default:
		fmt.Println("args failed !!!!")
	}
}

func flaginit() (start string, size string, field string) {
	flag.StringVar(&field, "t", "", "field String value")
	flag.StringVar(&size, "size", "10", "size String value")
	flag.StringVar(&start, "start", "0", "start String value")
	flag.Parse()
	return
}
