/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 17:58:31
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 17:58:42
 * @FilePath: /quake_go/src/model/reqjson.go
 * @Description:model
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package model

type Reqjson struct {
	Query        string `json:"query"`
	Start        string `json:"start"`
	Size         string `json:"size"`
	Ignore_cache bool   `json:"ignore_cache"`
	Start_time   string `json:"start_time"`
	End_time     string `json:"end_time"`
	Field        string `json:"field"`
}
