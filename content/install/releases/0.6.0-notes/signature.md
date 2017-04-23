+++
date = "2017-04-15T14:39:04+02:00"
title = "0.6.0 Signature Removed"
url = "release-0.6.0-signature"
+++

The yaml signature file is now removed. This means you no longer have to sign your yaml every time it changes in order to expose secrets. This does have some security implications for open source projects. Please read below.


# Pull Requests

Secrets are not exposed to pull requests by default. If your repository is public and you choose to override this behavior please note that your secrets are not safe. It is possible for a bad actor to submit a pull request and expose your secrets.

# Pull Requests + Gating

Drone has experimental support for gating pull requests and blocking builds from unknown accounts. This is perhaps the best way to mitigate security issues when exposing secrets to pull requests.
