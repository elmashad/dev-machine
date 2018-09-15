#!/bin/bash

go get -v ./...

go get -v github.com/codegangsta/gin

gin -i -p 3333 run main.go
