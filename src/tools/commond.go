/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 10:49:31
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 18:12:57
 * @FilePath: /quake_go/src/tools/commond.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func ApisPost(url string, data string, token string) string {
	var jsonStr = []byte(data)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("X-QuakeToken", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	fmt.Println("result:")
	// fmt.Println(string(body))
	if strings.Contains(string(body), "quake/login") {
		fmt.Println("token expired,please init new token")
	}
	return string(body)
}
func ApisGet(url string, token string) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("X-QuakeToken", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	fmt.Println("result:")
	// fmt.Println(string(body))
	if strings.Contains(string(body), "quake/login") {
		fmt.Println("token expired,please init new token")
	}
	return string(body)
}
