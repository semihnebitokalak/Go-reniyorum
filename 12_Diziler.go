// Bu dersimizde çok önemli bir konu olan dizileri anlamaya çalışalım...

package main

import "fmt"

func main() {
	var sayilar [5]int // Eğer bu şekilde sayılara değer vermezsek çıktı olarak sıfır alırız...
	fmt.Println(sayilar)
	sayilar[2] = 12
	fmt.Println(sayilar)
	fmt.Println(sayilar[2])
	// Bu kodların çıktısı sırasıyla 0 0 0 0 0
	// 								 0 0 12 0 0
	// 								 12
	// Sayılar 2 . ye değer atamıştık ama onu 3.çıktı olarak yazdı.
	// Bunun sebebi dizilerin indeks yani değer atamaları 0'dan başlar.
}
