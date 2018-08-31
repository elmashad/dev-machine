#!/bin/bash

go get -d -v ./...

go get -v github.com/codegangsta/gin


gin -i -p 3333 run main.go
