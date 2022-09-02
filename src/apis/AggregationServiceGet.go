/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:05:52
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:09:42
 * @FilePath: /quake_go/src/apis/AggregationServiceGet.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"quake/src/setting"
	"quake/src/utils"
)

func AggregationServiceGet(token string) {
	// 获取聚合数据筛选字段
	// curl -X GET "https://quake.360.cn/api/v3/aggregation/quake_service" -H "X-QuakeToken: d17140ae-xxxx-xxx-xxxx-c0818b2bbxxx"
	uri := "/aggregation/quake_service"
	utils.ApisGet(setting.URL+uri, token)
}
