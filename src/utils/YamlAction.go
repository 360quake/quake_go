/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-06 10:04:23
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-06 14:06:08
 * @FilePath: /quake_go/src/utils/YamlAction.go
 * @Description:读写yaml
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package utils

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type TokenInfo struct {
	Token string `yaml:"token"`
}

func WriteYaml(path string, token string) {
	// 进行写yaml操作
	var tokeninfo TokenInfo
	tokeninfo.Token = token
	data, err := yaml.Marshal(tokeninfo)
	//config赋值777操作
	err = ioutil.WriteFile(path, data, 0777)
	if err = ioutil.WriteFile(path, data, 0777); err != nil {
		fmt.Printf("token init fail!")
	} else {
		fmt.Printf("token init success!")
	}

}

func checkError(err error) bool {
	// 判断是否有错误，有错误nil，返回false，无返回true
	if err != nil {
		return true
	}
	return false
}

func ReadYaml(path string) (token TokenInfo, status bool) {
	// 读取yaml里面的内容
	content, err := ioutil.ReadFile(path)
	status = checkError(err)
	err = yaml.Unmarshal(content, &token)
	status = checkError(err)
	return
}
