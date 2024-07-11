/*
Pointer Nedir? Pointerlar, bir değişkenin bellekteki adresini saklayan değişkenlerdir.
Pointer Tanımlama: var p *int ifadesi bir int pointer tanımlar.
Adres Alma: & operatörü kullanılarak bir değişkenin adresi alınır.
Değere Erişme: * operatörü kullanılarak bir pointerın gösterdiği değere erişilir veya bu değer değiştirilir.
Fonksiyonlarda Kullanım: Fonksiyonlara pointerlar geçirilerek, fonksiyonun argüman olarak aldığı değerleri doğrudan değiştirmesi sağlanabilir.
Bellek Ayırma: new fonksiyonu kullanılarak yeni bellek alanı ayrılabilir ve bu alanın adresi bir pointerda saklanabilir.

Array ve Pointer Tanımlama: Bir dizinin belirli elemanının adresini almak için & operatörü kullanılır.
Pointer Kullanarak Değiştirme: Bir dizinin elemanını pointer kullanarak değiştirmek mümkündür.
Fonksiyonlara Geçirme: Bir diziyi pointer olarak bir fonksiyona geçirmek, fonksiyonun diziyi doğrudan değiştirmesine olanak tanır.
Pointer Arasında Geçiş: Pointer aritmetiği kullanarak dizinin belirli bir elemanından diğer elemanlara erişmek mümkündür.
*/
// Aşağıda new fonksiyonu kullanılarak yapılmış bir örnek mevcuttur.
package main

import "fmt"

func main() {
	// new kullanarak int türünde yeni bir bellek alanı ayırma
	p := new(int)

	// Pointerın gösterdiği bellek alanına değer atama
	*p = 25
	fmt.Println("p pointerinin gösterdiği değer:", *p)
}
