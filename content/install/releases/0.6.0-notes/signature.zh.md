+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 移除签名"
url = "zh/release-0.6.0-signature"
+++

<!--The yaml signature file is now removed. This means you no longer have to sign your yaml every time it changes in order to expose secrets. This does have some security implications for open source projects. Please read below.-->

yaml 签名已经被移除，这意味着不再需要签名 yaml 文件来使用密文。这对开源项目有一定的安全隐患，请阅读下面的章节来了解更多信息。

<!--# Pull Requests-->

# Pull Requests

<!--Secrets are not exposed to pull requests by default. If your repository is public and you choose to override this behavior please note that your secrets are not safe. It is possible for a bad actor to submit a pull request and expose your secrets.-->

默认在 pull request 中不使用密文。如果您的项目是开源项目，您希望修改这一默认行为，请注意这是不安全的。恶意的 pull request 可以暴露密文。

<!--# Pull Requests + Gating-->

# Pull Requests + Gating 门禁

<!--Drone has experimental support for gating pull requests and blocking builds from unknown accounts. This is perhaps the best way to mitigate security issues when exposing secrets to pull requests.-->

Drone 实验性支持 门禁 gating pull request 以及拦截来自未知账户的 pull request。这可能是处理密文被暴露等安全问题的最好方式。
