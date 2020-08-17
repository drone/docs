---
date: 2000-01-01T00:00:00+00:00
title: DRONE_TLS_CERT
author: bradrydzewski
aliases:
- /installation/reference/drone-tls-cert/
---

Path to an SSL certificate used by the server to accept HTTPS requests. This configuration parameter is of type string and is optional.

Please note that the cert file should be the concatenation of the server’s certificate, any intermediates, and the certificate authority’s certificate.

```
DRONE_TLS_CERT=/path/to/cert.pem
```