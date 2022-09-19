/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-06 10:04:23
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-19 10:52:11
 * @FilePath: /quake_go/utils/YamlAction.go
 * @Description:读写yaml
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package utils

import (
	"fmt"
	"io/ioutil"
	"os"

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

func ReadYaml(path string) (TokenInfo, bool) {
	// 读取yaml里面的内容
	var token TokenInfo
	token.Token = ""
	_, err := os.Stat(path)
	if err == nil {
		content, _ := ioutil.ReadFile(path)
		yaml.Unmarshal(content, &token)
		return token, true
	}
	if os.IsNotExist(err) { //如果返回的错误类型使用os.isNotExist()判断为true，说明文件或者文件夹不存在
		fmt.Println("!!!!please ./quake init token!!!!")
		return token, false
	}
	fmt.Println("!!!!please ./quake init token!!!!")
	return token, false
}
