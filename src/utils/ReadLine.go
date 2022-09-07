/*
 * @Author: rootphantomer rootphantomy@hotmail.com
 * @Date: 2022-09-07 16:27:34
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-07 16:31:15
 * @FilePath: /quake_go/src/utils/ReadLine.go
 * @Description:逐行读取内容
 *
 * Copyright (c) 2022 by rootphantomer, All Rights Reserved.
 */
package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func ReadLine(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	buf := bufio.NewReader(f)
	var result []string
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				return result, nil
			}
			return nil, err
		}
		result = append(result, line)
	}
}
