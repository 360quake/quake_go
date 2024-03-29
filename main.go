/*
 * @Author: ph4nt0mer
 * @Date: 2022-08-31 17:03:03
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-19 11:13:00
 * @FilePath: /quake_go/main.go
 * @Description:主函数
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package main

import (
	"github.com/360quake/quake_go/utils"
	"github.com/fatih/color"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hpifu/go-kit/hflag"

	"go.uber.org/automaxprocs/maxprocs"
)

var (
	reqjson      utils.Reqjson
	successColor = color.New(color.FgBlue)
	errorColor   = color.New(color.FgRed)
)

func main() {
	_, _ = maxprocs.Set(maxprocs.Logger(func(string, ...any) {}))

	action(hflagInit())
}

func hflagInit() (num int) {
	_, _ = successColor.Println("Starting Quake Cli...")
	hflag.AddFlag("start", "-st to start number", hflag.Shorthand("st"), hflag.Type("string"), hflag.DefaultValue("0"))
	hflag.AddFlag("size", "-sz to size number ", hflag.Shorthand("sz"), hflag.Type("string"), hflag.DefaultValue("10"))
	hflag.AddFlag("ignore_cache", "-ic true or false,default false", hflag.Shorthand("ic"), hflag.Type("bool"), hflag.DefaultValue("false"))
	hflag.AddFlag("start_time", "-s time flag , default time is time.now.year", hflag.Shorthand("s"), hflag.Type("time"), hflag.DefaultValue(strconv.Itoa(time.Now().Year())+"-01-01"))
	hflag.AddFlag("end_time", "-e time to end time flag", hflag.Shorthand("e"), hflag.Type("time"), hflag.DefaultValue(time.Now().Format("2006-01-02 15:04:05")))
	hflag.AddFlag("field", "-fe swich body,title,host,html_hash,x_powered_by  to show infomation", hflag.Shorthand("fe"), hflag.Type("string"), hflag.DefaultValue(""))
	hflag.AddFlag("file_txt", "-ft ./file.txt file to query search", hflag.Shorthand("ft"), hflag.Type("string"), hflag.DefaultValue(""))
	hflag.AddPosFlag("option", "init,info,search,host")
	hflag.AddPosFlag("args", "query value,example port:443", hflag.DefaultValue(""))
	if err := hflag.Parse(); err != nil {
		panic(err)
	}
	num = len(os.Args)
	if num < 2 {
		_, _ = errorColor.Println("./quake -h get help!")
		os.Exit(0)
	}
	return
}

func action(num int) {
	reqjson.Query = hflag.GetString("args")
	reqjson.Start = hflag.GetString("start")
	reqjson.Size = hflag.GetString("size")
	reqjson.StartTime = hflag.GetTime("start_time")
	reqjson.EndTime = hflag.GetTime("end_time")
	reqjson.IgnoreCache = hflag.GetBool("ignore_cache")
	reqjson.Field = hflag.GetString("field")
	reqjson.QueryTxt = hflag.GetString("file_txt")
	if sizelen, _ := strconv.Atoi(reqjson.Size); sizelen > 50 {
		_, _ = errorColor.Println("size only less than or equal to 50")
		return
	}
	switch strings.ToLower(hflag.GetString("option")) {
	case "version":
		_, _ = successColor.Printf("Clash %s %s %s with %s \n", "2.0.3", runtime.GOOS, runtime.GOARCH, runtime.Version())
	case "init":
		if num < 3 {
			_, _ = errorColor.Println("!!!!token is empty !!!!")
			return
		}
		utils.WriteYaml("./config.yaml", reqjson.Query)
	case "info":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			os.Exit(0)
		}
		info := utils.InfoGet(token.Token)
		dataResult, userResult := utils.InfoLoadJson(info)
		_, _ = successColor.Println("#用户名:", userResult["username"])
		_, _ = successColor.Println("#邮  箱:", userResult["email"])
		_, _ = successColor.Println("#手机:", dataResult["mobile_phone"])
		_, _ = successColor.Println("#月度积分:", dataResult["month_remaining_credit"])
		_, _ = successColor.Println("#长效积分:", dataResult["constant_credit"])
		_, _ = successColor.Println("#Token:", dataResult["token"])
	case "search":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		body := utils.SearchServicePost(reqjson, token.Token)
		dataResult := utils.RespLoadJson[utils.SearchJson](body).Data
		if reqjson.Field != "" && reqjson.Field != "ip,port" {
			for index, value := range dataResult {
				if value.Service.HTTP[reqjson.Field] == nil {
					_, _ = successColor.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + "  " + strconv.Itoa(value.Port))

				} else {
					_, _ = successColor.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + strconv.Itoa(value.Port) + "  " + value.Service.HTTP[reqjson.Field].(string))

				}
			}
		} else {
			for index, value := range dataResult {
				_, _ = successColor.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + strconv.Itoa(value.Port))
			}
		}
	case "host":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			os.Exit(0)
		}
		body := utils.HostSearchPost(reqjson, token.Token)
		dataResult := utils.RespLoadJson[utils.SearchJson](body).Data
		for index, value := range dataResult {
			_, _ = successColor.Println(strconv.Itoa(index+1) + "# " + value.IP)
		}
	// case "favicon":
	// 	fmt.Println("favicon相似度接口待完成。。。")
	// 	token, status := utils.ReadYaml("./config.yaml")
	// 	if !status {
	// 		return
	// 	}
	// 	apis.HostSearchPost(reqjson, token.Token)
	// case "domain":
	// 	fmt.Println("domain ")
	default:
		_, _ = errorColor.Println("args failed !!!!")
	}
}
