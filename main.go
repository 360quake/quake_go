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
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/360quake/quake_go/utils"

	"github.com/hpifu/go-kit/hflag"
)

func main() {
	hflagInit()
}

func hflagInit() {
	fmt.Println("Starting Quake Cli...")
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
	num := len(os.Args)
	if num < 2 {
		fmt.Println("./quake -h get help!")
		return
	}
	var reqjson utils.Reqjson
	reqjson.Query = hflag.GetString("args")
	reqjson.Start = hflag.GetString("start")
	reqjson.Size = hflag.GetString("size")
	reqjson.Start_time = hflag.GetTime("start_time")
	reqjson.End_time = hflag.GetTime("end_time")
	reqjson.Ignore_cache = hflag.GetBool("ignore_cache")
	reqjson.Field = hflag.GetString("field")
	reqjson.Query_txt = hflag.GetString("file_txt")
	if sizelen, _ := strconv.Atoi(reqjson.Size); sizelen > 50 {
		fmt.Println("size only less than or equal to 50")
		return
	}
	switch strings.ToLower(hflag.GetString("option")) {
	case "version":
		fmt.Println("version:2.0")
	case "init":
		if num < 3 {
			fmt.Println("!!!!token is empty !!!!")
			return
		}
		utils.WriteYaml("./config.yaml", reqjson.Query)
	case "info":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		info := utils.InfoGet(token.Token)
		data_result, user_result := utils.InfoLoadJson(info)
		fmt.Println("#用户名:", user_result["username"])
		fmt.Println("#邮  箱:", user_result["email"])
		fmt.Println("#手机:", data_result["mobile_phone"])
		fmt.Println("#月度积分:", data_result["month_remaining_credit"])
		fmt.Println("#长效积分:", data_result["constant_credit"])
		fmt.Println("#Token:", data_result["token"])
	case "search":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		body := utils.SearchServicePost(reqjson, token.Token)
		data_result := utils.RespLoadJson[utils.SearchJson](body).Data
		if reqjson.Field != "" && reqjson.Field != "ip,port" {
			for index, value := range data_result {
				if value.Service.HTTP[reqjson.Field] == nil {
					fmt.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + "  " + strconv.Itoa(value.Port))
				} else {
					fmt.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + strconv.Itoa(value.Port) + "  " + value.Service.HTTP[reqjson.Field].(string))
				}
			}
		} else {
			for index, value := range data_result {
				fmt.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + strconv.Itoa(value.Port))
			}
		}
	case "host":
		token, status := utils.ReadYaml("./config.yaml")
		if !status {
			return
		}
		body := utils.HostSearchPost(reqjson, token.Token)
		data_result := utils.RespLoadJson[utils.SearchJson](body).Data
		for index, value := range data_result {
			fmt.Println(strconv.Itoa(index+1) + "# " + value.IP)
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
		fmt.Println("args failed !!!!")
	}
}
