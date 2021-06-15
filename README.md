# goal-http-template

HTTP API服务开发模板

## 准备工作：

1. 指定运行环境变量：APP_RUN_ENV
    
    * dev：开发环境
    * test: 测试环境
    * pro: 生产环境

2. 将 .env.template 复制为 .env.dev 文件 (与运行环境变量APP_RUN_ENV的值一致即可), 按实际需要修改配置文件内的各项配置


## 启动服务:

1. 启动前台API服务:

```shell script
$ APP_RUN_ENV=dev go run app.go
```

2. 启动后台管理API服务:

```shell script
$ APP_RUN_ENV=dev go run admin.go
```
