# Go Gin App [![rcard](https://goreportcard.com/badge/github.com/luanluanxu/go-gin-app)](https://goreportcard.com/report/github.com/luanluanxu/go-gin-app) [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/luanluanxu/go-gin-app) [![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/EDDYCJY/go-gin-app/master/LICENSE)

An example of gin contains many useful features

## Tips:

This project is the multiple program version fork from `go-gin-example`

Whats needs attention：
- The `config`,`runtime` directly now in the `cmd/my_app`
- The internal files now in the `internal/my_app`
- When you use `gin-swagger` to generate Swagger files,you need add 2 args such `swag -d internal/my_app -o internal/my_app`


[简体中文](https://github.com/luanluanxu/go-gin-app/blob/master/README_ZH.md)

## Installation
```
$ go get github.com/luanluanxu/go-gin-app
```

## How to run

### Required

- Mysql
- Redis

### Ready

Create a **blog database** and import [SQL](https://github.com/luanluanxu/go-gin-app/blob/master/docs/sql/blog.sql)

### Conf

You should modify `conf/app.ini`

```
[database]
Type = mysql
User = root
Password =
Host = 127.0.0.1:3306
Name = blog
TablePrefix = blog_

[redis]
Host = 127.0.0.1:6379
Password =
MaxIdle = 30
MaxActive = 30
IdleTimeout = 200
...
```

### Run
```
$ cd $GOPATH/src/go-gin-app

$ go run main.go 
```

Project information and existing API

```
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /auth                     --> github.com/luanluanxu/go-gin-app/routers/api.GetAuth (3 handlers)
[GIN-debug] GET    /swagger/*any             --> github.com/luanluanxu/go-gin-app/vendor/github.com/swaggo/gin-swagger.WrapHandler.func1 (3 handlers)
[GIN-debug] GET    /api/v1/tags              --> github.com/luanluanxu/go-gin-app/routers/api/v1.GetTags (4 handlers)
[GIN-debug] POST   /api/v1/tags              --> github.com/luanluanxu/go-gin-app/routers/api/v1.AddTag (4 handlers)
[GIN-debug] PUT    /api/v1/tags/:id          --> github.com/luanluanxu/go-gin-app/routers/api/v1.EditTag (4 handlers)
[GIN-debug] DELETE /api/v1/tags/:id          --> github.com/luanluanxu/go-gin-app/routers/api/v1.DeleteTag (4 handlers)
[GIN-debug] GET    /api/v1/articles          --> github.com/luanluanxu/go-gin-app/routers/api/v1.GetArticles (4 handlers)
[GIN-debug] GET    /api/v1/articles/:id      --> github.com/luanluanxu/go-gin-app/routers/api/v1.GetArticle (4 handlers)
[GIN-debug] POST   /api/v1/articles          --> github.com/luanluanxu/go-gin-app/routers/api/v1.AddArticle (4 handlers)
[GIN-debug] PUT    /api/v1/articles/:id      --> github.com/luanluanxu/go-gin-app/routers/api/v1.EditArticle (4 handlers)
[GIN-debug] DELETE /api/v1/articles/:id      --> github.com/luanluanxu/go-gin-app/routers/api/v1.DeleteArticle (4 handlers)

Listening port is 8000
Actual pid is 4393
```
Swagger doc

![image](https://i.imgur.com/bVRLTP4.jpg)

## Features

- RESTful API
- Gorm
- Swagger
- logging
- Jwt-go
- Gin
- Graceful restart or stop (fvbock/endless)
- App configurable
- Cron
- Redis