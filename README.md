# goal-http-template

HTTP API服务开发模板

## 环境变量说明:

运行本服务需要指定两个环境变量: 


* ENV_NAME: env配置文件名称, 如未指定, 则默认使用根目录 .env 文件

* APP_SERVICE: 服务的类型, 有两个值.

    * admin 后台管理接口服务; 
    * api 前台API接口服务


## 本地开发准备工作：

1. 将 .env.template 复制为 .env 文件, 按实际需要修改配置文件内的各项配置;


## 运行本地开发服务:


1. 启动前台API服务:

```shell script
$ ENV_NAME=.env;APP_SERVICE=api go run main.go
```

2. 启动后台管理API服务:

```shell script
$ ENV_NAME=.env;APP_SERVICE=admin go run main.go
```

### Docker-compose快速启动

```shell script
$ docker-compose up -d 
```


## Docker分步操作

### 生成Image

```shell script
$ make build
```

### 运行Docker容器

运行前台API服务
```shell script
$ make run-api
```

运行后台API服务

```shell script
$ make run-admin
```

### 测试服务是否正常启动: 

{IP}/api/ping

