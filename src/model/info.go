/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 16:52:51
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:07:12
 * @FilePath: /quake_go/src/model/info.go
 * @Description:定义info的json返回值结构
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package model

type Infojson struct {
	Code int64 `json:"code"`
	Data struct {
		BanStatus             string `json:"ban_status"`
		Baned                 bool   `json:"baned"`
		Credit                int64  `json:"credit"`
		EnterpriseInformation struct {
			Email  interface{} `json:"email"`
			Name   interface{} `json:"name"`
			Status string      `json:"status"`
		} `json:"enterprise_information"`
		ID                        string `json:"id"`
		MobilePhone               string `json:"mobile_phone"`
		PersistentCredit          int64  `json:"persistent_credit"`
		PersonalInformationStatus bool   `json:"personal_information_status"`
		PrivacyLog                struct {
			Status bool        `json:"status"`
			Time   interface{} `json:"time"`
		} `json:"privacy_log"`
		Role []struct {
			Credit   int64  `json:"credit"`
			Fullname string `json:"fullname"`
			Priority int64  `json:"priority"`
		} `json:"role"`
		Source string `json:"source"`
		Token  string `json:"token"`
		User   struct {
			Email    string `json:"email"`
			Fullname string `json:"fullname"`
			ID       string `json:"id"`
			Username string `json:"username"`
		} `json:"user"`
	} `json:"data"`
	Message string   `json:"message"`
	Meta    struct{} `json:"meta"`
}
