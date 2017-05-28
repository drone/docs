+++
date = "2017-04-15T14:39:04+02:00"
title = "Database Settings"
url = "zh/database-settings"

[menu.install]
  identifier = "database-settings-zh"
  parent = "install_server"
  weight = 1
+++

<!--This guide provides instructions for using alternate storage engines. Please note this is optional. The default storage engine is an embedded SQLite database which requires zero installation or configuration.-->

这个指南提供了使用其它储存引擎的步骤。请注意这个是可选的。Drone 默认使用嵌入式的 SQLite 数据库，不需要配置。

<!--# Configure MySQL-->

# 配置 MySQL

<!--The below example demonstrates mysql database configuration. See the official driver [documentation](https://github.com/go-sql-driver/mysql#dsn-data-source-name) for configuration options and examples.-->

下面的例子展示了配置 mysql 数据库。参考官方[驱动文档](https://github.com/go-sql-driver/mysql#dsn-data-source-name)来了解配置选项和例子。

```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     DRONE_DATABASE_DRIVER: mysql
+     DRONE_DATABASE_DATASOURCE: root:password@tcp(1.2.3.4:3306)/drone?parseTime=true
```

<!--# Configure Postgres-->
# 配置 Postgres

<!--The below example demonstrates postgres database configuration. See the official driver [documentation](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING) for configuration options and examples.-->

下面的例子展示了配置 postgres 数据库。参考官方[驱动文档](https://www.postgresql.org/docs/current/static/libpq-connect.html#LIBPQ-CONNSTRING)来了解配置选项和例子。


```diff
version: '2'

services:
  drone-server:
    image: drone/drone:{{% version %}}
    environment:
+     DRONE_DATABASE_DRIVER: postgres
+     DRONE_DATABASE_DATASOURCE: postgres://root:password@1.2.3.4:5432/postgres?sslmode=disable
```

<!--# Database Creation-->

# 数据库生成

<!--Drone does not create your database automatically. If you are using the mysql or postgres driver you will need to manually create your database using `CREATE DATABASE`-->

Drone 不自动新建数据库。您需要手动使用 `CREATE DATABASE` 来新建 mysql 和 postgres 数据库。

<!--# Database Migration-->

# 数据库迁移

<!--Drone automatically handles database migration, including the initial creation of tables and indexes. New versions of Drone will automatically upgrade the database unless otherwise specified in the release notes.-->

Drone 自动处理数据库迁移，包括初始的表和索引新建。新版本的 Drone 将会自动升级数据库除非在发布说密个闹钟功能说明。

<!--# Database Backups-->

# 数据库备份

<!--Drone does not perform database backups. This should be handled by separate third party tools provided by your database vendor of choice.-->

Drone 不进行数据库备份。这需要使用数据库厂商选择的第三方工具来完成。

<!--# Database Archiving-->

# 数据库归档

<!--Drone does not perform data archival; it considered out-of-scope for the project. Drone is rather conservative with the amount of data it stores, however, you should expect the database logs to grow the size of your database considerably.-->

Drone 不处理数据归档，这不属于这个项目的目标。 Drone 在它储存数据方面相对保守，但是数据库日志将占到数据库容量的很大一部分。
