#!/usr/bin/env bash
GOPATH=`pwd` go get -v app/...
rm bin/*