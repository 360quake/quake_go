/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 15:36:10
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-01 16:03:48
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
)

func ServePost(query string, start string, size string, token string) {
	if query == "" || query == "?" {
		fmt.Println("No query specified")
		return
	}
	uri := "/search/quake_service"
	payload := "{\"query\":\"" + query +
		"\",\"start\":\"" + start + "\",\"size\":\"" + size +
		"\"}"
	tools.ApisPost(setting.URL+uri, payload, start, size, token)
}
func InfoGet(token string) {
	uri := "/user/info"
	tools.ApisGet(setting.URL+uri, token)
}
func FaviconPost(query string, start string, size string, token string) {
	uri := "/query/similar_icon/aggregation"
	tools.ApisGet(setting.URL+uri, token)
}
