#!/bin/bash -e
cd $(dirname $0)
PATH=$HOME/gopath/bin:$GOPATH/bin:$PATH

function announce
{
  echo
  echo $@
}

function v
{
  announce $@
  $@
}

v gofmt -l -w *.go

v go test -v -covermode=count -coverprofile=cover.out .
v go tool cover -func=cover.out
rm cover.out

v go vet .
