/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 17:58:31
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-07 16:19:32
 * @FilePath: /quake_go/src/model/reqjson.go
 * @Description:model
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package model

import . "time"

type Reqjson struct {
	Query        string `json:"query"`
	Start        string `json:"start,omitempty"`
	Size         string `json:"size,omitempty"`
	Ignore_cache bool   `json:"ignore_cache,omitempty"`
	Start_time   Time   `json:"start_time,omitempty"`
	End_time     Time   `json:"end_time,omitempty"`
	Field        string `json:"-"`
	Query_txt    string `json:"-"`
}
