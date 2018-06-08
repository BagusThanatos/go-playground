@echo off
set GOPATH=%cd%\
set GOBIN=%cd%\bin\
echo "RUN TEST"
go test .\src\bagus\... -v

go build -o bin\main.exe  .\src\bagus\...
bin\main.exe
