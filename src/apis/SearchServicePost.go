/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:03:14
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:03:28
 * @FilePath: /quake_go/src/apis/SearchServicePost.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"fmt"
	"quake/src/setting"
	"quake/src/utils"
)

func SearchServicePost(query string, start string, size string, token string) {
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
	utils.ApisPost(setting.URL+uri, payload, start, size, token)
}
