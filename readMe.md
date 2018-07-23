**1. Introduction**

This application connects to mongoDb to fetch the KV configurations for the client.

<br>

**2. Execute the following for building docker image and running image:**

    - go clean
    - CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001-mongoDB-driver cmd/main.go
    - docker build -t vickeyshrestha/ap0001-mongodriver-go:00.00.01 .
    - docker run --restart=always -p 8085:8085 -d vickeyshrestha/ap0001-mongodriver-go:00.00.01
    - docker push vickeyshrestha/ap0001-mongodriver-go:00.00.01

`**NOTE**`    

You can simply get this docker image from 

    https://hub.docker.com/u/vickeyshrestha/

_I prefer Artifactory as a docker hub, but I don't know if its free :(_

<br>

**3. Application Endpoints**

Here the the endpoints for this application:

| EndPoint Syntax        | Example           | Detail  |
| ------------- |:-------------:| -----:|
| /health      | http://192.168.202.131:8085/health | Get the health status of this application |
| /allconfigs      | http://192.168.202.131:8085/allconfigs      |   Gets whole data response from collection |
| /getconfig?app=testApplication&bin=0.0.2&site=dev | http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev      |    Returns the document based on mandatory parameters. The mandatory parameters are app, bin and site |


