/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:14:52
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:15:13
 * @FilePath: /quake_go/src/model/aggregation.go
 * @Description:聚合数据筛选字段接口
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package model

type MyJsonName struct {
	Code    int64    `json:"code"`
	Data    []string `json:"data"`
	Message string   `json:"message"`
	Meta    struct{} `json:"meta"`
}
