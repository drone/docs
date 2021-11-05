[![Build Status](https://cloud.drone.io/api/badges/drone/docs/status.svg)](https://cloud.drone.io/drone/docs)

This repository contains the source for [docs.drone.io](http://docs.drone.io).
To generate the documentation you will need to download and install the [hugo](https://gohugo.io/overview/installing/) static website engine.
If you are following the [Windows installation instructions](https://gohugo.io/getting-started/installing/#chocolatey-windows) you will need the extended version.

Generate the documentation:

```
$ hugo

                   | EN   
-------------------+------
  Pages            | 869  
  Paginator pages  |   0  
  Non-page files   |   1  
  Static files     |  43  
  Processed images |   0  
  Aliases          | 355  
  Sitemaps         |   1  
  Cleaned          |   0  

Total in 926 ms
```

Generate the documentation and serve at `localhost:1313`:

```
$ hugo server -b http://localhost:1313 -w

                   | EN   
-------------------+------
  Pages            | 869  
  Paginator pages  |   0  
  Non-page files   |   1  
  Static files     |  43  
  Processed images |   0  
  Aliases          | 355  
  Sitemaps         |   1  
  Cleaned          |   0  

Built in 379 ms
Web Server is available at http://localhost:1313/ (bind address 127.0.0.1)
Press Ctrl+C to stop
```

**For Windows Users**

We have found an issue with `highlight yaml` on Hugo. 
<br>
You will not be able to see any yaml syntax highlighting when previewing the docs at http://localhost:1313/ on a Windows machine.
<br>
If run into this issue, you can try running Hugo on [WSL](https://docs.microsoft.com/en-us/windows/wsl/install) to preview the page.