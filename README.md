The Webserv Project
=======================

WebServ is an simple http file server. It support both upload & download & manage operation.
The Project has no third-party dependence. It would works on any platfrom if golang is support.
Support image and video transfer on iPhone, iPad, Android device.
The server implement basic auth base on base64, which could fulfill most environment.

## Usage
![idx shot](doc/shot.gif "The Webserv Project")

## Install

### Compile from code
``` Shell
- go get github.com/aoaolion/webserv
- go build
```
### Download binary
- Mac
> NOT available
- Linux
> NOT available
- Windows
> NOT available

## Run

``` Shell
./weberv -d [file root] -h [listen ip] -p [listen port] -t [ttl] -u [username] -P [password]
```

The service listen to 0.0.0.0:8080 as default. And the default dir is file_root.
The default ttl is set 600, which means 10mins. If you do not want ttl, take to value to 0.

Notice: Basic auth works when uername is not empty!

## Close

There are three ways to close this service.
- access api, http://ip:port/close 
- signals, such as kill, ctl+c
- ttl way can close the service, only if ttl value is not 0
