---
date: 2000-01-01T00:00:00+00:00
title: Anka
author: eoinmcafee
weight: 100
---

<div class="alert">
Anka is in the Beta phase.
</div>

The goal of this document is to give you enough technical specifics to configure and run osx vms.
When properly configured, it will automatically provision and terminate osx vms based on your Drone server’s build volume.

# Prerequisites

Pre-build image are available in the AWS marketplace from Anka, just search of “anka” or using the following step's, use your own mac instance:

Recommend Specs:
* Volume size: 500GB
* Volume type: gp3
* Iops: 6000

If you find that the disk has been split over multiple volumes then do the following:

https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-mac-instances.html#mac-instance-increase-volume

Anka setup : Full documentation can be found at: https://docs.veertu.com/anka/intel/

Download Anka:
{{< highlight bash "linenos=table" >}}
FULL_FILE_NAME="$(curl -Ls -r 0-1 -o /dev/null -w %{url_effective} https://veertu.com/downloads/anka-virtualization-latest | cut -d/ -f5)"
curl -S -L -o ./$FULL_FILE_NAME https://veertu.com/downloads/anka-virtualization-latest
sudo installer -pkg $FULL_FILE_NAME -tgt /
{{< / highlight >}}

Download ISO:
{{< highlight bash "linenos=table" >}}
softwareupdate --list-full-installers
{{< / highlight >}}

Select version:
{{< highlight bash "linenos=table" >}}
softwareupdate --fetch-full-installer --full-installer-version 12.2.1
{{< / highlight >}}

Activate license
{{< highlight bash "linenos=table" >}}
sudo anka license activate {LICENSE}
{{< / highlight >}}

Create the VM:
{{< highlight bash "linenos=table" >}}
anka --debug  create {{vm-name}} --ram-size 8G --cpu-count 4 --disk-size 50G --app /Applications/Install\ macOS\ Monterey.app
{{< / highlight >}}

Run this to get the ID of the VM which has recently been created.
{{< highlight bash "linenos=table" >}}
anka list // lists vms
{{< / highlight >}}

Install the following on the VM:
<div class="alert">
Docker is not support on M1 Architectures
</div>
{{< highlight bash "linenos=table" >}}
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

brew install wget

brew install docker —-cask
{{< / highlight >}}

Please note you can use the run command to execute commands inside of VM from the host.
 - https://docs.veertu.com/anka/arm/command-line-reference/#run 

If using a paid license version you can then suspend the VM & all the vm's will startup from a suspended state - providing a faster provisioning time

# Installation

Download Binary
{{< highlight bash "linenos=table" >}}
curl https://github.com/drone-runners/drone-runner-aws/releases/{{version}}
{{< / highlight >}}

Install [Tmux](https://github.com/tmux/tmux/wiki)
{{< highlight bash "linenos=table" >}}
brew install tmux
{{< / highlight >}}

Create your env file:

The VM runner can be configured using environment variables. This article references the below configuration options. See [Reference]({{< relref "../../configuration/reference" >}}) for a complete list of configuration options.

- __DRONE_RPC_HOST__
  : provides the hostname (and optional port) of your Drone server. The runner connects to the server at the host address to receive pipelines for execution.
- __DRONE_RPC_PROTO__
  : provides the protocol used to connect to your Drone server. The value must be either http or https.
- __DRONE_RPC_SECRET__
  : provides the shared secret used to authenticate with your Drone server. This must match the secret defined in your Drone server configuration.
- __DRONE_MIN_POOL_SIZE__ **(optional)**
  : provides the minimum size of the pool. The pool will not shrink below this size. The default is 1.
- __DRONE_MAX_POOL_SIZE__ **(optional)**
  : provides the maximum size of the pool. The pool will not grow above this size. The default is 2.

Create pool file:

Example pool file
{{< highlight yaml "linenos=table" >}}
version: "1"
instances:
- name: osx-anka
  default: true
  type: anka   # type of instance.  amazon | gcp | vmfusion | anka
  pool: 2    # total number of warm instances in the pool at all times
  limit: 100  # limit the total number of running servers. If exceeded block or error.
  platform:
  os: darwin
  arch: amd64
  spec:
  account:
  username: admin
  password: admin
  vm_id: vm-id
  {{< / highlight >}}

To start runner and write to external log:
{{< highlight bash "linenos=table" >}}
tmux new -d './drone-runner-aws daemon --envfile=.env --pool=pool.yml |& tee runner.log'
{{< / highlight >}}
