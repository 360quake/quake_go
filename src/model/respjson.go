/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 18:10:53
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-14 10:39:09
 * @FilePath: /quake_go/src/model/respjson.go
 * @Description:返回包结构体
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */

package model

type ServiceJson struct {
	Code int `json:"code"`
	Data []struct {
		Asn int `json:"asn"`
		Cdn struct {
			Domain string `json:"domain"`
			IsCdn  bool   `json:"is_cdn"`
		} `json:"cdn"`
		Components []map[string]interface{} `json:"Components"`
		// Components []struct {
		// 	ID             string   `json:"id"`
		// 	ProductCatalog []string `json:"product_catalog"`
		// 	ProductLevel   string   `json:"product_level"`
		// 	ProductNameCn  string   `json:"product_name_cn"`
		// 	ProductNameEn  string   `json:"product_name_en"`
		// 	ProductType    []string `json:"product_type"`
		// 	ProductVendor  string   `json:"product_vendor"`
		// 	Version        string   `json:"version"`
		// } `json:"components"`
		Hostname string   `json:"hostname"`
		ID       string   `json:"id"`
		Images   []string `json:"images"`
		IP       string   `json:"ip"`
		IsIpv6   bool     `json:"is_ipv6"`
		// Location map[string]string `json:"location"`
		Location struct {
			CityCn      string  `json:"city_cn"`
			CityEn      string  `json:"city_en"`
			CountryCn   string  `json:"country_cn"`
			CountryCode string  `json:"country_code"`
			CountryEn   string  `json:"country_en"`
			DistrictCn  string  `json:"district_cn"`
			DistrictEn  string  `json:"district_en"`
			Isp         string  `json:"isp"`
			ProvinceCn  string  `json:"province_cn"`
			ProvinceEn  string  `json:"province_en"`
			Radius      float64 `json:"radius"`
			SceneCn     string  `json:"scene_cn"`
			SceneEn     string  `json:"scene_en"`
		} `json:"location"`
		Org       string `json:"org"`
		OsName    string `json:"os_name"`
		OsVersion string `json:"os_version"`
		Port      int    `json:"port"`
		Service   struct {
			Banner string                 `json:"banner"`
			Cert   string                 `json:"cert"`
			HTTP   map[string]interface{} `json:"http"`
			// HTTP   struct {
			// 	Body    string `json:"body"`
			// 	Favicon struct {
			// 		Data     string `json:"data"`
			// 		Hash     string `json:"hash"`
			// 		Location string `json:"location"`
			// 	} `json:"favicon"`
			// 	Host            string `json:"host"`
			// 	HTMLHash        string `json:"html_hash"`
			// 	MetaKeywords    string `json:"meta_keywords"`
			// 	Path            string `json:"path"`
			// 	ResponseHeaders string `json:"response_headers"`
			// 	Server          string `json:"server"`
			// 	StatusCode      int    `json:"status_code"`
			// 	Title           string `json:"title"`
			// 	XPoweredBy      string `json:"x_powered_by"`
			// } `json:"http"`
			Name     string `json:"name"`
			Product  string `json:"product"`
			Response string `json:"response"`
			Version  string `json:"version"`
		} `json:"service"`
		SysTag    []string `json:"sys_tag"`
		Time      string   `json:"time"`
		Transport string   `json:"transport"`
	} `json:"data"`
	Message string `json:"message"`
	Meta    struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
type HostJson struct {
	Code int `json:"code"`
	Data []struct {
		Asn int `json:"asn"`
		Cdn struct {
			Domain string `json:"domain"`
			IsCdn  bool   `json:"is_cdn"`
		} `json:"cdn"`
		Components []map[string]interface{} `json:"Components"`
		// Components []struct {
		// 	ID             string   `json:"id"`
		// 	ProductCatalog []string `json:"product_catalog"`
		// 	ProductLevel   string   `json:"product_level"`
		// 	ProductNameCn  string   `json:"product_name_cn"`
		// 	ProductNameEn  string   `json:"product_name_en"`
		// 	ProductType    []string `json:"product_type"`
		// 	ProductVendor  string   `json:"product_vendor"`
		// 	Version        string   `json:"version"`
		// } `json:"components"`
		Hostname string   `json:"hostname"`
		ID       string   `json:"id"`
		Images   []string `json:"images"`
		IP       string   `json:"ip"`
		IsIpv6   bool     `json:"is_ipv6"`
		// Location map[string]string `json:"location"`
		Location struct {
			CityCn      string  `json:"city_cn"`
			CityEn      string  `json:"city_en"`
			CountryCn   string  `json:"country_cn"`
			CountryCode string  `json:"country_code"`
			CountryEn   string  `json:"country_en"`
			DistrictCn  string  `json:"district_cn"`
			DistrictEn  string  `json:"district_en"`
			Isp         string  `json:"isp"`
			ProvinceCn  string  `json:"province_cn"`
			ProvinceEn  string  `json:"province_en"`
			Radius      float64 `json:"radius"`
			SceneCn     string  `json:"scene_cn"`
			SceneEn     string  `json:"scene_en"`
		} `json:"location"`
		Org       string `json:"org"`
		OsName    string `json:"os_name"`
		OsVersion string `json:"os_version"`
		Port      int    `json:"port"`
		Service   struct {
			Banner string                 `json:"banner"`
			Cert   string                 `json:"cert"`
			HTTP   map[string]interface{} `json:"http"`
			// HTTP   struct {
			// 	Body    string `json:"body"`
			// 	Favicon struct {
			// 		Data     string `json:"data"`
			// 		Hash     string `json:"hash"`
			// 		Location string `json:"location"`
			// 	} `json:"favicon"`
			// 	Host            string `json:"host"`
			// 	HTMLHash        string `json:"html_hash"`
			// 	MetaKeywords    string `json:"meta_keywords"`
			// 	Path            string `json:"path"`
			// 	ResponseHeaders string `json:"response_headers"`
			// 	Server          string `json:"server"`
			// 	StatusCode      int    `json:"status_code"`
			// 	Title           string `json:"title"`
			// 	XPoweredBy      string `json:"x_powered_by"`
			// } `json:"http"`
			Name     string `json:"name"`
			Product  string `json:"product"`
			Response string `json:"response"`
			Version  string `json:"version"`
		} `json:"service"`
		SysTag    []string `json:"sys_tag"`
		Time      string   `json:"time"`
		Transport string   `json:"transport"`
	} `json:"data"`
	Message string `json:"message"`
	Meta    struct {
		Pagination struct {
			Count     int `json:"count"`
			PageIndex int `json:"page_index"`
			PageSize  int `json:"page_size"`
			Total     int `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
type InfoJson struct {
	Code int `json:"code"`
	Data struct {
		AvatarID       string `json:"avatar_id"`
		BanStatus      string `json:"ban_status"`
		Baned          bool   `json:"baned"`
		ConstantCredit int    `json:"constant_credit"`
		Credit         int    `json:"credit"`
		Disable        struct {
			DisableTime interface{} `json:"disable_time"`
			StartTime   interface{} `json:"start_time"`
		} `json:"disable"`
		EnterpriseInformation struct {
			Email  interface{} `json:"email"`
			Name   interface{} `json:"name"`
			Status string      `json:"status"`
		} `json:"enterprise_information"`
		FreeQueryAPICount  int    `json:"free_query_api_count"`
		ID                 string `json:"id"`
		InvitationCodeInfo struct {
			Code                string `json:"code"`
			InviteAcquireCredit int    `json:"invite_acquire_credit"`
			InviteNumber        int    `json:"invite_number"`
		} `json:"invitation_code_info"`
		IsCashedInvitationCode    bool   `json:"is_cashed_invitation_code"`
		MobilePhone               string `json:"mobile_phone"`
		MonthRemainingCredit      int    `json:"month_remaining_credit"`
		PersistentCredit          int    `json:"persistent_credit"`
		PersonalInformationStatus bool   `json:"personal_information_status"`
		PrivacyLog                struct {
			AnonymousModel bool   `json:"anonymous_model"`
			QuakeLogStatus bool   `json:"quake_log_status"`
			QuakeLogTime   string `json:"quake_log_time"`
			Status         bool   `json:"status"`
			Time           string `json:"time"`
		} `json:"privacy_log"`
		Role []struct {
			Credit   int    `json:"credit"`
			Fullname string `json:"fullname"`
			Priority int    `json:"priority"`
		} `json:"role"`
		RoleValidity struct {
		} `json:"role_validity"`
		Source string `json:"source"`
		Time   string `json:"time"`
		Token  string `json:"token"`
		User   struct {
			Email    string   `json:"email"`
			Fullname string   `json:"fullname"`
			Group    []string `json:"group"`
			ID       string   `json:"id"`
			Username string   `json:"username"`
		} `json:"user"`
	} `json:"data"`
	Message string   `json:"message"`
	Meta    struct{} `json:"meta"`
}
