/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:13:28
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:13:32
 * @FilePath: /quake_go/src/model/search.go
 * @Description:实时查询接口
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package model

type SearchJson struct {
	Code    int64         `json:"code"`
	Data    []interface{} `json:"data"`
	Message string        `json:"message"`
	Meta    struct {
		Pagination struct {
			Count     int64 `json:"count"`
			PageIndex int64 `json:"page_index"`
			PageSize  int64 `json:"page_size"`
			Total     int64 `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
