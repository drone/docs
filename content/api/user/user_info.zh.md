+++
date = "2000-01-01T00:00:00+00:02"
title = "当前用户信息"
url = "zh/api-user-info"

[menu.api]
  weight = 1
  identifier = "api-user-info-zh"
  parent = "api_user"
+++

<!--Returns the currently authenticated user.-->

返回当前登录用户。

```text
GET /api/user
```

<!--Example Response Body:-->

响应 Body 示例：

```json
{
  "id": 1,
  "login": "octocat",
  "email": "octocat@github.com",
  "avatar_url": "http://www.gravatar.com/avatar/7194e8d48fa1d2b689f99443b767316c",
  "admin": false,
  "active": true
}
```
