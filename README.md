## hertz框架



## 目录结构


```bash
.
├── app
│   ├── http http模块
│   │   ├── controllers 控制器
│   │   ├── middleware 中间件
│   │   └── validator 验证器
│   ├── logic 逻辑模块
│   └── models 模型
├── app.ini 配置文件
├── config 配置文件目录
├── datebase 数据库
│   ├── model.go 基本模型
│   └── softDeletes.go 软删除
├── go.mod
├── go.sum
├── main.go 服务入口文件
├── migrate.go 迁移脚本
├── pkg
│   └── setting
│       └── setting.go
├── README.md
└── routes 路由
    └── api.go

```



## 模块

- gin框架：github.com/gin-gonic/gin
- 配置读取：github.com/go-ini/ini
- orm：github.com/jinzhu/gorm
- html模板： golang 的html/template库
- 热更新：https://github.com/gravityblast/fresh


## Feature

- 中间件
- rpc
- http、rpc分块，按需加载
- tool管理
- 命令：平滑重启
- view


## 规范


## 使用

### 迁移

```bash
go run migrate.go
```
