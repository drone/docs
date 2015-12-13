[![Build Status](http://test.drone.io/api/badge/github.com/drone/docs/status.svg?style=flat)](http://test.drone.io/github.com/drone/docs) [![Gitter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/drone/drone?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

Documentation for the [Drone](https://github.com/drone/drone) continuous integration server, found at [readme.drone.io](http://readme.drone.io)

## Setup

This website uses the [Hugo](https://github.com/spf13/hugo) static site generator. If you are planning to contribute you'll want to download and install Hugo on your local machine.

Install on OSX with brew:

```
brew install hugo
```

Install using the Go tool:

```
go get github.com/spf13/hugo
```

Or follow these instructions for your platform: http://gohugo.io/overview/installing/

## Build

Generate the website and output to `./public`

```
hugo
```

Generate the website, serve on localhost:1313, and automatically refresh the browser when a change is detected:

```
hugo server -w
```
