---
date: 2000-01-01T00:00:00+00:00
title: Setup Command
author: tphoney
weight: 2
toc: true
description: |
  Test your vm/cloud configuration.
---

This command will help check you get up and running with the runner, it will verify that your credentials can create an instance and that we can run a build on it. It will output an example pool file `pool.yml` based on what you VM credentials you pass it, if you want to create a custom pool file. (It performs more checks than the daemon command)

You can locate the binaries here (if you have not already done so): [releases](https://github.com/drone-runners/drone-runner-aws/releases)

# Amazon

To use the setup command for amazon you will need to pass through your key and secret. Below is an example of how to use the setup command.

{{< highlight bash "linenos=table" >}}
./drone-runner-aws setup --aws-access-key-id"="your key" --aws-access-key-secret="your secret"
INFO[0000] setup: using amazon
INFO[0000] no pool file provided, creating in memory pool for amazon
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

# Anka

To use the setup command for Anka you will need to pass through your vm name. Below is an example of how to use the setup command.

{{< highlight bash "linenos=table" >}}
./drone-runner-aws setup --anka-vm-name master-vm
INFO[0000] setup: using anka
INFO[0000] no pool file provided
DEBU[0003] Not there yet 1/60, error: exit status 2
DEBU[0007] Not there yet 2/60, error: exit status 2
DEBU[0011] Not there yet 3/60, error: exit status 2
DEBU[0013] Not there yet 4/60, error: exit status 2q
DEBU[0016] got IP 192.168.64.32                          cloud=anka name=setup--5577006791947779410 pool=testpool
INFO[0017] Running script in VM                          cloud=anka name=setup--5577006791947779410 pool=testpool
DEBU[0273] anka: [creation] complete                     cloud=anka fields.time=273.27s ip=192.168.64.32 name=setup--5577006791947779410 pool=testpool
INFO[0273] setup: instance logs for da894645-10c1-460f-886f-df6d65677c52:
Pool file:
version: "1"
instances:
- name: testpool
  default: true
  type: anka
  pool: 1
  limit: 2
  platform:
    os: darwin
    arch: amd64
    spec:
    account:
      username: ""
      password: ""
    vm_id: master-vm
    root_directory: ""
    user_data: ""
{{< / highlight >}}

For more information about the Anka configuration options, see [Anka]({{< relref "anka" >}})

# Digital Ocean

To use the setup command for digital ocean you will need to pass through your project id. Below is an example of how to use the setup command.

{{< highlight bash "linenos=table" >}}
./drone-runner-aws setup --digital-ocean-pat XXXXXXXXX
INFO[0000] setup: using digital ocean
INFO[0000] no pool file provided
INFO[0000] in memory pool is using digitalocean
INFO[0000] digitalocean: creating instance setup-testpool-1655379844  driver=digitalocean hibernate=false image=docker-18-04 pool=testpool
INFO[0001] digitalocean: instance created setup-testpool-1655379844  driver=digitalocean hibernate=false image=docker-18-04 pool=testpool
INFO[0003] digitalocean: firewall configured setup-testpool-1655379844  driver=digitalocean hibernate=false image=docker-18-04 pool=testpool
DEBU[0003] find instance network                         driver=digitalocean hibernate=false image=docker-18-04 name=setup-testpool-1655379844 pool=testpool
DEBU[0063] find instance network                         driver=digitalocean hibernate=false image=docker-18-04 name=setup-testpool-1655379844 pool=testpool
INFO[0064] setup: instance logs for 304473745: no logs here
TRAC[0064] setup: running healthcheck and waiting for an ok response
TRAC[0129] RetryHealth: health check completed           duration=1m5.3480743s
TRAC[0129] LE.RetryHealth check complete                 response="&{Version:0.1.0 DockerInstalled:true GitInstalled:true LiteEngineLog:time=\"2022-06-16T11:45:44Z\" level=info msg=\"server listening at port :9079\"\ntime=\"2022-06-16T11:45:44Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-06-16T11:45:44Z\" level=info msg=\"git is installed\"\ntime=\"2022-06-16T11:45:44Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES\ntime=\"2022-06-16T11:45:44Z\" level=info msg=\"docker is installed\"\ntime=\"2022-06-16T11:46:13Z\" level=info msg=\"handler: HandleHealth()\"\ntime=\"2022-06-16T11:46:13Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES\ntime=\"2022-06-16T11:46:13Z\" level=info msg=\"docker is installed\"\ntime=\"2022-06-16T11:46:13Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-06-16T11:46:13Z\" level=info msg=\"git is installed\"\n OK:true}"
Pool file:
version: "1"
instances:
- name: testpool
  default: true
  type: digitalocean
  pool: 1
  limit: 2
  platform:
    os: linux
    arch: amd64
  spec:
    account:
      pat: XXXXXXXXX

DEBU[0130] deleting droplet                              driver=amazon id="[304473745]"
DEBU[0130] droplet deleted                               driver=amazon id="[304473745]"
TRAC[0130] digitalocean: VM terminated                   driver=amazon id="[304473745]"
{{< / highlight >}}

For more information about the Amazon configuration options, see [Digital Ocean]({{< relref "../drivers/digitalocean/_index.md" >}})

# Azure

To use the setup command for azure you will need to pass through ...

# Google

To use the setup command for Google you will need to pass through your project id. Below is an example of how to use the setup command.

{{< highlight bash "linenos=table" >}}
./drone-runner-aws setup --google-project-id myProjectId
INFO[0000] setup: using google
INFO[0000] no pool file provided
INFO[0000] in memory pool is using google
DEBU[0000] finding default firewall rules
DEBU[0000] found default firewall rule
TRAC[0000] google: creating VM                           cloud=google image=projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131 name=setup-testpool-1655380357 pool=testpool size=e2-small zone=northamerica-northeast1-a
DEBU[0010] instance insert operation completed           cloud=google image=projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131 name=setup-testpool-1655380357 pool=testpool size=e2-small zone=northamerica-northeast1-a
DEBU[0010] google: [provision] VM provisioned            cloud=google fields.time=9.85s image=projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131 ip=2360638475713554281 name=setup-testpool-1655380357 pool=testpool size=e2-small zone=northamerica-northeast1-a
DEBU[0010] google: [provision] complete                  cloud=google fields.time=10.09s image=projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131 ip=34.152.16.141 name=setup-testpool-1655380357 pool=testpool size=e2-small zone=northamerica-northeast1-a
ERRO[0010] setup: unable to get instance logs            error=Unimplemented
INFO[0010] setup: instance logs for 4546377208440465258:  
TRAC[0010] setup: running healthcheck and waiting for an ok response
TRAC[0043] health check failed. Retrying                 error="Get \"https://34.152.16.141:9079/healthz\": dial tcp 34.152.16.141:9079: connect: connection refused" retry_num=0
TRAC[0143] RetryHealth: health check completed           duration=2m12.159454s
TRAC[0143] LE.RetryHealth check complete                 response="&{Version:0.1.0 DockerInstalled:true GitInstalled:true LiteEngineLog:time=\"2022-06-16T11:54:58Z\" level=info msg=\"server listening at port :9079\"\ntime=\"2022-06-16T11:54:58Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-06-16T11:54:58Z\" level=info msg=\"git is installed\"\ntime=\"2022-06-16T11:54:58Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES\ntime=\"2022-06-16T11:54:58Z\" level=info msg=\"docker is installed\"\ntime=\"2022-06-16T11:54:59Z\" level=info msg=\"handler: HandleHealth()\"\ntime=\"2022-06-16T11:54:59Z\" level=info msg=\"checking docker is installed\"\nCONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES\ntime=\"2022-06-16T11:54:59Z\" level=info msg=\"docker is installed\"\ntime=\"2022-06-16T11:54:59Z\" level=info msg=\"checking git is installed\"\ntime=\"2022-06-16T11:54:59Z\" level=info msg=\"git is installed\"\n OK:true}"
Pool file:
version: "1"
instances:
- name: testpool
  default: true
  type: google
  pool: 1
  limit: 2
  platform:
    os: linux
    arch: amd64
  spec:
    account:
      project_id: drone-ci-289110
      json_path: ~/.config/gcloud/application_default_credentials.json
    image: projects/ubuntu-os-pro-cloud/global/images/ubuntu-pro-1804-bionic-v20220131
    machine_type: e2-small
    zone:
    - northamerica-northeast1-a
{{< / highlight >}}

For more information about the Amazon configuration options, see [Google]({{< relref "google" >}})
