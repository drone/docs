---
title: Selenium 示例项目
url: zh/selenium-example

menu:
  usage:
    weight: 6
    identifier: selenium_example-zh
    parent: usage_examples
---

Example Yaml configuration for a project with a Selenium service dependency. Note the selenium service will be available to at `http://selenium:4444/wd/hub`

```diff
pipeline:
  build:
    image: golang
    commands:
      - go get
      - go test

services:
+ selenium:
+   image: selenium/standalone-chrome
```

Example Yaml configuration for a project with a Selenium Grid service dependency, using Chrome and Firefox Selenium nodes.

```yaml
services:
  selenium:
    image: selenium/hub

  chrome:
    image: selenium/node-chrome
    environment:
      HUB_PORT_4444_TCP_ADDR: selenium
      HUB_PORT_4444_TCP_PORT: 4444
      DISPLAY: ":99.0"

  firefox:
    image: selenium/node-firefox
    environment:
      HUB_PORT_4444_TCP_ADDR: selenium
      HUB_PORT_4444_TCP_PORT: 4444
      DISPLAY: ":98.0"
```
