/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:04:59
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:05:04
 * @FilePath: /quake_go/src/apis/FaviconPost.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"quake/src/setting"
	"quake/src/utils"
)

func FaviconPost(query string, start string, size string, token string) {
	uri := "/query/similar_icon/aggregation"
	utils.ApisGet(setting.URL+uri, token)
}
