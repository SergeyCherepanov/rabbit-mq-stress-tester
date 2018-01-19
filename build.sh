#!/bin/bash

for GOOS in linux darwin; do
   for GOARCH in amd64; do
     docker run --rm -e GOOS=$GOOS -e GOARCH=$GOARCH -v "$PWD":/go/src/github.com/SergeyCherepanov/rabbit-mq-stress-tester -v "$GOPATH"/src:/go/src -w /go/src/github.com/SergeyCherepanov/rabbit-mq-stress-tester/ golang:1.8 go build -v -o /go/src/github.com/SergeyCherepanov/rabbit-mq-stress-tester/rabbit-mq-stress-tester-$GOOS-$GOARCH
   done
done