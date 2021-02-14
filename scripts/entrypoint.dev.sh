#!/bin/bash

set -e

go run cfg/db/migrate/main.go

GO111MODULE=off go get github.com/githubnemo/CompileDaemon

CompileDaemon --build="go build -o main main.go" --command ./main