## fullstack

基于Gin + Vue +minigui UI的前后端分离框架，系统初始化简单，只需要配置文件中，修改数据库连接，系统支持多指令操作

## 获取代码

```bash
git clone https://github.com/longzhoufeng/fullstack.git
```

## 服务端启动说明

```bash
# 进入 fullstack 项目
cd ./fullstack
```

## 初始化数据库，以及服务启动

```bash
 # 首次配置需要初始化数据库资源信息
 
 # 1、macOS or linux 下使用
 $ ./fullstack migrate -c=config/settings.dev.yml
 
 # 2、⚠️注意:windows 下使用
$ fullstack.exe migrate -c=config/settings.dev.yml
or
$ go run main.go migrate -c=config/settings.dev.yml

# 3、启动项目，也可以用IDE进行调试,在IDE的Edit configurations=>Go build=>Configuration=>Program arguments里设置如下参数

./fullstack server -c config/settings.dev.yml

# 4、⚠️注意:windows 下使用
server -c config/settings.dev.yml
```

Copyright (c) 2021 longzhoufeng
