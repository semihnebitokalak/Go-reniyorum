// Burada if elseif ve else kullanımını gördük...
package conditionals

import "fmt"

func Demo2() {
	var hesap float64 = 1000
	var cekilmekIstenen float64 = 1000
	if cekilmekIstenen > hesap {
		fmt.Println("Para yetersizdir.")
	} else if cekilmekIstenen == hesap {
		fmt.Println("Paraniz hazirlaniyor.")
		fmt.Println("Dikkat! Hesabinizda para kalmadi.")
	} else {
		fmt.Println("Paraniz hazirlaniyor.")
	}
}
