// Go diline girmiş bulunmaktayız. Sonunu getirme dileğiyle...
// Bu dersimizde basit bir ekrana yazma işlemi yapacağız ve birkaç bilgi vereceğiz...

package main

import "fmt"

func main() {
	fmt.Println("Hello")
	fmt.Println("World")
	fmt.Println("!")
	fmt.Print("Hello")
	fmt.Print("World")
	fmt.Print("!")
	fmt.Println()
	fmt.Print("My name is Semih Nebi Tokalak")
}

// Bu kodun çıktısından anlaşılacağı üzere Println ve Print arasındaki farkın sadece alt satıra inmesi olduğunu gördük.
// Ayrıca main içerisindeki yazdırma işlemlerini silip command + s yaparsak import ettiğimiz "fmt" de silinecektir.
// Çünkü Go dili kullanmayacağı şeyleri bünyesinde bulundurmaz.
