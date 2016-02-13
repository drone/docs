+++

# THIS DOCUMENT IS AUTOMATICALLY GENERATED. PLEASE DO NOT EDIT

title = "Azure Storage"
description = "Publish files and artifacts to Azure Storage"
user = "drone-plugins"
repo = "drone-azure-storage"
image = "plugins/drone-azure-storage"
tags = ["azure", "storage"]
categories = "publish"
draft = false
date = 2016-02-13T08:58:43Z
menu = ""
weight = 1

+++

Use this plugin for publishing files and artifacts to Azure Storage. The plugin is power by [blobxfer](https://github.com/Azure/azure-batch-samples/tree/master/Python/Storage) - an upload tool from the Microsoft HFC team. The following parameters are required:

* `account_key` - Storage Account Key, required for authentication
* `storage_account` - Storage Account name
* `container` - The target storage container
* `source` - Location of folder to sync relative to the workspace root

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
publish:
  azure_storage:
    account_key: 123889asd89u8hsfdjh98128hh
    storage_account: my-storage-account
    container: my-storage-container
    source: folder/to/upload
```

