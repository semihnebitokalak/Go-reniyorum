// Bu dersimizde önceki sorunun aynısı baz alıp kaç adımda doğru bildiğimizi yazacağız.
// Ayrıca 1. adımda doğru bildiysek = süper
// 2. adım ve 5. adım arasında bildiysek = oluru var
// 5. adım üstü bildiysek = gelişmen lazım
// yazılarını yazdırmamız gerekmektedir...
package main

import "fmt"

func main() {
	var akilda_olan_sayi int = 73
	var tahmin_edilen_sayi int = 0

	for counter := 1; counter < akilda_olan_sayi; counter++ {
		fmt.Println("Lütfen bir sayi tahmin ediniz.")
		fmt.Scanln(&tahmin_edilen_sayi)
		if tahmin_edilen_sayi > akilda_olan_sayi {
			fmt.Println("Lütfen daha düşük bir sayi giriniz.")
		} else if tahmin_edilen_sayi < akilda_olan_sayi {
			fmt.Println("Lütfen daha yüksek bir sayi giriniz.")
		} else {
			fmt.Println("Sayiyi doğru tahmin ettiniz.")
			fmt.Println("Tahmin sayiniz = " + fmt.Sprintf("%v", counter))
			if counter == 1 {
				fmt.Println("Süpersinnn")
			} else if counter < 6 {
				fmt.Println("Oluru var")
			} else {
				fmt.Println("Gelişmen lazimm")
			}
			break
		}

	}

}
