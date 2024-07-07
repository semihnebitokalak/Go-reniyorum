// Bundan önceki dizi derslerimizde tek boyutlu dizileri görmüştük bu dersimizde çok boyutlu dizileri göreceğiz.
// Eğer Matematikte matrisleri biliyorsanız anlamanız daha da kolay olacaktır.
package main

import "fmt"

func main() {
	// Tek boyutlu dizi tanımlaması gibi ama extra olarak bir tane daha kapalı parantez ekleyeceğiz.
	var sayilar [2][3]int
	sayilar[0][0] = 1
	sayilar[0][1] = 2
	sayilar[0][2] = 3
	sayilar[1][0] = 4
	sayilar[1][1] = 5
	sayilar[1][2] = 6
	fmt.Println(sayilar[1][1])
	// Bunu for döngüsü ile yapabiliriz.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(sayilar[i][j])
		}
	}
}
