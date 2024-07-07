// Bu dersimizde klavyeden girilen bir sayının asal bir sayı olup olmadığını kontrol edeceğiz...

package main

import "fmt"

func main() {
	var sayi int
	fmt.Println("Lütfen bir sayi giriniz")
	fmt.Scanln(&sayi)
	for i := 0; i <= sayi/2; i++ {
		if sayi%2 == 0 || sayi%3 == 0 || sayi%5 == 0 || sayi%7 == 0 {
			fmt.Println("Girilen sayi bir asal sayi değildir.")
			break
		} else {
			fmt.Println("Girilen sayi bir asal sayidir.")
			break
		}

	}
}

// Yukarıdaki kodu şöyle düzenleyebiliriz. i değerini 2 den başlatıp if koşulunu da sayi % i == 0 yapabiliriz...
