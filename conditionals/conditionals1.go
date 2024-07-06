// package conditionals: Bu, kodun conditionals adlı bir pakete ait olduğunu belirtir.
package conditionals

import "fmt"

// Demo1 adlı bir fonksiyon oluşturduk ve içini doldurduk.
func Demo1() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 900
	if cekilmekIstenen > hesap {
		fmt.Println("Para yetersizdir.")
	}
	if cekilmekIstenen < hesap {
		fmt.Println("Paraniz hazirlaniyor.")
	}
	hesap = hesap - cekilmekIstenen
	// printf ile bu şekilde yazabiliriz...
	fmt.Printf("Kalan bakiye = %v\n", hesap)
	// println ile yazmak istersek eğer;
	fmt.Println("Kalan bakiye = " + fmt.Sprintf("%v", hesap))
	// Bu yukarıdaki koddaki amaç metinsel bir kod ile sayı(value) bir değeri aynı anda yazmaktır.
	// Format dönüşümü için ise Sprintf i kullanırız...
}
