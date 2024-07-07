// Bu dersimizde tanımlı olan bir integer değeri döngüler yardımı ile bilelim...
// İleride bunu kullanıcıdan alıp ondan sonra da bulabiliriz ama şimdi o sayısı sabit olucak...
package main

import "fmt"

func main() {
	var akilda_olan_sayi int = 73
	var tahmin_edilen_sayi int = 0

	for counter := 0; counter < akilda_olan_sayi; counter++ {
		fmt.Println("Lütfen bir sayi tahmin ediniz.")
		// Burada klavyeden bir değer isteniyor ve bunu da fmt.Scanln formatı ile yapacağız...
		// Fakat bunu kullanırken adres göstermeyi unutmayıp & işaretini kullanmalıyız...
		fmt.Scanln(&tahmin_edilen_sayi)
		// Klavyeden tahmin ettiğimiz sayıyı ekrana yazdırmak istersek eğer...
		//fmt.Println(tahmin_edilen_sayi)
		if tahmin_edilen_sayi > akilda_olan_sayi {
			fmt.Println("Lütfen daha düşük bir sayi giriniz.")
		} else if tahmin_edilen_sayi < akilda_olan_sayi {
			fmt.Println("Lütfen daha yüksek bir sayi giriniz.")
		} else {
			fmt.Println("Sayiyi doğru tahmin ettiniz.")
			break
		}

	}
	// Farklı çözümler yapıp bulmuş olabilirsiniz. Her problemin tek bir çözümü var diye bir şey yok...
}
