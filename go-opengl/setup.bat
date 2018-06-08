@echo off

set GOPATH=%cd%\
go get -v github.com/go-gl/gl/v4.6-core/gl
go get -v github.com/go-gl/glfw/v3.2/glfw