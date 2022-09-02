<!--
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 18:39:52
 * @LastEditors: ph4nt0mer
 * @LastEditTime: 2022-09-02 15:42:38
 * @FilePath: /quake_go/README.md
 * @Description:
 *
 * Copyright (c) 2022 by ph4nt0mer, All Rights Reserved.
-->

# Quake_go

Quake Command-Line Application With Golang

## rust 版本

rust 项目代码 移步这里——https://github.com/360quake/quake_rs/

## 安装

1. 直接下载即可使用
2. 或者本地编译：

   ```bash
   // 安装golang后编译
   go  build .
   ```

## 更新日志

- 2022-09-01 v1.0:
  - 新增 info 接口功能
  - 新增 service 接口功能

## 使用

```bash
Starting Quake Cli...
example:
 ./quake -token=<token> -model=<model> -query=<query>

Usage of ./quake_go:
  -model string
        model String value
  -query string
        query String value
  -size string
        size String value (default "10")
  -start string
        start String value (default "0")
  -token string
        token Sting value
```
