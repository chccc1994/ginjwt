## JWT Demo
[1. jwt验证](https://jwt.io/)

- 任务时间：11.2
> gin: jwtdemo 大概框架搭建
```text
.
├── conf/
│   └── app.ini
├── docs/
│   └── mysql.sql
├── go.mod
├── go.sum
├── main.go
├── models/
├── pkg/
│   ├── e/
│   │   └── code.go
│   ├── logging/
│   │   ├── file.go
│   │   └── log.go
│   └── setting/
│       └── setting.go
├── README.md
└── routers/
    ├── api/
    │   └── auth.go
    └── router.go
```

+ 任务时间：11.3
>任务进程：log日志实现、路由、token生成

+ 存在问题
1. endless执行监听，再次执行时没有显示端口号。
2. 