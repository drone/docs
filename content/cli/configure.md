---
date: 2000-01-01T00:00:00+00:00
title: Configuration
author: bradrydzewski
weight: 2
aliases:
- /cli-authentication/
- /cli/setup
---

The command line tools interact with the server using REST endpoints. You will need to provide the CLI tools with the server addresses and your personal authorization token. You can find your authorization token in your Drone account settings (click your Avatar in the user interface).

1. Configure your Drone server address:
   ```
   export DRONE_SERVER=http://drone.mycompany.com
   ```

2. Configure your Drone personal token:
   ```
   export DRONE_TOKEN=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   ```