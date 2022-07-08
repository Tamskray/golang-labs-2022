package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Facts struct {
	Title    string
	FactInfo string
}

type InfoFacts struct {
	Data []Facts
}

type VariablesAbout struct {
	AboutCity  string
	Attraction []string
	Cities     []string
}

type City struct {
	Name       string
	Population int
	Areas      []string
	Media      []string
}

type News struct {
	News []string
}

type Weather struct {
	Days       string
	DegreesMin int
	DegreesMax int
	Wind       float64
	Pressure   int
	Humidity   int
}

type InfoWeather struct {
	Data []Weather
}

const tmplFile = "main.tpl"
const tmplFileHomePage = "homePage.html"
const tmplFileFacts = "facts.html"
const tmplFileAbout = "aboutPage.html"
const tmplFileLastNews = "lastNews.html"

const tmlFileWeather = "weather.html"

const (
	anError = `<p class="error">%s</p>`
)

type Solution string

var HttpSolution Solution = "Site"

func (s Solution) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/" {
		zhytomyr := City{"Житомир", 261624, []string{"Крошня", "Мальованка", "Польова", "Корбутівка", "Смолянка", "Центр"}, []string{"«Житомирщина»", "«Місто»", "«20 Хвилин»", "«Житомир info»"}}

		tmpl, err := template.ParseFiles(tmplFileHomePage)
		templates := template.Must(tmpl, err)
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		templates.ExecuteTemplate(w, tmplFileHomePage, zhytomyr)
	}

	if r.URL.Path == "/facts" {
		data := InfoFacts{
			Data: []Facts{
				{Title: "Область", FactInfo: "Житомирська область посідає 5 місце за площею серед усіх регіонів України, маючи загальну площу 29 832 км². Це більше за такі країни, як Вірменія чи Ізраїль."},
				{Title: "Соколовський кар’єр", FactInfo: "Соколовський кар’єр, який знаходиться поблизу Житомира, своєю глибиною не поступається озеру Байкал"},
				{Title: "Морозиво", FactInfo: "Кожна третя пачка морозива в Україні виготовлена в нашому місті."},
				{Title: "Театр", FactInfo: "Житомирська обласна філармонія ім. Святослава Ріхтера – найстаріший театр України."},
				{Title: "Валюта", FactInfo: "Житомирщина – Батьківщина нашої національної валюти. Саме тут Центральна Рада прийняла закон про введення грошової одиниці – гривні."},
				{Title: "Фабрика", FactInfo: "Малинська фабрика виробляє банкнотний папір для всієї Європи. Подібних підприємств лише 50 у світі."},
				{Title: "Трамвай", FactInfo: "Житомир був третім містом України, в якому з’явився трамвай. Спочатку він був грузовим, та за десятки років цей вид транспорту став на стільки популярним, що у Житомирі діяло 5 трамвайних маршрутів."},
				{Title: "Лампочка", FactInfo: "Першу електричну лампочку запалили у Житомирському міському театрі 121 рік тому."},
			},
		}

		tmpl, err := template.ParseFiles(tmplFileFacts)
		templates := template.Must(tmpl, err)
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		//templates.ExecuteTemplate(w, tmplFileHomePage, bob)
		templates.ExecuteTemplate(w, tmplFileFacts, data)
	}

	if r.URL.Path == "/about" {
		AboutCity := VariablesAbout{"Місто на півночі України, адміністративний центр Житомирської області та Житомирського району, центр Житомирської агломерації. До 2020 року був містом обласного підпорядкування. Ділиться на Богунський та Королівський райони.",
			[]string{"Кафедральний собор св. Софії (1737 р.)", "Церква св. Михайла (1856 р.)", "Підвісний міст", "Вулиця Михайлівська (при Свято-Михайлівському кафедральному соборі)",
				"Старий театр (польський)", "Старий єврейський цвинтар (XIX ст.)"},
			[]string{"Бердичів", "Звягель", "Овруч", "Коростень", "Коростишів", "Малин"}}

		tmpl, err := template.ParseFiles(tmplFileAbout)
		templates := template.Must(tmpl, err)
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		templates.ExecuteTemplate(w, tmplFileAbout, AboutCity)
	}

	if r.URL.Path == "/lastNews" {
		lastNewsinCity := News{[]string{
			"За півтора місяця прибирань тротуарів у Житомирі комунальники готові заплатити 2,3 млн грн",
			"Артилерія житомирської бригади ДШВ продовжує нищити бронетехніку російських окупантів", "Щоп’ятниці у гідропарку житомиряни можуть безкоштовно отримати інвентар для занять спортом",
			"Робочий тиждень у Житомирі розпочався зі спеки, але через два дні погода змінить характер",
			"У Житомирі на Михайлівській провели благодійну акцію «Віват, ліцей 25»"}}

		tmpl, err := template.ParseFiles(tmplFileLastNews)
		templates := template.Must(tmpl, err)
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		templates.ExecuteTemplate(w, tmplFileLastNews, lastNewsinCity)
	}

	if r.URL.Path == "/weather" {
		/*
			wtr := Weather{[]string{"Понеділок", "Вівторок", "Середа", "Четвер", "П'ятниця", "Субота", "Неділя"},
				[]int{14, 12, 14, 16, 15, 12, 14}, []int{17, 20, 20, 23, 23, 21, 31}, []float64{1.4, 2.4, 3.2, 1.6, 1.3, 3.2, 0.3},
				[]int{737, 740, 738, 742, 746, 740, 742}, []int{76, 87, 64, 99, 93, 79, 82}}
		*/
		data := InfoWeather{
			Data: []Weather{
				{Days: "Понеділок", DegreesMin: 14, DegreesMax: 17, Wind: 1.4, Pressure: 737, Humidity: 76},
				{Days: "Вівторок", DegreesMin: 12, DegreesMax: 20, Wind: 2.4, Pressure: 740, Humidity: 87},
				{Days: "Середа", DegreesMin: 14, DegreesMax: 23, Wind: 3.2, Pressure: 739, Humidity: 64},
				{Days: "Четвер", DegreesMin: 16, DegreesMax: 23, Wind: 1.6, Pressure: 738, Humidity: 99},
				{Days: "П'ятниця", DegreesMin: 15, DegreesMax: 21, Wind: 1.3, Pressure: 746, Humidity: 93},
				{Days: "Субота", DegreesMin: 12, DegreesMax: 20, Wind: 3.2, Pressure: 740, Humidity: 79},
				{Days: "Неділя", DegreesMin: 14, DegreesMax: 31, Wind: 0.3, Pressure: 742, Humidity: 82},
			},
		}
		tmpl, err := template.ParseFiles(tmlFileWeather)
		templates := template.Must(tmpl, err)
		err = r.ParseForm()

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		templates.ExecuteTemplate(w, tmlFileWeather, data)
	}

}

func main() {
	//http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	http.ListenAndServe("localhost:80", HttpSolution)
}
