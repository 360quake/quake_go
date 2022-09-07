/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-07 18:00:02
 * @FilePath: /quake_go/main.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package main

import (
	"fmt"
	"os"
	"quake/src/apis"
	. "quake/src/model"
	"quake/src/utils"
	"strconv"
	"strings"
	"time"

	"github.com/hpifu/go-kit/hflag"
)

func main() {
	hflag_init()
}

func hflag_init() {
	fmt.Println("Starting Quake Cli...")
	hflag.AddFlag("start", "-st to start number", hflag.Shorthand("st"), hflag.Type("string"), hflag.DefaultValue("0"))
	hflag.AddFlag("size", "-sz to size number ", hflag.Shorthand("sz"), hflag.Type("string"), hflag.DefaultValue("10"))
	hflag.AddFlag("ignore_cache", "-ic true or false,default false", hflag.Shorthand("ic"), hflag.Type("bool"), hflag.DefaultValue("false"))
	hflag.AddFlag("start_time", "-s time flag , default time is time.now.year", hflag.Shorthand("s"), hflag.Type("time"), hflag.DefaultValue(strconv.Itoa(time.Now().Year())+"-01-01"))
	hflag.AddFlag("end_time", "-e time to end time flag", hflag.Shorthand("e"), hflag.Type("time"), hflag.DefaultValue(time.Now().Format("2006-01-02 15:04:05")))
	hflag.AddFlag("field", "-fe body to show body infomation", hflag.Shorthand("fe"), hflag.Type("string"), hflag.DefaultValue(""))
	hflag.AddFlag("query_txt", "-qt ./file.txt file to query search", hflag.Shorthand("qt"), hflag.Type("string"), hflag.DefaultValue(""))
	hflag.AddPosFlag("option", "init,info,search")
	if err := hflag.Parse(); err != nil {
		panic(err)
	}
	num := len(os.Args)
	if num < 2 {
		fmt.Println("./quake -h get help!")
		return
	}

	switch strings.ToLower(hflag.GetString("option")) {
	case "version":
		fmt.Println("version:1.4")
	case "init":
		if num < 3 {
			fmt.Println("!!!!token is empty !!!!")
			return
		}
		utils.WriteYaml("./config.yaml", os.Args[2])
	case "info":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		apis.InfoGet(token.Token)
	case "search":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		var reqjson Reqjson
		reqjson.Query = os.Args[2]
		reqjson.Start = hflag.GetString("start")
		reqjson.Size = hflag.GetString("size")
		reqjson.Start_time = hflag.GetTime("start_time")
		reqjson.End_time = hflag.GetTime("end_time")
		reqjson.Ignore_cache = hflag.GetBool("ignore_cache")
		reqjson.Field = hflag.GetString("field")
		reqjson.Query_txt = hflag.GetString("query_txt")
		apis.SearchServicePost(reqjson, token.Token)
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
