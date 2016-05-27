+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Samples"
weight = 31
toc = true


[menu.main]
	parent="usage"
+++

# Overview

This section of the documentation provides sample Yaml files for common configurations and frequently asked questions. This is not meant to be an exhaustive list. If you would like to share your custom configurations please post to our [discourse forum](https://discuss.drone.io/c/how-tos).

# Mysql database

Example configuration that starts a mysql container. The mysql container is available to the golang container at 127.0.0.1:3306. Please note that environment variables are used to create the default database, username and password.

```
pipeline:
  test:
    image: golang
    commands:
      - sleep 10 # gives the database enough time to initialize
      - go get
      - go test

services:
  database:
    image: mysql:5.6.27
    environment:
      - MYSQL_DATABASE=test
      - MYSQL_ALLOW_EMPTY_PASSWORD=yes  
```

# Postgres database

Example configuration that starts a postgres container. The postgres container is available to the golang container at 127.0.0.1:5432. Please note that environment variables are used to create the default database, username and password.

```
pipeline:
  test:
    image: golang
    commands:
      - sleep 10 # gives the database enough time to initialize
      - go get
      - go test

services:
  database:
    image: postgres:9.4.5
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_DB=test
```

# Publish Docker Images

Example configuration that uses the `docker` plugin to build and publish an image to DockerHub. Please note that you will need to provide the plugin with your DockerHub credentials.

```
pipeline:
  test:
    image: golang
    commands:
      - go test
      - go build

  docker:
    repo: octocat/hello-world
    tag: latest
```

# Publish GitHub Pages

Example configuration that uses the `gh_pages` plugin to publish your website. You can use this plugin in conjunction with static site generators like jekyll or hugo. Please note that you will need to provide the plugin with your repository credentials.

```
pipeline:
  build:
    image: publysher/hugo
    commands:
      - hugo

  gh_pages:
    source: public/
```

# Selenium

Example configuration that starts a selenium container. The selenium container is available to the golang container at 127.0.0.1:4444, with the executor available at http://127.0.0.1:4444/wd/hub

```
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go run main.go &  # start program in background
      - go test           # test program with selenium

services:
  selenium:
    image: selenium/standalone-chrome
```

# Selenium Grid

Example configuration that starts selenium grid with chrome and firefox nodes. The selenium grid container is available to the golang container at 127.0.0.1:4444, with the executor available at http://127.0.0.1:4444/wd/hub

```
services:
  selenium:
    image: selenium/hub

  chrome:
    image: selenium/node-chrome
    environment:
      SE_OPTS: "-port 4445"
      HUB_PORT_4444_TCP_ADDR: 127.0.0.1
      HUB_PORT_4444_TCP_PORT: 4444

  firefox:
    image: selenium/node-firefox
    environment:
      SE_OPTS: "-port 4446"
      HUB_PORT_4444_TCP_ADDR: 127.0.0.1
      HUB_PORT_4444_TCP_PORT: 4444
      DISPLAY: ":98.0"
```
