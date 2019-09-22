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

## jwt验证流程

- 1.登录获取token，缓存到cookie（后台管理系统）或前端自己保存（前后端分离）
- 2.将请求header带上token请求
- 3.后端中间件进行验证


token登录刷新时间有效为7天
token登录有效时间为2小时

token验证
- 1.当token小于2小时，直接允许通过
- 2.当token大于2小时小于7天时候，自动（后台）或接口（app）进行刷新接口
- 3.当token大于7天的时候，token无效，提示重新登录


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


### 测试部署

git pull 更新
swag init 生成文档
重启服务
