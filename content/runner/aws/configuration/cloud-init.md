---
date: 2000-01-01T00:00:00+00:00
title: Custom cloud-init
author: tphoney
weight: 2
toc: true
description: |
  Custom cloud-init or powershell scripts.
---

In some cases you may need to customize an instance at creation, before the build is started.

**Overriding the default cloud-init file is an advanced feature and should be avoided unless absolutely necessary.**

You can customize your instance configuration by providing a custom cloud-init file.

Templating is provided using golang template syntax, you can find more information [here](https://golangdocs.com/templates-in-golang).

There are 2 variables available in the spec section of your pool file:

# `user_data` field

The contents of the cloud-init file as a string in the `pool.yml`

{{< highlight yaml "linenos=table,hl_lines=3-9" >}}
version: "1"
instances:
  - name: test_pool
    default: true
    type: amazon
    pool: 1    
    limit: 1 
    platform:
      os: linux
      arch: amd64
    spec:
      account:
        region: us-east-2
        access_key_id: XXXXXXXXXXXXXXXXXXXXX
        access_key_secret: XXXXXXXXXXXXXXXXXXXXX
      ami: ami-051197ce9cbb023ea
      size: t2.micro
      user_data: |
        #cloud-config
        apt:
          sources:
            docker.list:
              source: deb [arch={{ .Architecture }}] https://download.docker.com/linux/ubuntu $RELEASE stable
              keyid: 9DC858229FC7DD38854AE2D88D81803C0EBFCD88
        packages:
        - wget
        - docker-ce
        write_files:
        - path: {{ .CaCertPath }}
          permissions: '0600'
          encoding: b64
          content: {{ .CACert | base64  }}
        - path: {{ .CertPath }}
          permissions: '0600'
          encoding: b64
          content: {{ .TLSCert | base64 }}
        - path: {{ .KeyPath }}
          permissions: '0600'
          encoding: b64
          content: {{ .TLSKey | base64 }}
        runcmd:
        - 'wget "{{ .LiteEnginePath }}/lite-engine-{{ .Platform }}-{{ .Architecture }}" -O /usr/bin/lite-engine'
        - 'chmod 777 /usr/bin/lite-engine'
        - 'touch /root/.env'
        - 'touch /tmp/magic_lives_here'
        - '/usr/bin/lite-engine server --env-file /root/.env > /var/log/lite-engine.log 2>&1 &'
    {{< / highlight >}}

# `user_data_path` field

 The path to the cloud-init file as a string in the `pool.yml`

{{< highlight yaml "linenos=table" >}}
version: "1"
instances:
  - name: test_pool
    default: true
    type: amazon
    pool: 1    
    limit: 1 
    platform:
      os: linux
      arch: amd64
    spec:
      account:
        region: us-east-2
        access_key_id: XXXXXXXXXXXXXXXXXXXXX
        access_key_secret: XXXXXXXXXXXXXXXXXXXXX
      ami: ami-051197ce9cbb023ea
      size: t2.micro
      user_data_path: /tmp/user-data.yml
    {{< / highlight >}}


You can see an example of a custom cloud-init template file [here](https://github.com/drone-runners/drone-runner-aws/blob/master/template/cloud-init).
