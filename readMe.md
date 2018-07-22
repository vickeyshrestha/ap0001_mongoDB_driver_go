Execute the following for building docker image and running image:

<br>

go clean
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001-mongoDB-driver cmd/main.go

docker build -t docker.vickey.com/ap0001/mongodriver-go:00.00.01 .

docker run --restart=always -p 8085:8085 -d docker.vickey.com/ap0001/mongodriver-go:00.00.01