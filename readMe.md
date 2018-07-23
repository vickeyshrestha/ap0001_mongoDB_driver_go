**1. Introduction**

This application connects to mongoDb to fetch the KV configurations for the client. The client can request for configs using HTTP interface.

<br>

**2. Requirement**

A mongo DB must be running in one of the servers and the mongo must have a database name "config" and a collection name "vic_application". Yup, that's mandatory for now. That's where all of the application config will be stored.

<br>

**3. Execute the following for building docker image and running image:**

    - go clean
    - CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001-mongoDB-driver cmd/main.go
    - docker build -t vickeyshrestha/ap0001-mongodriver-go:00.00.01 .
    - docker run --restart=always -e mongoHostAndPort=192.168.202.131:27017 -p 8085:8085 -d vickeyshrestha/ap0001-mongodriver-go:00.00.01
    - docker push vickeyshrestha/ap0001-mongodriver-go:00.00.01

`**NOTE**`    

You can simply get this docker image from 

    https://hub.docker.com/u/vickeyshrestha/

_I prefer Artifactory as a docker hub, but I don't know if its free :(_

<br>

**4. Application Endpoints**

Here the the endpoints for this application:

| EndPoint Syntax        | Example           | Detail  |
| ------------- |:-------------:| -----:|
| /health      | http://192.168.202.131:8085/health | Get the health status of this application |
| /getallconfigs      | http://192.168.202.131:8085/getallconfigs      |   Gets whole data response from collection |
| /getconfig?app=<APPLICATION_NAME>&bin=<BINARY_VERSION>&site=<SITE_NAME> | http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev      |    Returns the document based on mandatory parameters. The mandatory parameters are app, bin and site |


