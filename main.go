/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-07 12:41:06
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

//	func Init() {
//		fmt.Println("Starting Quake Cli...")
//		num := len(os.Args)
//		if num < 2 {
//			fmt.Println(`
//
// Usage of ./quake:
//
//	  init <token>
//			token initialization
//	  info
//	        person infomation
//	  query <string>
//	        query string value
//	  -t <string>
//	        field string value(example:./quake query port:8088 -t=body,)
//	  -ic bool
//	        ignore_cache value (true or false)
//	  -s time
//			start time value, USE UTC
//	  -e time
//			end time value, USE UTC
//	  -start string
//	        start String value (default :./quake query port:8088 -start=0)
//	  -size string
//	        size String value (default :./quake query port:8088 -size=10)`)
//			return
//		}
//		path := "./config.yaml"
//		if strings.ToLower(os.Args[1]) == "init" {
//			if num < 3 {
//				fmt.Println("!!!!token is empty !!!!")
//				return
//			}
//			utils.WriteYaml(path, os.Args[2])
//			return
//		}
//		token, status := utils.ReadYaml(path)
//		if status {
//			fmt.Println("!!!!please ./quake init token!!!!")
//			return
//		}
//		var reqjson Reqjson
//		reqjson.Field = "ip,port"
//		for _, value := range os.Args {
//			tmp := strings.Split(value, "=")
//			if strings.Contains(tmp[0], "-start") {
//				// reqjson.Start = tmp[1]
//			}
//			if strings.Contains(tmp[0], "-size") {
//				// reqjson.Size = tmp[1]
//			}
//			if strings.Contains(tmp[0], "-t") {
//				// reqjson.Field = tmp[1]
//			}
//			if strings.Contains(tmp[0], "-s") {
//				// reqjson.Start_time = tmp[1]
//			}
//			if strings.Contains(tmp[0], "-e") {
//				// reqjson.End_time = tmp[1]
//			}
//			if strings.Contains(tmp[0], "-ic") {
//				if strings.Contains(tmp[1], "true") {
//					reqjson.Ignore_cache = true
//				}
//			}
//		}
//		switch strings.ToLower(os.Args[1]) {
//		case "info":
//			apis.InfoGet(token.Token)
//		case "query":
//			if num < 3 {
//				fmt.Println("!!!!query is empty !!!!")
//				return
//			}
//			reqjson.Query = os.Args[2]
//			apis.SearchServicePost(reqjson, token.Token)
//		case "host":
//			fmt.Println("主机数据接口待完成。。。")
//		case "favicon":
//			fmt.Println("favicon相似度接口待完成。。。")
//		case "domain":
//			fmt.Println("domain ")
//		default:
//			fmt.Println("args failed !!!!")
//		}
//	}
func hflag_init() {
	fmt.Println("Starting Quake Cli...")
	hflag.AddFlag("start", "start flag", hflag.Shorthand("st"), hflag.Type("string"), hflag.DefaultValue("0"))
	hflag.AddFlag("size", "size flag", hflag.Shorthand("sz"), hflag.Type("string"), hflag.DefaultValue("10"))
	hflag.AddFlag("ignore_cache", "ignore_cache value (true or false)", hflag.Shorthand("ic"), hflag.Type("bool"), hflag.DefaultValue("false"))
	hflag.AddFlag("start_time", "start time flag", hflag.Shorthand("s"), hflag.Type("time"), hflag.DefaultValue(strconv.Itoa(time.Now().Year())+"-01-01"))
	hflag.AddFlag("end_time", "end time flag", hflag.Shorthand("e"), hflag.Type("time"), hflag.DefaultValue(time.Now().Format("2006-01-02 15:04:05")))
	hflag.AddFlag("field", "field flag", hflag.Shorthand("fe"), hflag.Type("string"))
	hflag.AddPosFlag("pos", "pos flag")
	if err := hflag.Parse(); err != nil {
		panic(err)
	}
	num := len(os.Args)
	if num < 2 {
		fmt.Println(`
usage: quake [pos] [-e,end_time time=2022-09-07 12:39:56] [-fe,field string] [-h,help bool] [-ic,ignore_cache bool=false] [-sz,size string=10] [-st,start string=0] [-s,start_time time=2022-01-01]

positional options:
       pos             [string]                    pos flag

options:
   -e, --end_time      [time=2022-09-07 12:39:56]  end time flag
  -fe, --field         [string]                    field flag
   -h, --help          [bool]                      show usage
  -ic, --ignore_cache  [bool=false]                ignore_cache value (true or false)
  -sz, --size          [string=10]                 size flag
  -st, --start         [string=0]                  start flag
   -s, --start_time    [time=2022-01-01]           start time flag`)
		return
	}
	path := "./config.yaml"
	token, status := utils.ReadYaml(path)
	if status {
		fmt.Println("!!!!please ./quake init token!!!!")
		return
	}
	switch strings.ToLower(hflag.GetString("pos")) {
	case "init":
		if num < 3 {
			fmt.Println("!!!!token is empty !!!!")
			return
		}
		utils.WriteYaml(path, os.Args[2])
	case "info":
		apis.InfoGet(token.Token)
	case "query":
		var reqjson Reqjson
		reqjson.Query = os.Args[2]
		reqjson.Start = hflag.GetString("start")
		reqjson.Size = hflag.GetString("size")
		reqjson.Start_time = hflag.GetTime("start_time")
		reqjson.End_time = hflag.GetTime("end_time")
		reqjson.Ignore_cache = hflag.GetBool("ignore_cache")
		reqjson.Field = hflag.GetString("field")
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
