// Bu dersimizde for döngüsünü göreceğiz...
// for döngüsü tekrar etme verilen koşula kadar gidip istediğimiz işlemi yaptırma demektir...

package main

import "fmt"

func main() {
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	fmt.Println(metin)
	// Bu şekilde birçok kez ekrana yazdırabiliriz.
	// Bunu döngü ile tek bir işlemde yapabiliriz...
	// Şimdi for döngüsünün başlangıç değeri için i diye bir değişken tanımlıyoruz...
	i := 1
	// infinite loop
	for i <= 5 {
		fmt.Println(metin)
		// yukarıdaki şekilde çalıştırırsanız sonsuz döngüye girecektir çünkü i değeri sabit.
		i = i + 1
	}
	// for döngüsünün asıl kullanımı bu şekildedir.
	for i := 0; i < 4; i++ {
		fmt.Println(metin)
	}
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}
