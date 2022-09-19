/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-19 10:54:59
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-19 10:55:01
 * @FilePath: /quake_go/utils/test.go
 * @Description:
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package utils

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
