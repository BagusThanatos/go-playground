@echo off
set GOPATH=%cd%\src\
set GOBIN=%cd%\bin\
echo "RUN TEST"
go test ./... -v

go build -o bin\main.exe  .\src\...
bin\main.exe
