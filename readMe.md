Execute the following for building docker image and running image:

<br>

go clean

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ap0001-mongoDB-driver cmd/main.go

docker build -t vickeyshrestha/ap0001-mongodriver-go:00.00.01 .

docker run --restart=always -p 8085:8085 -d vickeyshrestha/ap0001-mongodriver-go:00.00.01

NOTE

You can simply get this docker image from https://hub.docker.com/u/vickeyshrestha/

I prefer Artifactory as a docker hub, but I don't know if its free :(