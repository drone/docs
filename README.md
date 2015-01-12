This is the [Drone](https://github.com/drone/drone) documentation found at [readme.drone.io](http://readme.drone.io)

## Install Hugo

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

## Generate Docs

Generate the website and output to `./public`

```
hugo -t drone
```

Generate the website, serve on localhost:1313, and automatically refresh the browser when a change is detected:

```
hugo server -w -t drone
```
