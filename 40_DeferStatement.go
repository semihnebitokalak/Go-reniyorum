// Bu dersimizde defer kavramını kavramak için farklı bir örnek yazdık.

package main

import "fmt"

func A(sayi int32) string {
	defer DeferFunc()
	if sayi%2 == 0 {
		return "Çift sayidir."
	}
	if sayi%2 != 0 {
		return "Tek sayidir."
	}
	return "Belli degildir."

}

func Test() {
	sonuc := A(15)
	fmt.Println(sonuc)
}
func DeferFunc() {
	fmt.Println("Bitti.")
}

func main() {
	Test()
}
