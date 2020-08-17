---
date: 2000-01-01T00:00:00+00:00
title: Blobs
author: bradrydzewski
weight: 20
aliases:
- /installation/storage/object/
description: |
  Configure object storage.
---


Drone stores large text files (e.g. build logs) in the database by default. Drone _optionally_ supports storing large files in S3, or any object store compatible with the S3 protocol. This can significantly reduce the size of your database, and may improve overall performance and scalability.

1. Provide the Drone server with your AWS credentials:
    ```
    AWS_ACCESS_KEY_ID=AC7CBED7971FDCD0
    AWS_SECRET_ACCESS_KEY=1fb023a72bd510053ab88ec3dd50a8eb5b8a1d58d9
    AWS_DEFAULT_REGION=us-east-1
    AWS_REGION=us-east-1
    ```

2. Provide the Drone server with your Bucket name:

    ```
    DRONE_S3_BUCKET=name-of-your-bucket
    ```
