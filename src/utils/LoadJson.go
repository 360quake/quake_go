/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 16:04:43
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-14 11:51:59
 * @FilePath: /quake_go/src/utils/LoadJson.go
 * @Description:解析json的逻辑代码
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package utils

import (
	"encoding/json"
	"fmt"
)

func RespLoadJson[T any](body string) (result T) {
	var respjson T
	if err := json.Unmarshal([]byte(body), &respjson); err == nil {
		result = respjson
	} else {
		fmt.Println(err)
	}
	return
}

func InfoLoadJson[T any](body string) (T, T) {
	// 多层map，深度解析两次
	var infomapjson map[string]T
	var infomapjson2 map[string]map[string]T
	if err := json.Unmarshal([]byte(body), &infomapjson); err != nil {
	}
	if err := json.Unmarshal([]byte(body), &infomapjson2); err != nil {
	}
	// fmt.Println("###", infomapjson["data"])
	return infomapjson["data"], infomapjson2["data"]["user"]
}
