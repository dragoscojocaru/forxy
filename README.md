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
- [3. Developer guide](#3-developer-guide)
  - [Routes](#routes)
    - [ /http/fork](#-httpfork)
    - [ /http/sequential](#-httpsequential)
- [4. Known limitations](#4-known-limitations)

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
        ports:
          - "1480:1480"


## 2.2 Building the binary
  Build your own binary from the source code running **go build** in the project root. The current version is **go 1.22.5**.

  *! the docker image already contains the binary, manual build is not needed*
## 2.3 Configuration file
  The configuration file is named **forxy.yaml** and by default it is located in /etc/froxy/ path. The path can be changed using the FORXY_CONFIG_PATH environment variable.
### server
  The **server** config contains the following options:
  <li> <i>bind port (integer)</i>: binding port for the forxy http server.

    server:
      bind_port: 1480

### log
  The **log** config contains the following options:
  <li> <i>path (string)</i>: path of the error log.

    log:
      path: "/var/log/forxy/error.log"

<br>

# 3. Developer guide
Forxy exposes a HTTP server for communication.

## Routes

### <li> <b>/http/fork<b> 

This is the route that forks your requests into the distributed world 🙂. To make sure both stability and functionality are maintained, this route uses a dedicated request payload:

<br>

**request payload:**

    {
      "timeout": 400,
      "requests":
      {
          "1":{
              "url": "https://integration-1.forxy.dev/",
              "method": "GET",
              "headers": {
                "key1": "value1"
              },
              "body": {
                  "count": 2,
                  "category": "Time"
              }
          },
          "2":{
              "url": "https://integration-2.forxy.dev/",
              "method": "GET",
              "body": {
                  "count": 1,
                  "category": "Comedy"
              }
          },
          "3":{
              "url": "https://integration-1.forxy.dev/",
              "method": "GET",
              "body": {
                  "count": 1,
                  "category": "Time"
              }
          }
      }
    }

*timeout (integer)* - applied as a global timeout in miliseconds for the target requests.

<br>

*requests (object)* - json object containing the target request index as *key(string)* & target request data as *value(object)*

<br>

*target request* - object containg data for the target http request. 

*url (string)* - target url

*method (string)* - target HTTP method

*headers (object), optional* - target http headers object containg multiple *key(string)*-*value(string)* pairs. 

*body (object), optional* - target HTTP request body.

<br>

**response payload:**

    {
      "responses": {
          "1": {
              "forxy_control": {
                  "ok": true,
                  "message": "Forxy pass OK."
              },
              "status": 200,
              "body": {
                  "count": 2,
                  "quotes": [
                      {
                          "category": "Time",
                          "quote": {
                              "quote": "Lost time is never found again.",
                              "author": "Benjamin Franklin"
                          }
                      },
                      {
                          "category": "Time",
                          "quote": {
                              "quote": "The two most powerful warriors are patience and time.",
                              "author": "Leo Tolstoy"
                          }
                      }
                  ]
              }
          },
          "2": {
              "forxy_control": {
                  "ok": true,
                  "message": "Forxy pass OK."
              },
              "status": 200,
              "body": {
                  "count": 1,
                  "movies": [
                      {
                          "category": "Comedy",
                          "movie": {
                              "name": "Groundhog Day",
                              "release_year": 1993,
                              "description": "A weatherman finds himself living the same day over and over again, forcing him to reevaluate his life and actions."
                          }
                      }
                  ]
              }
          },
          "3": {
              "forxy_control": {
                  "ok": true,
                  "message": "Forxy pass OK."
              },
              "status": 200,
              "body": {
                  "count": 1,
                  "quotes": [
                      {
                          "category": "Time",
                          "quote": {
                              "quote": "Time is what we want most, but what we use worst.",
                              "author": "William Penn"
                          }
                      }
                  ]
              }
          }
      }
    }


<br>

*responses (object)*: contains multiple pairs of target request index *key(str)* & target response data *value(object)*;

<br>

*forxy_control (object)*: object used for error handling inside the forxy environment;

*ok (boolean)*: success status of the target request execution;

*message (string)*: satus message the target request execution;

<br>

*status (integer)* - target HTTP response status code;

<br>

*body (object)* - target HTTP response body;

<br>
<br>

*! Forxy only supports application/json communication for now. request and response headers are passed automatically by the application, so manual handling is not necessary.*


### <li> <b>/http/sequential<b> 


Route used for testing and development purposes. It uses the same request payload as the forxy route, but executes the target http requests sequentially.

<br>

# 4. Known limitations
  Forxy supports for now only HTTPS(S) application/json communication.
