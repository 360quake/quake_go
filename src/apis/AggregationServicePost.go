/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:05:37
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:05:45
 * @FilePath: /quake_go/src/apis/AggregationServicePost.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"quake/src/setting"
	"quake/src/utils"
)

func AggregationServicePost(query string, start string, size string, token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	utils.ApisPost(setting.URL+uri, query, start, size, token)
}
