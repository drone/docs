---
date: 2016-01-01T00:00:00+00:00
title: "Bash 插件示例"
author: bradrydzewski
url: zh/creating-custom-plugins-bash
weight: 4

menu:
  plugin:
    parent: pipeline
    identifier: creating-custom-plugins-bash-zh
---

<!--This provides a brief tutorial for creating a Drone webhook plugin, using simple shell scripting, to make an http requests during the build pipeline. The below example demonstrates how we might configure a webhook plugin in the Yaml file:-->

这个简单的教程使用 shell 脚本来生成一个 Drone webhook 插件，在构建工作流中生成 http 请求。下面的例子展示了我们在 yaml 文件中如何配置 webhook 插件：

```yaml
pipeline:
  webhook:
    image: foo/webhook
    url: http://foo.com
    method: post
    body: |
      hello world
```

<!--Create a simple shell script that invokes curl using the Yaml configuration parameters, which are passed to the script as environment variables in uppercase and prefixed with `PLUGIN_`.-->

使用 yaml 参数来使用简单的 shell 脚本来启动 curl 命令，这些参数是使用前置字符串全大写 `PLUGIN_` 的环境变量：

```bash
#!/bin/sh

curl \
  -X ${PLUGIN_METHOD} \
  -d ${PLUGIN_BODY} \
  ${PLUGIN_URL}

```

<!--Create a Dockerfile that adds your shell script to the image, and configures the image to execute your shell script as the main entrypoint.-->

新建一个 Dockerfile ，并将 shell 脚本添加到这个镜像中，并将 shell 脚本作为镜像运行的入口。

```dockerfile
FROM alpine
ADD script.sh /bin/
RUN chmod +x /bin/script.sh
RUN apk -Uuv add curl ca-certificates
ENTRYPOINT /bin/script.sh
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
