package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func setupUI() {
	window := ui.NewWindow("Калькулятор склопакета", 400, 300, false)
	window.SetMargined(true)

	labelSize := ui.NewLabel("Розмір вікна")
	labelWidth := ui.NewLabel("Ширина, см")
	widthEntry := ui.NewEntry()
	labelHight := ui.NewLabel("Висота, см")
	hightEntry := ui.NewEntry()
	labelMaterial := ui.NewLabel("Матеріал")

	cbox := ui.NewCombobox()
	cbox.Append(("Дерево"))
	cbox.Append(("Метал"))
	cbox.Append(("Металопластик"))
	cbox.SetSelected(0)

	labelDoubleGlazing := ui.NewLabel("Склопакет")

	cbox1 := ui.NewCombobox()
	cbox1.Append(("Однокамерний"))
	cbox1.Append(("Двокамерний"))
	cbox1.SetSelected(0)

	checkBoxWindowSill := ui.NewCheckbox("Підвіконня")

	button := ui.NewButton("Розрахувати")
	costLabel := ui.NewLabel("")
	//name := ui.NewEntry()
	//button := ui.NewButton("Greet")
	//greeting := ui.NewLabel("")
	//cbox := ui.NewCombobox()
	//cbox.Append(("1. dsadsad"))
	//cbox.Append(("2. sdadsadsad"))

	vbox := ui.NewVerticalBox()
	vbox.Append(labelSize, false)
	vbox.Append(labelWidth, false)
	vbox.Append(widthEntry, false)
	vbox.Append(labelHight, false)
	vbox.Append(hightEntry, false)
	vbox.Append(labelMaterial, false)
	vbox.Append(cbox, false)
	vbox.Append(labelDoubleGlazing, false)
	vbox.Append(cbox1, false)
	vbox.Append(checkBoxWindowSill, false)
	vbox.Append(button, false)
	vbox.Append(costLabel, false)

	window.SetChild(vbox)

	button.OnClicked(func(*ui.Button) {

		widthValue, err := strconv.ParseFloat(widthEntry.Text(), 64)
		fmt.Println("Width ->", widthValue, err, reflect.TypeOf(widthValue))

		hightValue, err := strconv.ParseFloat(hightEntry.Text(), 64)
		fmt.Println("Hight ->", hightValue, err, reflect.TypeOf(hightValue))

		square := widthValue * hightValue
		var cost float64 = 0

		if cbox.Selected() == 0 && cbox1.Selected() == 0 {
			cost = square * 2.5
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		if cbox.Selected() == 0 && cbox1.Selected() == 1 {
			cost = square * 3
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		if cbox.Selected() == 1 && cbox1.Selected() == 0 {
			cost = square * 0.5
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		if cbox.Selected() == 1 && cbox1.Selected() == 1 {
			cost = square * 1
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		if cbox.Selected() == 2 && cbox1.Selected() == 0 {
			cost = square * 1.5
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		if cbox.Selected() == 2 && cbox1.Selected() == 1 {
			cost = square * 2
			if checkBoxWindowSill.Checked() {
				cost += 350
			}
		}

		fmt.Println("Square ->", square, "Cost ->", cost)

		costSting := fmt.Sprintf("%f", roundFloat(cost, 2))
		costLabel.SetText("Вартість ->" + costSting + " грн")
	})

	window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})

	window.Show()
}

func main() {
	ui.Main(setupUI)
}
