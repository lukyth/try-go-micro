#!/bin/bash

cd proto && protoc --go_out=plugins=micro:. *.proto && cd ..
go build -o bin/client cmd/client/client.go
go build -o bin/service cmd/service/service.go