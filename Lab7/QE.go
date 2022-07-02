package main

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
)

const (
	pageHeader = `<!DOCTYPE HTML>
	<html>
	<head>
	<title>Квадратне рівняння</title>
	<style>
	*{
		padding: 5px;
	}
	.error{
	color:#FF0000;
	}
	</style>
	</head>`
	pageBody = `<body>
	<h3>Квадратне рівняння</h3>`
	form = `<form action="/" method="POST">
	<label for="numbers">Введіть коефіцієнти рівняння:</label><br /><br />
	<input type="text" name="a" size="2" value="0"> x^2
	 + <input type="text" name="b" size="2" value="0"> x
	 + <input type="text" name="c" size="2" value="0">
	 = 0
	<p><input type="submit" value="Виконати"></p>
	</form>`
	anError = `<p class="error">%s</p>`
)

type Solution string

var HttpSolution1 Solution = "Квадратне рівняння"

func QuadraticEquation(a, b, c float64) []float64 {
	var d, x1, x2 float64
	results := []float64{}
	d = b*b - 4*a*c
	if d > 0 {
		x1 = ((-b) + math.Sqrt(d)) / (2 * a)
		x2 = ((-b) - math.Sqrt(d)) / (2 * a)
		fmt.Println("x1 ->", x1, "\nx2 ->", x2)
		results = append(results, x1, x2)
	}
	if d == 0 {
		x1 = -(b / (2 * a))
		fmt.Println("x1 ->", x1)
		results = append(results, x1)
	}
	if d < 0 {
		fmt.Println("D < 0, рівняння немає дійсних рішень")
	}
	return results
}

func (s Solution) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, pageHeader, pageBody, form) // Формируем страницу в браузере
	if r.Method == "POST" {                   // Обрабатываем входные данные
		err := r.ParseForm() // Парсим форму
		post := r.PostForm

		if err != nil {
			fmt.Fprintf(w, anError, err)
			return
		}

		pa := post.Get("a")
		pb := post.Get("b")
		pc := post.Get("c")
		a, _ := strconv.ParseFloat(pa, 64)
		b, _ := strconv.ParseFloat(pb, 64)
		c, _ := strconv.ParseFloat(pc, 64)
		fmt.Fprintf(w, "a = %v	b = %v	c = %v", a, b, c) // Выводим сообщение в браузер
		fmt.Fprintf(w, "\n\t<br /><br />")
		//fmt.Println(reflect.TypeOf(a))
		res := QuadraticEquation(a, b, c)
		//fmt.Println(res)
		if len(res) == 2 {
			fmt.Fprintf(w, "<h3>x1 = %v, x2 = %v</h3>", res[0], res[1])
		}
		if len(res) == 1 {
			fmt.Fprintf(w, "<h3>x = %v</h3>", res[0])
		}
		if len(res) == 0 {
			fmt.Fprintf(w, "<h3>D<0, рівняння немає дійсних рішень</h3>")
		}
	}
	if r.Method == "GET" {
		va := r.FormValue("a")
		vb := r.FormValue("b")
		vc := r.FormValue("c")
		a, _ := strconv.ParseFloat(va, 64)
		b, _ := strconv.ParseFloat(vb, 64)
		c, _ := strconv.ParseFloat(vc, 64)

		fmt.Fprintf(w, "a = %v	b = %v	c = %v", a, b, c) // Выводим сообщение в браузер
		fmt.Fprintf(w, "\n\t<br /><br />")

		res := QuadraticEquation(a, b, c)
		//fmt.Println(res)
		if len(res) == 2 {
			fmt.Fprintf(w, "<h3>x1 = %v, x2 = %v</h3>", res[0], res[1])
		}
		if len(res) == 1 {
			fmt.Fprintf(w, "<h3>x = %v</h3>", res[0])
		}
		if len(res) == 0 {
			fmt.Fprintf(w, "<h3>D<0, рівняння немає дійсних рішень</h3>")
		}
	}
}
func main() {
	// Запускаем локальный сервер
	http.ListenAndServe("localhost:80", HttpSolution1)
}
