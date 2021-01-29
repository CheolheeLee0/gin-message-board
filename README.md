# gin-message-board

基于Gin框架的网站留言系统

## 功能分析

- 让用户使用用户名和密码注册(仅限非登录用户)
- 让用户使用用户名和密码登录(仅限非登录用户)
- 允许用户注销(仅限已登录的用户)
- 允许用户创建留言(仅限登录用户)
- 在主页上显示所有留言的列表(适用于所有用户)
- 在自己的页面上显示留言(针对所有用户)

## 目录结构

```shell
├── gin-message-board	# 根目录
│   ├──  config		    # 配置目录
│   ├──  crontrollers	# 控制器目录	
│   ├──  database	    # 数据库目录
│   ├──  middlewares	# 中间件目录
│   ├──  models		    # 模型目录
│   ├──  templates	    # 模板目录
│   ├──  tests	        # 测试目录
│   ├──  tools	        # 其他工具目录
└── go.mod
└── go.sum	
└── main.go	            # 项目入口文件
```



##  技术栈

- 框架路由使用 [Gin](https://github.com/gin-gonic/gin) 路由
- 中间件使用 [Gin](https://github.com/gin-gonic/gin) 框架的中间件
- 数据库组件 [GORM](https://github.com/jinzhu/gorm)
- 模块测试gin

## TODO：

- 用户密码使用密文存储
- 持续集成
- 留言关联文章与用户
- 文章模型



