/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 15:36:10
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-14 12:08:25
 * @FilePath: /quake_go/src/apis/api.go
 * @Description:封装请求接口
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"encoding/json"
	"fmt"
	. "quake/src/model"
	"quake/src/setting"
	"quake/src/tools"
	"quake/src/utils"
	"strconv"
	"strings"
)

func FilterableServiceGET(token string) {
	// 获取服务数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/filterable/field/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/filterable/field/quake_service"
	tools.Apis(setting.URL+uri, "GET", []byte{}, token)
}
func SearchServicePost(reqjson Reqjson, token string) {
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
	reqjson.Query = strings.ReplaceAll(reqjson.Query, " ", "")
	if reqjson.Query == "" || reqjson.Query == "?" {
		fmt.Println("No query specified")
		return
	}
	if reqjson.Query_txt != "" {
		bytedata, _ := utils.ReadLine(reqjson.Query_txt)
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
	}
	datajson, err := json.Marshal(reqjson)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", string(setting.URL+uri))
	fmt.Println(string(datajson))
	body := tools.Apis(setting.URL+uri, "POST", datajson, token)
	data_result := utils.RespLoadJson[SearchJson](body).Data
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
	tools.Apis(setting.URL+uri, "POST", query, token)
}
func AggregationServiceGet(token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.Apis(setting.URL+uri, "GET", []byte{}, token)
}
func AggregationServicePost(query []byte, start string, size string, token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.Apis(setting.URL+uri, "POST", query, token)
}

func InfoGet(token string) {
	// 个人信息接口
	uri := "/user/info"
	info := tools.Apis(setting.URL+uri, "GET", []byte{}, token)
	data_result, user_result := utils.InfoLoadJson[map[string]interface{}](info)
	fmt.Println("#用户名:", user_result["username"])
	fmt.Println("#邮  箱:", user_result["email"])
	fmt.Println("#手机:", data_result["mobile_phone"])
	fmt.Println("#月度积分:", data_result["month_remaining_credit"])
	fmt.Println("#长效积分:", data_result["constant_credit"])
	fmt.Println("#Token:", data_result["token"])

}
func FaviconPost(query string, token string) {
	uri := "/query/similar_icon/aggregation"
	tools.Apis(setting.URL+uri, "GET", []byte{}, token)
}

func HostSearchPost(reqjson Reqjson, token string) {
	//	curl -X POST "https://quake.360.cn/api/v3/search/quake_host" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//	     "query": "service: http",
	//	     "start": 20,
	//	     "size": 10,
	//	     "ignore_cache": false,
	//	     "start_time": "2021-01-01 00:00:00",
	//	     "end_time": "2021-02-01 00:00:00"
	//	}'

	uri := "/search/quake_host"
	reqjson.Query = strings.ReplaceAll(reqjson.Query, " ", "")
	if reqjson.Query == "" || reqjson.Query == "?" {
		fmt.Println("No query specified")
		return
	}
	if reqjson.Query_txt != "" {
		bytedata, _ := utils.ReadLine(reqjson.Query_txt)
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
	}
	datajson, err := json.Marshal(reqjson)
	if err != nil {
		panic(err)
	}
	fmt.Println("->", string(setting.URL+uri))
	fmt.Println(string(datajson))
	body := tools.Apis(setting.URL+uri, "POST", datajson, token)
	data_result := utils.RespLoadJson[SearchJson](body).Data
	for index, value := range data_result {
		fmt.Println(strconv.Itoa(index+1) + "# " + value.IP)
	}
}
