+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "NuGet"
description = "Publish files and artifacts to NuGet repository"
user = "drone-plugins"
repo = "drone-nuget"
image = "plugins/drone-nuget"
tags = ["nuget"]
categories = "publish"
draft = false
date = 2016-02-13T09:01:59Z
menu = ""
weight = 1

+++

Use this plugin to publish artifacts from the build to a NuGet Repository
You can override the default configuration with the following parameters:

* `source` - NuGet Repository URL
* `api_key` - Api Key used for authentication
* `verbosity` - NuGet output verbosity, default to 'quiet'. Accept: normal, quiet or detailed
* `files` - List of files to upload

All file paths must be relative to current project sources

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  nuget:
    source: http://nuget.company.com
    api_key: <Your Key>
    files: 
      - *.nupkg
```

## .nuspec / .nupkg

If a file to upload does have ```.nuspec``` extension, the __nuget pack__ command is called before a push

If a file to upload does have ```.nupkg``` extension, it is pushed directly

