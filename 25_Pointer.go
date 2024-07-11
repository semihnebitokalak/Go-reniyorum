// Bu dersimizde pointer kavramını derinlemesine inceleyeceğiz...
// Pointer kavramı kullanmak "*" kavramını kullanmak demektir.
package main

import "fmt"

func Arttir(sayi int) {
	sayi += 1
	fmt.Println("Arttir fonksiyonundaki sayi : ", sayi)
	// Biz yukarıda fonksiyonun içindeki sayıyı arttırdık. Fonksiyonun içinde bunu yazdırdığımızda sayı +1 artmış şekilde yazılacaktır.
}
func Arttir1(sayi *int) {
	*sayi = *sayi + 1
	// Yukarıdaki fonksiyondan farkı adres olan int ile değer olan int farklıdır.
	fmt.Println("Arttir fonksiyonundaki sayi : ", sayi)
	// Pointer kavramında adreslerle işlem yaptığımız için yukarıdaki yazdırma işleminde bellekte olduğu yeri gösterecektir.
	// Bunun çözümü için sayi kelimesinin başına * işareti konur.
	fmt.Println("Arttir fonksiyonundaki sayi : ", *sayi)
}

func main() {
	sayi := 20
	Arttir(sayi)
	// Arttir fonksiyonunu çağırdık ve 21 değerini ekranda gördük.
	// Ondan sonra mainde tekrar sayi değeri yazdırdık ve 20 değerini gördük. Bunun sebebi mainde tanımlı olan sayı ile
	// herhangi bir oynama(ekleme) yapmadık.
	// Arttir ve main fonksiyonlarındaki sayi değişkenler isim benzerliğine rağmen aynı değişkenler değildir.
	// Eğer sadece Arttir fonksiyonunu kullanarak değeri değiştirmek isteseydik sayi değerinin adresini referans göstermemiz gerekecekti.
	// Bunun için pointerları kullanıyoruz.
	fmt.Println("Maindeki sayi : ", sayi)
	// pointer ile aynı kodu tekrar yazalım.
	Arttir1(&sayi)
	fmt.Println("Maindeki sayi : ", sayi)

}
