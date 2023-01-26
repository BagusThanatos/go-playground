package main

import (
  "encoding/csv"
  "fmt"
  "log"
  "os"

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

func generateTable() *g.TableWidget{
  data := loadCSV("data.csv")
  rows := make([]*g.TableRowWidget, len(data))
  for index, value := range data {
    rows[index] = g.TableRow(g.Label(value[0]), g.Label(value[1]), g.Label(value[2]))
  }
  return g.Table().Columns(g.TableColumn("column1"), g.TableColumn("column2"), g.TableColumn("last Column")).
          Rows(rows...)
}

func loadCSV(path string) [][]string{
  file, err := os.Open(path)
  if err != nil {
    log.Fatal("Error opening file: ", err)
    return nil
  }
  reader := csv.NewReader(file)
  result, err := reader.ReadAll()
  if err != nil {
    log.Fatal("Error parsing csv file: ", err)
  }
  return result
}

func loop() {
  window1 := g.Window("First Window")
  window2 := g.Window("Second window")
  window3 := g.Window("Table Window")

  layoutWin1 := g.Layout{
    g.Label("Hello world from giu"),
    g.Labelf("Window 1 has focus: %t", window1.HasFocus()),
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
    g.Labelf("Window 2 has focus: %t", window2.HasFocus()),
  }

  layoutWin3 := g.Layout{generateTable()}

  window1.Layout(layoutWin1)
  window2.Layout(layoutWin2)
  window3.Layout(layoutWin3)
}

func main() {
  wnd := g.NewMasterWindow("Hello world", 800, 600, g.MasterWindowFlagsNotResizable)
  wnd.Run(loop)
}