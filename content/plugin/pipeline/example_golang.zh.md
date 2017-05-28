---
date: 2016-01-01T00:00:00+00:00
title: "Go 插件示例"
author: bradrydzewski
url: zh/creating-custom-plugins-golang
weight: 5

menu:
  plugin:
    parent: pipeline
    identifier: creating-custom-plugins-golang
---

<!--This guide provides a brief tutorial for creating a webhook plugin, using the Go programming language, to trigger http requests during the build pipeline. The below example demonstrates how we might configure a webhook plugin in the Yaml file:-->

这个简单的教程使用 Go 语言来生成一个 Drone webhook 插件，在构建工作流中生成 http 请求。这个例子展示了如何使用 yaml 文件配置 webhook 插件：

```yaml
pipeline:
  webhook:
    image: foo/webhook
    url: http://foo.com
    method: post
    body: hello world
```

<!--Create a Go program that makes an http request using the Yaml configuration parameters, which are passed to the program as environment variables in uppercase, prefixed with `PLUGIN_`.-->

新建一个 Go 程序来根据 yaml 配置参数来生成一个 http 请求，这些参数是使用前置字符串全大写 `PLUGIN_` 的环境变量：

```Go
package main

import (
  "net/http"
  "os"
)

func main() {
  body := strings.NewReader(
    os.GetEnv("PLUGIN_BODY"),
  )

  req, err := http.NewRequest(
    os.GetEnv("PLUGIN_METHOD"),
    os.GetEnv("PLUGIN_URL"),
    body,
  )
  if err != nil {
    os.Exit(1)
  }
}
```

<!--Compile your binary on the host machine for the target platform. Compiling on the host machine and adding the binary to the image is considered a best practice because it reduces the overall image size.-->

在宿主机器上编译程序为对应平台的二进制文件，编译程序后将对应二进制文件添加到镜像文件中可以减小 Docker 镜像体积，是最佳实践之一。

<!--It is very important to compile using the correct target platform, otherwise your plugin will fail with a cryptic Go runtime error.-->

{{% alert warn %}}
选择正确的编译平台十分重要，否则插件可能会遇到 Go 运行时错误。
{{% /alert %}}

```nohighlight
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o webhook
```

<!--Create a Dockerfile that adds your compiled binary to the image, and configures the image to run your binary as the main entrypoint.-->

新建一个 Dockerfile 来将编译后的二进制文件到镜像中，并把这个程序作为镜像运行的入口。

```dockerfile
FROM alpine
ADD webhook /bin/
RUN apk -Uuv add ca-certificates
ENTRYPOINT /bin/webhook
```

<!--Build and publish your plugin to the Docker registry. Once published your plugin can be shared with the broader Drone community.-->

构建和发布镜像到 Docker registry。镜像发布后你的插件就可以被 Drone 社区的用户使用。

```nohighlight
docker build -t foo/webhook .
docker push foo/webhook
```

<!--Execute your plugin locally from the command line to verify it is working:-->

在本地执行你的插件来确定你的插件的运行情况：

```nohighlight
docker run --rm \
  -e PLUGIN_METHOD=post \
  -e PLUGIN_URL=http://foo.com \
  -e PLUGIN_BODY="hello world" \
  foo/webhook
```
