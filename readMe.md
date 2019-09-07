**1. Introduction**

This application connects to mongoDb to fetch the KV configurations for the client. The client can request for configs using HTTP interface.

<br>

**2. Requirement / prerequisite**

A mongoDB must be running in one of the servers and the mongo must have a database name "config" and a collection name "vic_application". Yup, that's mandatory for now. That's where all of the application config will be stored. The FQDN or IP Address of the server and the port number where MongoDB is serving must be passed as environment variable. See Section 3 and section 7 of this document for detail and example.

<br>

For required external dependencies, check **vendor** directory.

<br>

**3. Execute the following for building docker image and running image:**

    - go clean
    
    - CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001_mongo_engine cmd/main.go
    
    - docker build -t vickeyshrestha/ap0001_mongo_engine:<IMAGE_VERSION> .
    - EXAMPLE: docker build -t vickeyshrestha/ap0001_mongo_engine:00.00.01 .
    
    - docker run --restart=always -e mongoHostAndPort=<MONGODB_HOST>:<PORT> -p <PORT_TO_MAP>:8085 -d vickeyshrestha/ap0001_mongo_engine:<IMAGE_VERSION>
    - EXAMPLE: docker run --restart=always -e HOSTNAME=vickey_ubuntu_1 --name=MONGODRIVER -e mongoHostAndPort=192.168.@@@.@@@:27017 -p 8085:8085 -d vickeyshrestha/ap0001_mongo_engine:00.00.01
    
        Note: To run the container in SSL secured mode, refer to section 10 of this readme file.
    
    - Verify that your env is properly set by entering the container. You can enter using following command:
        $ docker exec -it <CONTAINER_ID> /bin/bash
    
    - docker push vickeyshrestha/ap0001_mongo_engine:00.00.01

`**NOTE**`    

You can simply get this docker image from 

    https://hub.docker.com/u/vickeyshrestha/

<br>

**4. Application Endpoints**

Here the the endpoints for this application:

| EndPoint Syntax        | Example           | Detail  |
| ------------- |:-------------:| -----:|
| /health      | http://192.168.@@@.@@@:8085/health | Get the health status of this application |
| /getallconfigs      | http://192.168.@@@.@@@:8085/getallconfigs      |   Gets whole data response from collection |
| /getconfig?app=<APPLICATION_NAME>&bin=<BINARY_VERSION>&site=<SITE_NAME> | http://localhost:8085/getconfig?app=testApplication&bin=0.0.2&site=dev      |    Returns the document based on mandatory parameters. The mandatory parameters are app, bin and site |

<br>

**5. Insert a new config record**

The endpoint to insert new record is `/insertnew`. It is a POST request and accepts JSON body. For example:

    Example: 
        Type-> POST
        Http Interface-> http://localhost:8085/insert
        Body ->
            {
                "applicationName": "testApplication",
                "binaryVersion": "0.0.2",
                "httpConnectionTimeout": 5,
                "servingPort": 8099,
                "site": "dev"
             }
             
     Note: The fields applicationName, binaryVersion and site are mandatory. Other fields are up to developers on how they want to add them.
     
     - You can use your favorite HTTP API Interacting tool. My favorite is Postman. 

<br>

**6. Delete record(s))**

To delete a record, simply use the endpoint `/delete`. It is a request of DELETE type. You will need to provide 3 parameters.
    
    Example:
        Type -> DELETE
        HTTP request -> http://localhost:8085/delete?app=testApplication&bin=0.0.2&site=dev
        
    Note: The parameters app, bin and site are mandatory that represents applicationName, binaryVersion and site respectively.
    
    - You can use your favorite HTTP API Interacting tool. My favorite is Postman. 

<br>

**7. Running in non-containerized environment** 

If you simply want to run this application in its non-containerized form, i.e. just through cmd/main.go, you have to setup the follwoing two environment variables:

    - configFile
        - Example -> configFile=C:/Projects-Golang/src/ap0001_mongo_engine/resources/config.json
        - Description -> An env var pointing to the location of the config file. This is found inside the resources directory of the project.
        
    - mongoHostAndPort
        - Example -> some_server:27017
        - Description -> FQDN and the port where MongoDB is running
        
<br>

**8. Vendorizing dependencies**

_`prerequisite: Please install govendor tool available at https://github.com/kardianos/govendor`_

Remember to vendorize your dependencies as well:

TO ADD:

govendor init

govendor add +external


TO REMOVE ALL/ CLEAN VENDOR DIRECTORY:

govendor remove +v

<br>

**9. To use behavior driven test (BDT) and running Unit test**

_`prerequisite: Please install godog available at https://github.com/DATA-DOG/godog`_

Navigate to folder ap0001_mongo_engine\behaviorDrivenTest and Use the following godog command:

$ godog testMongoDriver.feature

<br>

I have tried to use the concept of Dependency Injection using interfaces. Unit tests in this application uses a tool called Mockery to mock up the interface. Download it using 

_`go get github.com/vektra/mockery/.../`_

<br>

Now in future, to auto create the mock of the interface, navigate to interface directory and use the following command to create a mock of that interface:

_`mockery -name <INTERFACE_NAME>`_

<br>

**10. Running application in SSL secured mode**

To run this application serving a https secured connection, you need to pass the following 3 additional environment variables:

    1. devmode  (Set it to false. This will enable the https server)
    2. sslKey   (Path to the ssl private Key)
    3. sslCert  (Path to the ssl private certificate)
    
The certificates and keys are confidential files and will not be commited to github. In fact, each server requires their own certificate and key. Each company might have their own csr signing tool. But for development purpose, you can use the free certificate signing tool at 
    
    http://getacert.com/signacert.html

You can generate your own csr file and key using free tools like Openssl. Then use that csr in the above link to generate a free 60 days signed cert. Donate to the site to get for 10 years validation.

When running a docker image, I would recommend the newly generated cert to place somewhere locally in the server. Then you can volume map the docker conatiner's location to the local host location using -v option. For example, your certs and keys are located under /server/certificate folders. You can do something like:

    docker run --restart=always -e HOSTNAME=vickey_ubuntu_1 --name=MONGODRIVER -v /server/certificate:/server/certificate -e devmode=false -e sslCert=/server/certificate/server.cer -e sslKey=/server/certificate/server.key -e mongoHostAndPort=192.168.@@@.@@@:27017 -p 8085:8085 -d vickeyshrestha/ap0001_mongo_engine:00.00.01


**Remember** that you need to install the public certificate to your OS cert repo or your browser repository also. When you generate the certificate using the link above, it should provide the public certificate as well.  