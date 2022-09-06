/*
 * @Author: rootphantomer zhoufei1@360.cn
 * @Date: 2022-09-06 18:10:53
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 18:10:57
 * @FilePath: /quake_go/src/model/respjson.go
 * @Description:
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */

package model

type ServiceJson struct {
	Code int64 `json:"code"`
	Data []struct {
		Asn int64 `json:"asn"`
		Cdn struct {
			Domain string `json:"domain"`
			IsCdn  bool   `json:"is_cdn"`
		} `json:"cdn"`
		Components []struct {
			ID             string   `json:"id"`
			ProductCatalog []string `json:"product_catalog"`
			ProductLevel   string   `json:"product_level"`
			ProductNameCn  string   `json:"product_name_cn"`
			ProductNameEn  string   `json:"product_name_en"`
			ProductType    []string `json:"product_type"`
			ProductVendor  string   `json:"product_vendor"`
			Version        string   `json:"version"`
		} `json:"components"`
		Hostname string        `json:"hostname"`
		ID       string        `json:"id"`
		Images   []interface{} `json:"images"`
		IP       string        `json:"ip"`
		IsIpv6   bool          `json:"is_ipv6"`
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
		Port      int64  `json:"port"`
		Service   struct {
			Banner string `json:"banner"`
			Cert   string `json:"cert"`
			HTTP   struct {
				Body    string `json:"body"`
				Favicon struct {
					Data     string `json:"data"`
					Hash     string `json:"hash"`
					Location string `json:"location"`
				} `json:"favicon"`
				Host            string `json:"host"`
				HTMLHash        string `json:"html_hash"`
				MetaKeywords    string `json:"meta_keywords"`
				Path            string `json:"path"`
				ResponseHeaders string `json:"response_headers"`
				Server          string `json:"server"`
				StatusCode      int64  `json:"status_code"`
				Title           string `json:"title"`
				XPoweredBy      string `json:"x_powered_by"`
			} `json:"http"`
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
			Count     int64 `json:"count"`
			PageIndex int64 `json:"page_index"`
			PageSize  int64 `json:"page_size"`
			Total     int64 `json:"total"`
		} `json:"pagination"`
	} `json:"meta"`
}
