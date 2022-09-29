/*
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 10:49:31
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-19 10:51:48
 * @FilePath: /quake_go/utils/commond.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package utils

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"io"
	"net/http"
	"strings"
)

func Apis(url string, method string, data []byte, token string) string {
	errcolor := color.New(color.FgRed)
	client := &http.Client{}
	var req *http.Request
	var err error
	switch method {
	case "POST":
		req, err = http.NewRequest("POST", url, bytes.NewBuffer(data))
	case "GET":
		req, err = http.NewRequest("GET", url, nil)
	}
	if err != nil {
		errcolor.Println(err)
		return err.Error()
	}
	req.Header.Add("X-QuakeToken", token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		errcolor.Println(err)
		return err.Error()
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			errcolor.Println(err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		errcolor.Println(err)
		return err.Error()
	}
	fmt.Println("result:")
	// fmt.Println(string(body))
	if strings.Contains(string(body), "quake/login") {
		errcolor.Println("token expired,please init new token")
	}
	return string(body)
}
