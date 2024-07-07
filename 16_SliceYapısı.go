// Bu dersimizde Slice yapısını anlamaya çalışacağız...
package main

import "fmt"

func main() {
	// Slice yapısı nasıl oluşturulur önce ona bakalım.
	isimler := make([]string, 3)
	fmt.Println(isimler)
	// Bize string tipiinde 3 tane boş yer olurşturdu.
	isimler[0] = "Semih"
	isimler[1] = "Nebi"
	isimler[2] = "Tokalak"
	fmt.Println(isimler)
	fmt.Println(len(isimler))
	// Şimdi bu isimler arrayine yeni bir dizi ekleyelim.
	// Bu ekleme işlemini de append ile yapıyoruz.
	isimler = append(isimler, "Ahmet")
	fmt.Println(isimler)
	fmt.Println(len(isimler))
}
