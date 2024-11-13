#!/bin/bash

mkdir -p bin
go build -o bin/main -gcflags "-m=3" . |& grep -i long --before-context=10 --after-context=5

