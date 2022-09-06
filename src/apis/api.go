/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 15:36:10
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 16:44:26
 * @FilePath: /quake_go/src/apis/api.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"fmt"
	"quake/src/setting"
	"quake/src/tools"
	"quake/src/utils"
	"strconv"
)

func FilterableServiceGET(token string) {
	// 获取服务数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/filterable/field/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/filterable/field/quake_service"
	tools.ApisGet(setting.URL+uri, token)
}
func SearchServicePost(query string, start string, size string, token string, field string) {
	// 服务数据实时查询接口
	// curl -X POST "https://quake.360.cn/api/v3/search/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//      "query": "service: http",
	//      "start": 20,
	//      "size": 10,
	//      "ignore_cache": false,
	//      "start_time": "2021-01-01 00:00:00",
	//      "end_time": "2021-02-01 00:00:00"
	// }'
	if query == "" || query == "?" {
		fmt.Println("No query specified")
		return
	}
	uri := "/search/quake_service"
	payload := "{\"query\":\"" + query +
		"\",\"start\":\"" + start + "\",\"size\":\"" + size +
		"\"}"

	body := tools.ApisPost(setting.URL+uri, payload, start, size, token)
	resut := utils.SeriveLoadJson(body)
	data := resut.Data
	// if field != "" {
	// 	fields := strings.Split(field, ",")
	// 	for _, fields_value := range fields {
	// 		for _, value := range data {
	// 			fmt.Println(value.fields_value)
	// 		}
	// 	}
	// }
	for _, value := range data {
		fmt.Println(value.IP + ":" + strconv.Itoa(value.Port))
	}

}
func ScrollServicePost(query string, start string, size string, token string) {
	// 服务数据深度查询接口
	// curl -X POST "https://quake.360.cn/api/v3/scroll/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx" -H "Content-Type: application/json" -d '{
	//     "query": "service: http",
	//     "size": 100,
	//     "ignore_cache": false,
	//     "start_time": "2021-01-07 00:13:14",
	//     "end_time": "2021-05-20 01:13:14"
	// }'
	uri := "/scroll/quake_service"
	tools.ApisPost(setting.URL+uri, query, start, size, token)
}
func AggregationServiceGet(token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.ApisGet(setting.URL+uri, token)
}
func AggregationServicePost(query string, start string, size string, token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.ApisPost(setting.URL+uri, query, start, size, token)
}
func InfoGet(token string) {
	// 个人信息接口
	uri := "/user/info"
	info := tools.ApisGet(setting.URL+uri, token)
	fmt.Println(info)
}
func FaviconPost(query string, start string, size string, token string) {
	uri := "/query/similar_icon/aggregation"
	tools.ApisGet(setting.URL+uri, token)
}
