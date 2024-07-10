// Aşağıdaki soruyu for range ile çözünüz...
// 1-10 a kadar olan tek sayıların toplamını yazdırın.
package main

import "fmt"

func main() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	toplam := 0
	for _, sayi := range sayilar {
		if sayi%2 != 0 {
			toplam += sayi
		}
	}
	fmt.Println("Tek sayilarin toplami : ", toplam)
	// Eğer 1 den 10 a kadar sayıların toplamını sorsaydı if karar yapısını kaldırmamız yeterli olacaktı.
}
