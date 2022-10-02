/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 15:36:10
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-19 10:56:23
 * @FilePath: /quake_go/utils/api.go
 * @Description:封装请求接口
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package utils

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
)

func FilterableServiceGET(token string) {
	// 获取服务数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/filterable/field/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/filterable/field/quake_service"
	Apis(URL+uri, "GET", []byte{}, token)
}
func SearchServicePost(reqjson Reqjson, token string) string {
	logcolor := color.New(color.FgGreen)
	// 服务数据实时查询接口
	// curl -X POST "https://quake.360.cn/api/v3/search/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//      "query": "service: http",
	//      "start": 20,
	//      "size": 10,
	//      "ignore_cache": false,
	//      "start_time": "2021-01-01 00:00:00",
	//      "end_time": "2021-02-01 00:00:00"
	// }'
	uri := "/search/quake_service"
	if reqjson.QueryTxt != "" {
		bytedata, _ := ReadLine(reqjson.QueryTxt)
		tmp := ""
		for _, v := range bytedata {
			if tmp == "" {
				tmp = v
			} else {
				tmp += " OR " + v
			}

		}
		// fmt.Println(tmp)
		reqjson.Query = tmp
	} else {
		//reqjson.Query = strings.ReplaceAll(reqjson.Query, " ", "")
		if reqjson.Query == "" || reqjson.Query == "?" {
			// fmt.Println("No query specified")
			// return
			panic("No query specified")
		}
	}
	datajson, err := json.Marshal(reqjson)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", string(URL+uri))
	logcolor.Println(string(datajson))
	body := Apis(URL+uri, "POST", datajson, token)
	return body

	// fields := strings.Split(reqjson.Field, ",")
	// for _, value := range fields {
	// 	if strings.Contains(value, "body") {
	// 		for _, value := range data_result {
	// 			fmt.Println(value)
	// 			return
	// 		}
	// 	}
	// }

}
func ScrollServicePost(query []byte, start string, size string, token string) {
	// 服务数据深度查询接口
	// curl -X POST "https://quake.360.cn/api/v3/scroll/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//     "query": "service: http",
	//     "size": 100,
	//     "ignore_cache": false,
	//     "start_time": "2021-01-07 00:13:14",
	//     "end_time": "2021-05-20 01:13:14"
	// }'
	uri := "/scroll/quake_service"
	Apis(URL+uri, "POST", query, token)
}
func AggregationServiceGet(token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	Apis(URL+uri, "GET", []byte{}, token)
}
func AggregationServicePost(query []byte, start string, size string, token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	Apis(URL+uri, "POST", query, token)
}

func InfoGet(token string) string {
	// 个人信息接口
	uri := "/user/info"
	info := Apis(URL+uri, "GET", []byte{}, token)
	return info

}
func FaviconPost(query string, token string) {
	uri := "/query/similar_icon/aggregation"
	Apis(URL+uri, "GET", []byte{}, token)
}

func HostSearchPost(reqjson Reqjson, token string) string {
	//	curl -X POST "https://quake.360.cn/api/v3/search/quake_host" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//	     "query": "service: http",
	//	     "start": 20,
	//	     "size": 10,
	//	     "ignore_cache": false,
	//	     "start_time": "2021-01-01 00:00:00",
	//	     "end_time": "2021-02-01 00:00:00"
	//	}'
	logcolor := color.New(color.FgGreen)
	uri := "/search/quake_host"
	if reqjson.QueryTxt != "" {
		bytedata, _ := ReadLine(reqjson.QueryTxt)
		tmp := ""
		for _, v := range bytedata {
			if tmp == "" {
				tmp = v
			} else {
				tmp += " OR " + v
			}

		}
		// fmt.Println(tmp)
		reqjson.Query = tmp
	} else {
		//reqjson.Query = strings.ReplaceAll(reqjson.Query, " ", "")
		if reqjson.Query == "" || reqjson.Query == "?" {
			// fmt.Println("No query specified")
			panic("No query specified")
		}
	}

	datajson, err := json.Marshal(reqjson)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", string(URL+uri))
	logcolor.Println(string(datajson))
	body := Apis(URL+uri, "POST", datajson, token)
	return body

}
