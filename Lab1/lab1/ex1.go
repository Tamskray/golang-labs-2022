/*Многострочный комментарий
Структура программы */
package main

//Однострочный комментарий
//Импорт пакетов
import "fmt"

func main() {
	//Объявление переменной
	var str string = "Golang!"

	//Вывод в консоль текста
	fmt.Println("Hello ", str)

	//Задание.
	//1. Вывести текст на украинском языке
	var strUkr string = "Привіт світ\nУкраїнська мова"
	fmt.Println(strUkr)

	strUkr1 := "Їжак"
	fmt.Println("-> " + strUkr1)
}
