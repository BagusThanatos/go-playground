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

func popMessageBox()  {
  g.Msgbox("Title", "Press OK to close")
}

func loop() {
  window1 := g.Window("First Window")
  window2 := g.Window("Second window")

  layoutWin1 := g.Layout{
    g.Label("Hello world from giu"),
    g.Row(
      g.Button("Click Me").OnClick(onClickMe),
      g.Button("Do Something").OnClick(doSomething),
    ),
    g.PrepareMsgbox(),
    g.Button("show message box").OnClick(popMessageBox),
    g.Button("show message box inline").OnClick(func(){
      g.Msgbox("messge box inline", "This is using anonymous function")
    }),
  }

  layoutWin2 := g.Layout{
    g.Label("Second window here"),
  }

  window1.Layout(layoutWin1)
  window2.Layout(layoutWin2)
}

func main() {
  wnd := g.NewMasterWindow("Hello world", 800, 600, g.MasterWindowFlagsNotResizable)
  wnd.Run(loop)
}