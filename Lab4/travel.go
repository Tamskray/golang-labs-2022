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

	labelTours := ui.NewLabel("Путівки")
	labelTickets := ui.NewLabel("Кількість путівок")
	ticketsEntry := ui.NewEntry()
	labelDays := ui.NewLabel("Кількість днів")
	daysEntry := ui.NewEntry()
	labelCountry := ui.NewLabel("Країна")

	cbox := ui.NewCombobox()
	cbox.Append(("Болгарія"))
	cbox.Append(("Німеччина"))
	cbox.Append(("Польща"))
	cbox.SetSelected(0)

	labelSeason := ui.NewLabel("Сезон")

	cbox1 := ui.NewCombobox()
	cbox1.Append(("Літо"))
	cbox1.Append(("Зима"))
	cbox1.SetSelected(0)

	checkBoxGuide := ui.NewCheckbox("Індивідуальний гід")

	checkBoxLuxHotel := ui.NewCheckbox("Проживання в номері люкс")

	button := ui.NewButton("Розрахувати")
	costLabel := ui.NewLabel("")
	//name := ui.NewEntry()
	//button := ui.NewButton("Greet")
	//greeting := ui.NewLabel("")
	//cbox := ui.NewCombobox()
	//cbox.Append(("1. dsadsad"))
	//cbox.Append(("2. sdadsadsad"))

	vbox := ui.NewVerticalBox()
	vbox.Append(labelTours, false)
	vbox.Append(labelTickets, false)
	vbox.Append(ticketsEntry, false)
	vbox.Append(labelDays, false)
	vbox.Append(daysEntry, false)
	vbox.Append(labelCountry, false)
	vbox.Append(cbox, false)
	vbox.Append(labelSeason, false)
	vbox.Append(cbox1, false)
	vbox.Append(checkBoxGuide, false)
	vbox.Append(checkBoxLuxHotel, false)
	vbox.Append(button, false)
	vbox.Append(costLabel, false)

	window.SetChild(vbox)

	button.OnClicked(func(*ui.Button) {

		ticketsValue, err := strconv.ParseFloat(ticketsEntry.Text(), 64)
		fmt.Println("Width ->", ticketsValue, err, reflect.TypeOf(ticketsValue))

		daysValue, err := strconv.ParseFloat(daysEntry.Text(), 64)
		fmt.Println("Hight ->", daysValue, err, reflect.TypeOf(daysValue))

		trip := ticketsValue * daysValue
		var cost float64 = 0

		if cbox.Selected() == 0 && cbox1.Selected() == 0 {
			cost = trip * 100

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		if cbox.Selected() == 0 && cbox1.Selected() == 1 {
			cost = trip * 150

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		if cbox.Selected() == 1 && cbox1.Selected() == 0 {
			cost = trip * 160

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		if cbox.Selected() == 1 && cbox1.Selected() == 1 {
			cost = trip * 200

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		if cbox.Selected() == 2 && cbox1.Selected() == 0 {
			cost = trip * 120

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		if cbox.Selected() == 2 && cbox1.Selected() == 1 {
			cost = trip * 180

			markUpCheck := false
			// націнка з гідом
			if checkBoxGuide.Checked() {
				cost += 50
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
					markUpCheck = true
				}
			}

			// націнка без гіда
			if markUpCheck == false {
				if checkBoxLuxHotel.Checked() {
					markUp := (cost * 20) / 100
					cost += markUp
				}
			}
		}

		fmt.Println("trip ->", trip, "Cost ->", cost)

		costSting := fmt.Sprintf("%f", roundFloat(cost, 2))
		costLabel.SetText("Вартість ->" + costSting + " $")
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
