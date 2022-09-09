/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 15:36:10
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-09 18:22:25
 * @FilePath: /quake_go/src/apis/api.go
 * @Description:
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
)

func FilterableServiceGET(token string) {
	// 获取服务数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/filterable/field/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/filterable/field/quake_service"
	tools.ApisGet(setting.URL+uri, token)
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
		fmt.Println(tmp)
		reqjson.Query = tmp
	}
	// payload := "{\"query\":\"" + req.Query +
	// 	"\",\"start\":\"" + req.Start + "\",\"size\":\"" + req.Size +
	// 	"\"}"
	datajson, err := json.Marshal(reqjson)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(datajson))
	body := tools.ApisPost(setting.URL+uri, datajson, token)
	data_result := utils.SeriveLoadJson(body).Data
	for index, value := range data_result {
		fmt.Println(strconv.Itoa(index+1) + "# " + value.IP + ":" + strconv.Itoa(value.Port))
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
	tools.ApisPost(setting.URL+uri, query, token)
}
func AggregationServiceGet(token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.ApisGet(setting.URL+uri, token)
}
func AggregationServicePost(query []byte, start string, size string, token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	tools.ApisPost(setting.URL+uri, query, token)
}
func InfoGet(token string) {
	// 个人信息接口
	uri := "/user/info"
	info := tools.ApisGet(setting.URL+uri, token)
	resut := utils.InfoLoadJson(info)
	data := resut.Data
	fmt.Println("#用户名:", data.User.Username)
	fmt.Println("#邮  箱:", data.User.Email)
	fmt.Println("#手机:", data.MobilePhone)
	fmt.Println("#月度积分:", data.MonthRemainingCredit)
	fmt.Println("#长效积分:", data.ConstantCredit)
}
func InfoGet2(token string) {
	// 个人信息接口
	uri := "/user/info"
	info := tools.ApisGet(setting.URL+uri, token)
	data_result, user_result := utils.InfoLoadJson2(info)
	fmt.Println("#用户名:", user_result["username"])
	fmt.Println("#邮  箱:", user_result["email"])
	fmt.Println("#手机:", data_result["mobile_phone"])
	fmt.Println("#月度积分:", data_result["month_remaining_credit"])
	fmt.Println("#长效积分:", data_result["constant_credit"])
	fmt.Println("#Token:", data_result["token"])

}
func FaviconPost(query string, token string) {
	uri := "/query/similar_icon/aggregation"
	tools.ApisGet(setting.URL+uri, token)
}
