package main

import (
	"fmt"
	"log"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var input *walk.LineEdit
	var label *walk.Label

	MainWindow{
		Title:   "Saludo con Walk",
		MinSize: Size{300, 200},
		Layout:  VBox{},
		Children: []Widget{
			LineEdit{
				AssignTo: &input,
				Text:     Bind("nombre"),
			},
			PushButton{
				Text: "Listo",
				OnClicked: func() {
					nombre := input.Text()
					if nombre != "" {
						walk.MsgBox(nil, "Saludo", fmt.Sprintf("Hola %s", nombre), walk.MsgBoxIconInformation)
					}
				},
			},
			Label{
				AssignTo: &label,
				Visible:  false,
			},
		},
	}.Run()
}
