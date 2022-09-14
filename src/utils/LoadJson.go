/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 16:04:43
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-14 10:39:54
 * @FilePath: /quake_go/src/utils/LoadJson.go
 * @Description:解析json的逻辑代码
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package utils

import (
	"encoding/json"
	"fmt"
	. "quake/src/model"
)

func SeriveLoadJson(body string) (result ServiceJson) {
	var serivejson ServiceJson

	if err := json.Unmarshal([]byte(body), &serivejson); err == nil {
		result = serivejson
	} else {
		fmt.Println(err)
	}
	return
}
func HostLoadJson(body string) (result HostJson) {
	var hostjson HostJson

	if err := json.Unmarshal([]byte(body), &hostjson); err == nil {
		result = hostjson
	} else {
		fmt.Println(err)
	}
	return
}
func SeriveLoadJson2(body string) (result []map[string]string) {
	var servicemapjson map[string][]map[string]string

	if err := json.Unmarshal([]byte(body), &servicemapjson); err == nil {
		result = servicemapjson["data"]
	} else {
		fmt.Println(err)
	}
	return
}

func InfoLoadJson(body string) (result InfoJson) {
	var infojson InfoJson
	if err := json.Unmarshal([]byte(body), &infojson); err == nil {
		result = infojson
	} else {
		fmt.Println(err)
	}
	return
}

func InfoLoadJson2(body string) (map[string]interface{}, map[string]interface{}) {
	// 多层map，深度解析两次
	var infomapjson map[string]map[string]interface{}
	var infomapjson2 map[string]map[string]map[string]interface{}
	if err := json.Unmarshal([]byte(body), &infomapjson); err != nil {
	}
	if err := json.Unmarshal([]byte(body), &infomapjson2); err != nil {
	}
	// fmt.Println("###", infomapjson["data"])
	return infomapjson["data"], infomapjson2["data"]["user"]
}
