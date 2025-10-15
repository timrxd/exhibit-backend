# exhibit-backend

Simple backend server to to provide data for [exhibit-frontend](https://github.com/timrxd/exhibit-frontend) (coming soon™).

Simple API based in Go.

## Setup
 - Make sure the following are installed:
    - [Go](https://go.dev/doc/install)
        - make sure your GOBIN is on your $PATH
        - ```export PATH=`go env GOPATH`/bin/:$PATH```
    - [buf](https://buf.build/docs/cli/installation/)
    - [protobuf](https://protobuf.dev/installation/)

## Run
 - ```./run.sh```
 - REST gateway is located at http://localhost:8090/v1
    - ex:  
        ```curl -X POST -k http://localhost:8090/v1/getPlayer -d '{"name": "Federer"}'```

