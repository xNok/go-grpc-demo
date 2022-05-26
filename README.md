# Go gRPC demo

This demo gRPC application is a simple note-taking application. You can use the client to:
* Save a note by providing a title and a content
* Load a note by searching a keyword

Fully tutorial can be found here: https://speedscale.com/2022/05/03/using-grpc-with-golang/

## How to get started?

You will need `protoc` to generate proto buffer code.

```
sudo apt update
sudo apt install protobuf-compiler
```

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
export PATH="$PATH:$(go env GOPATH)/bin"
```

## How to test it?

Run the server

```
go run ./notes_server/main.go
```

Use the client to interact with the server.

Save a note:

```
go run notes_client/main.go save -title test -content "Lorem ipsum dolor sit amet, consectetur "
```

Load a note:

```
go run notes_client/main.go load -keyword Lorem
```
