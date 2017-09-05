A POC to try out go-micro to compare with go-kit

## Prereqs

### Consul

Install consul
```shell
brew install consul
```

Run Consul

```shell
consul agent -dev
```

### Protobuf

```shell
brew install protobuf
go get github.com/micro/protobuf/{proto,protoc-gen-go}
```

## Installation

```shell
dep ensure
```

Build
```shell
./build.sh
```

## Usage

Run the service

```shell
./bin/service
```

Run the client

```shell
./bin/client
```
