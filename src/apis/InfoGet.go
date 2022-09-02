/*
 * @Author: ph4nt0mer rootphantomy@hotmail.com
 * @Date: 2022-09-02 17:05:18
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 17:05:22
 * @FilePath: /quake_go/src/apis/InfoGet.go
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
 */
package apis

import (
	"quake/src/setting"
	"quake/src/utils"
)

func InfoGet(token string) {
	// 个人信息接口
	uri := "/user/info"
	utils.ApisGet(setting.URL+uri, token)
}
