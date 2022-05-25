---
date: 2000-01-01T00:00:00+00:00
title: Setup Command
author: tphoney
weight: 2
toc: true
description: |
  Test your vm/cloud configuration.
---

This command will help check you get upand running. It will constuct an example pool file `pool.yml` based on what you pass it. It performs more checks than the daemon command.

You can locate the binaries here (if you have not already done so): [releases](https://github.com/drone-runners/drone-runner-aws/releases)

Testing an amazon configuration:

{{< highlight bash "linenos=table" >}}
./drone-runner-aws setup --awsAccessKeyID="your key" --awsAccessKeySecret="your secret"
INFO[0000] setup: using amazon
INFO[0000] no pool file provided, creating in memmory pool for amazon
TRAC[0000] amazon: using default vpc, checking security groups  ami=ami-051197ce9cbb023ea hibernate=false image=ami-051197ce9cbb023ea pool=test_pool provider=amazon region=us-east-2 size=t2.micro
WARN[0000] aws: no security group specified assuming 'harness-runner'  ami=ami-051197ce9cbb023ea hibernate=false image=ami-051197ce9cbb023ea pool=test_pool provider=amazon region=us-east-2 size=t2.micro
TRAC[0001] amazon: provisioning VM                       ami=ami-051197ce9cbb023ea hibernate=false image=ami-051197ce9cbb023ea pool=test_pool provider=amazon region=us-east-2 size=t2.micro
DEBU[0002] amazon: [provision] created instance          ami=ami-051197ce9cbb023ea hibernate=false id=i-001caa6bd4a2c76a0 image=ami-051197ce9cbb023ea pool=test_pool provider=amazon region=us-east-2 size=t2.micro
TRAC[0003] amazon: [provision] checking instance IP address  ami=ami-051197ce9cbb023ea hibernate=false id=i-001caa6bd4a2c76a0 image=ami-051197ce9cbb023ea pool=test_pool provider=amazon region=us-east-2 size=t2.micro
DEBU[0003] amazon: [provision] complete                  ami=ami-051197ce9cbb023ea fields.time=2.92s hibernate=false id=i-001caa6bd4a2c76a0 image=ami-051197ce9cbb023ea ip=52.14.158.173 pool=test_pool provider=amazon region=us-east-2 size=t2.micro
INFO[0003] setup: instance logs for i-001caa6bd4a2c76a0: 'console output is empty'
TRAC[0003] setup: running healthcheck and waiting for an ok response
TRAC[0068] health check failed. Retrying                 error="Get \"https://52.14.158.173:9079/healthz\": dial tcp 52.14.158.173:9079: connect: connection refused" retry_num=0
TRAC[0100] RetryHealth: health check completed           duration=1m37.25737s
TRAC[0100] LE.RetryHealth check complete                 response="&{Version:0.1.0 DockerInstalled:true GitInstalled:true LiteEngineLog:time=\"2022-05-19T09:33:19Z\" level=info msg=\"server listening at port :9079\"\ntime=\"2022-05-19T09:33:19Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-05-19T09:33:19Z\" level=info msg=\"git is installed\"\ntime=\"2022-05-19T09:33:19Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"docker is installed\"\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"handler: HandleHealth()\"\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"docker is installed\"\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-05-19T09:33:20Z\" level=info msg=\"git is installed\"\n OK:true}"
Pool file:
version: "1"
instances:
- name: test_pool
  default: true
  type: amazon
  pool: 1
  limit: 2
  platform:
    os: linux
    arch: amd64
  spec:
    account:
      access_key_id: your key
      access_key_secret: your secret
      region: us-east-2
    size: t2.micro
    ami: ami-051197ce9cbb023ea
    hibernate: false
  {{< / highlight >}}

For more information about the Amazon configuration options, see [Amazon]({{< relref "amazon" >}})
