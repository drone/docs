---
date: 2017-04-15T14:39:04+02:00
title: Services 服务
url: zh/services

xnext_steps:
  - file: mysql.zh.md
    name: MySQL 示例项目
  - file: postgres.zh.md
    name: Postgres 示例项目
  - file: redis.zh.md
    name: Redis 示例项目

menu:
  usage:
    weight: 6
    identifier: services-zh
    parent: usage_concepts
---

<!--Drone provides a services section in the Yaml file used for defining service containers. The below configuration composes database and cache containers.-->

Drone 在 Yaml 文件中使用 services 模块来定义服务容器。下面的配置文件组建了数据库服务容器和缓存服务容器。

```diff
pipeline:
  build:
    image: golang
    commands:
      - go build
      - go test

services:
  database:
    image: mysql

  cache:
    image: redis
```

<!--Services are accessed using custom hostnames. In the above example the mysql service is assigned the hostname `database` and is available at `database:3306`.-->

服务容器可以使用自定义的主机名称（hostnames）来访问。在上面的例子中， mysql 服务被指定了 `database` 的主机名称，可以使用 `database:3306` 来访问它。

# 配置

<!--Service containers generally expose environment variables to customize service startup such as default usernames, passwords and ports. Please see the official image documentation to learn more.-->

服务容器一般使用环境变量来自定义默认的用户名，密码以及端口。请查看官方的镜像文档来了解更多内容。

```diff
services:
  database:
    image: mysql
+   environment:
+     - MYSQL_DATABASE=test
+     - MYSQL_ALLOW_EMPTY_PASSWORD=yes

  cache:
    image: redis
```

# 初始化

<!--Service containers require time to initialize and begin accept connections. If you are unable to connect to a service you may need to wait a few seconds or implement a backoff.-->

服务容器需要一定时间来初始化，然后才能被访问。如果刚开始时无法连接到一个服务，可以先等几秒钟或者添加一个退让（backoff）的机制。

```diff
pipeline:
  test:
    image: golang
    commands:
+     - sleep 15
      - go get
      - go test

services:
  database:
    image: mysql
```
