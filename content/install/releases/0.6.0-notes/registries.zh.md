+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Registry 凭证"
url = "zh/release-0.6.0-registry-credentials"
+++

<!--Registry credentials for downloading private pipeline images are now managed separately from secrets, and are no longer specified in the yaml configuration file.-->

用来下载工作流镜像的 Registry 凭证已经和密文分开，不再在 yaml 文件里面配置。

<!--Example legacy configuration:-->

```diff
pipeline:
  build:
    image: bradrydzewski/private-image
-   auth_config:
-     username: bradrydzewski
-     password: password
    commands:
      - go build
      - go test
```

<!--Registry credentials are adding to drone using the command line utility:-->

使用 Drone 命令行工具来添加 Registry 凭证

```text
drone registry add \
  --repository=<repo> \
  --hostname=docker.io \
  --username=<username> \
  --password=<password>
```

<!--Example command:-->

使用示例：

```text
drone registry add \
  --repository=octocat/hello-world \
  --hostname=docker.io \
  --username=bradrydzewski \
  --password=password
```

<!--Drone matches the registry hostname to each image in your yaml. If the hostnames match the registry credentials are used to authenticate to your registry and pull the image. Note that registry credentials are used by the Drone agent and are never exposed to your build containers.-->

Drone 使用 hostname 来匹配 yaml 文件中的镜像。如果 hostname 匹配，对应凭据将被用来认证和拉取对应镜像文件。 Registry 凭证只被 Drone 代理客户端（Agent）使用，不会暴露给构建容器。

<!--Example registry hostnames:-->

示例 registry hostname：

* `gcr.io/foo/bar` hostname `gcr.io`
* `foo/bar` hostname `docker.io`
* `qux.com:8000/foo/bar` hostname `qux.com:8000`

<!--Example registry hostname matching logic:-->

示例 registry hostname 匹配逻辑：

* Hostname `gcr.io` 匹配 `gcr.io/foo/bar`
* Hostname `docker.io` 匹配 `golang`
* Hostname `docker.io` 匹配 `library/golang`
* Hostname `docker.io` 匹配 `bradyrydzewski/golang`
* Hostname `docker.io` 匹配 `bradyrydzewski/golang:latest`