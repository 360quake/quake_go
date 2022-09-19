/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 16:04:43
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-15 10:04:31
 * @FilePath: /quake_go/src/utils/LoadJson.go
 * @Description:解析json的逻辑代码
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package utils

import (
	"encoding/json"
)

func RespLoadJson[T any](body string) (respjson T) {
	if err := json.Unmarshal([]byte(body), &respjson); err != nil {
		panic(err)
	}
	return
}

func InfoLoadJson(body string) (data map[string]interface{}, user map[string]interface{}) {
	// 多层map，深度解析两次
	var infomapjson map[string]interface{}
	if err := json.Unmarshal([]byte(body), &infomapjson); err != nil {
		panic(err)
	}
	data = infomapjson["data"].(map[string]interface{})
	user = data["user"].(map[string]interface{})
	return
}
