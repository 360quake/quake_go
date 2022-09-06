/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 10:49:31
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 15:53:00
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

func ApisPost(url string, payload string, start string, size string, token string) {
	var jsonStr = []byte(payload)
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-QuakeToken", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:")
	fmt.Println(string(body))
	if strings.Contains(string(body), "login") {
		fmt.Println("token expired,please init new token")
	}
}
func ApisGet(url string, token string) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-QuakeToken", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("result:")
	fmt.Println(string(body))
	if strings.Contains(string(body), "login") {
		fmt.Println("token expired,please init new token")
	}
}
