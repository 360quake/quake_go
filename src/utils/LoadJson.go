/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 16:04:43
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 18:10:35
 * @FilePath: /quake_go/src/utils/LoadJson.go
 * @Description:
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
		// fmt.Println(serivejson.Data)
		// data := serivejson.Data
		result = serivejson
	} else {
		fmt.Println(err)
	}
	return
}
