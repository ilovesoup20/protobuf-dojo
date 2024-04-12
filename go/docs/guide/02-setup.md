# Setup

Guide for using protobuf with using Go language.

## Prerequisites

- Go (already installed)
- Protobuf Compiler (already installed)

## Steps

### Install protobuf compiler

https://github.com/protocolbuffers/protobuf/releases

### Install the Go Protobuf plugin

```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

```

### Define Protobuf Schema

```proto
syntax = "proto3";

package tutorial;

// Defines a Person message with some fields.
message Person {
  string name = 1;
  int32 id = 2;  // Unique ID number for this person.
  string email = 3;
}

```

### Generate Go Code

```sh
protoc --go_out=. --go_opt=paths=source_relative person.proto

```


## Self-notes

- `proto` files are defined in `api/proto` folder
- business logic code is stored under `cmd`
- compiled `proto` files are stored inside `pkg/pb`

### Importing compiled protobuf code

- compiled protobuf files reside under `pkg/pb`
- `person.proto` uses package name `tutorial`

**Use the following import statement**

```
import tutorial "protobuf-dojo/pkg/pb"
```

Note the import path specifies the directory where the compiled code lives. Then `tutorial` package is imported from that namesapce.

### Compiling proto files

```
protoc --proto_path=api/proto/ --go_out=pkg/pb --go_opt=paths=source_relative api/proto/*.proto
```