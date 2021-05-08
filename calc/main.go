// main.go
package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var widgets = []Widget{
		LineEdit{
			Text:      "0",
			TextColor: walk.RGB(255, 0, 0),
			//PasswordMode: true,
			//MaxLength: 12,
			//ReadOnly: true,
		},
	}
	var win = MainWindow{
		Title:    "go开发的计算器",
		MinSize:  Size{400, 500},
		Children: widgets,
		Layout:   VBox{},
	}
	win.Run()
}
