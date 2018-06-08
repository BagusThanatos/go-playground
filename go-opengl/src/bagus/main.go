package main

import (
	"fmt"
	"runtime"
	//"log"

	//"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	
)

const (
	width = 800
	height = 640
)

func main() {
	fmt.Println("Hello World")
	
	runtime.LockOSThread()
	window := initGlfw()
	defer glfw.Terminate()

	for !window.ShouldClose() {
		
	}
}

func initGlfw() *glfw.Window {
	if err := glfw.Init(); err != nil {
		panic(err)
	}
	
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 6)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Bagus Thanatos", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	
	return window
}