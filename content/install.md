+++
date = "2015-12-05T16:00:21-08:00"
draft = false
title = "Self Hosting"
weight = 5
menu = "main"
toc = true
+++

> Please note the self-hosted version is completely free for small teams and
> non-profits. There may be a very modest annual licensing fee for large
> teams and enterprises in the future. Please contact
> [@bradrydzewski](mailto:brad@drone.io) if you have any questions or concerns.

# Requirements

LGTM ships as a single binary file inside a minimalist 20 MB Docker image.
Docker is the only dependency. If you are planning on integrating
with GitHub Enterprise it requires version 2.4 or higher.

# Installation

Get started by downloading the image from DockerHub:

```
sudo docker pull lgtm/lgtm
```

Create a `/etc/lgtm/lgtmrc` file to hold your configuration parameters in
`KEY=VALUE` format. Please note that these variables should not be quoted:

```
GITHUB_URL=https://github.mycompany.com
GITHUB_SCOPE=user:email,read:org,repo
GITHUB_CLIENT=
GITHUB_SECRET=
```

Create and run your container:

```
sudo docker run \
	--volume /var/lib/lgtm:/var/lib/lgtm \
	--env-file /etc/lgtm/lgtmrc \
	--restart=always \
	--publish=80:8000 \
	--detach=true \
	--name=lgtm \
	lgtm/lgtm
```

Note the above example mounts a volume on the host machine. This is important
because the default configuration uses a sqlite database and should therefore
be mounted on the host machine as a volume to avoid data loss.

```
--volume /var/lib/lgtm:/var/lib/lgtm
```

# Configuration

This is a full list of configuration options. Please note that many of these
options use default configuration value that should work for the majorify of
installations.

* `DEBUG=false` runs the server in debug mode with more verbose logs
* `CACHE_TTL=15m` the cache duration for certain github data
* `GITHUB_URL` github server url when using GitHub Enterprise
* `GITHUB_CLIENT` github oauth client id
* `GITHUB_SECRET` github oauth client secret
* `GITHUB_SCOPE` github oauth scopes
* `DATABASE_DRIVER=sqlite` the database driver
* `DATABASE_DATASOURCE` the database connection string

# Registration

Register your application with GitHub (or GitHub Enterprise) to create your client
id and secret. It is very import that the redirect URL matches your http(s) scheme
and hostname exactly with `/login` as the path.

Please use this screenshot for reference:

![github registration](/docs/images/app_registration.png)

# Reverse Proxies

If you are running behind a reverse proxy please ensure the `X-Forwarded-For`
and `X-Forwarded-Proto` variables are configured.

This is an example nginx configuration:

```nginx
location / {
    proxy_set_header X-Forwarded-For $remote_addr;
    proxy_set_header X-Forwarded-Proto $scheme;
    proxy_set_header Host $http_host;
	proxy_pass http://127.0.0.1:8000;
}
```

This is an example caddy server configuration:

```nginx
lgtm.mycomopany.com {
        proxy / localhost:8000 {
                proxy_header X-Forwarded-Proto {scheme}
                proxy_header X-Forwarded-For {host}
                proxy_header Host {host}
        }
}
```

Note that when running behind a reverse proxy you should change the recommended
port mappings from `--publish=80:8000` to something like `--publish=8000:8000`.
