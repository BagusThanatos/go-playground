package main

import (
  "fmt"

  g "github.com/AllenDang/giu"
)

func onClickMe() {
  fmt.Println("Hello world!")
}

func doSomething() {
  fmt.Println("Do something")
}

func loop() {
  g.SingleWindow().Layout(
    g.Label("Hello world from giu"),
    g.Row(
      g.Button("Click Me").OnClick(onClickMe),
      g.Button("Do Something").OnClick(doSomething),
    ),
  )
}

func main() {
  wnd := g.NewMasterWindow("Hello world", 800, 600, g.MasterWindowFlagsNotResizable)
  wnd.Run(loop)
}