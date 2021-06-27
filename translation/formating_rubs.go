package main

import (
	"fmt"

	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	fmt.Println(transformToRub(1231231231231))
	fmt.Println(transformToRub(123123.123123))
	fmt.Println(transformToRub("12312312321"))

}

func transformToRub(total interface{}) string {
	tr := message.NewPrinter(language.Russian)

	return tr.Sprint(total)
}
