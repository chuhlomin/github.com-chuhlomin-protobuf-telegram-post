# Protobuf – Telegram post

This is an additional material to Telegram post about [Protobuf](https://developers.google.com/protocol-buffers/): https://t.me/chuhlomin_channel/104

This project has a few pieces:

* [`writer`](cmd/writer/main.go) – Go application that that encodes Person struct into protobuf message and saves it to [`person.pb`](person.pb).
* [`reader`](cmd/reader/main.go) – Go application that reads [`person.pb`](person.pb) file and decodes protobuf message into Person struct.
* [`person.proto`](person.proto) – Person protobuf message definition.
* [`pkg/demo/person.pb.go`](pkg/demo/person.pb.go) – Go code generated from [`person.proto`](person.pb) with protoc3 (see below).
* [`person.json`](person.json) – just an example of JSON file.

## Running Go applications

In case you want to experiment with protobuf message encoding/decoding, you'd need Go.

Install [Go](https://golang.org/dl/).

```bash
$ go run cmd/writer/main.go

File [`person.pb`](person.pb) updated with content: 
John Doe�

$ go run cmd/reader/main.go
File content: name:"John Doe"  id:1234  has_ponycopter:true
```

## Updating generated Go code

In case you want to experiment with .proto files and code generation, you'd need Go, protoc3 and protoc-gen-go.

Install [Go](https://golang.org/dl/).

Download [protoc3](https://github.com/protocolbuffers/protobuf/releases/latest), add `bin` directory to PATH.

Install [protoc-gen-go](https://google.golang.org/protobuf/cmd/protoc-gen-go):

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

Run protoc3:

```bash
protoc3 -I=. --go_out=. person.proto
```
