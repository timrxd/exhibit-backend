# exhibit-backend

Simple backend server to to provide data for [exhibit-frontend](https://github.com/timrxd/exhibit-frontend) (coming soon™).

Simple API based in Go.

## Setup
 - Make sure the following are installed:
    - [Go](https://go.dev/doc/install)
        - run ```go install tool``` to build cli tools 
        - make sure your GOBIN is on your $PATH
        - ```export PATH=`go env GOPATH`/bin/:$PATH```
    - [buf](https://buf.build/docs/cli/installation/)
    - [protobuf](https://protobuf.dev/installation/)

## Run
 - ```./scripts/run.sh```
 - REST gateway is located at http://localhost:8090/v1
    - ex:  
        ```curl -X GET -k http://localhost:8090/v1/player/Federer```

## API
Endpoints:
 - GET /v1/player/{name}
 - GET /v1/players
 - POST /v1/player
 - DELETE /v1/player/{name}  
  
Player:
 - string Name
 - string Country
 - int32 Age
 - int32 Points