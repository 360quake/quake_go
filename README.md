<!--
 * @Author: ph4nt0mer
 * @Date: 2022-09-01 18:39:52
 * @LastEditors: rootphantomer
 * @LastEditTime: 2022-09-07 14:55:42
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

- 2022-09-07 v1.4:

  - 替换 hflag 来解析命令行参数，使参数拓展开发更简单
  - 简化逻辑代码

- 2022-09-07 v1.3:

  - 支持-ic,-s,-e(ignore_cache,start_time,end_time)参数传入,start_time 初始化是当年年初 01-01，end_time 默认初始化是 now()
  - 优化发包的结构体解析

- 2022-09-06 v1.2:

  - 默认输出格式为 ip:port(见下图)
  - 新增-t=body 可以输出 body 信息

- 2022-09-06 v1.1:

  - 简化参数传参格式
  - 将 token 固化在当前目录

- 2022-09-01 v1.0:

  - 新增 info 接口功能
  - 新增 service 接口功能

## 使用

```bash
Starting Quake Cli...
usage: quake [pos] [-e,end_time time=2022-09-07 14:18:23] [-fe,field string] [-h,help bool] [-ic,ignore_cache bool=false] [-sz,size string=10] [-st,start string=0] [-s,start_time time=2022-01-01]

positional options:
       pos             [string]                    pos flag

options:
   -e, --end_time      [time=2022-09-07 14:18:23]  end time flag
  -fe, --field         [string]                    field flag
   -h, --help          [bool]                      show usage
  -ic, --ignore_cache  [bool=false]                ignore_cache value (true or false)
  -sz, --size          [string=10]                 size flag
  -st, --start         [string=0]                  start flag
   -s, --start_time    [time=2022-01-01]           start time flag
```

## 案例

![alt](./iShot_2022-09-07_14.19.19.jpg)
