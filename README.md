<h1 style="font-size: 40px">Forxy</h1>

Forxy is a fast HTTP proxy aggregator that forks the requests. It's main purpose is to bring a faster, more distributed approach to single threaded languages while maintaining stability and ease of use.

**Table of contents**
- [1. How to use it](#1-how-to-use-it)
- [2. Getting started](#2-getting-started)
  - [2.1 Docker installation](#21-docker-installation)
  - [2.2 Building the binary](#22-building-the-binary)
  - [2.3 Configuration file](#23-configuration-file)
    - [server](#server)
    - [log](#log)
    - [response](#response)
- [3. Known limitations](#3-known-limitations)

# 1. How to use it
  Forxy is designed to be a client side proxy, routing traffic parallely from the local network to the distributed world. It can sit on top of your microservices as well, with respect to network latency.
<br>

# 2. Getting started
## 2.1 Docker installation
   The forxy <a href="https://hub.docker.com/r/dragoscojocaru/forxy">docker image</a> makes it easy to just grep, test, and deploy it. 

  Integrate it in minutes with your docker compose stack using the following example:

    version: "3.4"

    services:
      forxy:
        image: dragoscojocaru/forxy
        container_name: forxy
        volumes:
          - ./forxy.yaml:/etc/forxy/forxy.yaml
        expose:
          - 8080


## 2.2 Building the binary
  Build your own binary from the source code running **go build** in the project root. The current version is **go 1.22.5**.
## 2.3 Configuration file
  The configuration file is named **forxy.yaml** and by default it is located in /etc/froxy/ path. The path can be changed using the FORXY_CONFIG_PATH environment variable.
### server
  The **server** config contains the following options:
  <li> <i>bind port</i>: binding port for the forxy http server.

### log
  The **log** config contains the following options:
  <li> <i>path</i>: path of the error log.

### response
  The **response** config contains the following options:
  <li> <i>validators</i>: string array containg name of the validators used in the response construct. The validators are run agains the target responses. Currently supported validators are:

    "content-type validator" - checks the target response against the Content-Type header. Forxy currently only supports "application/json" target responses, not including this validator will result in errors for other Content-Types.

<br>

# 3. Known limitations
  Forxy supports for now only HTTPS(S) application/json communication.