// Burada kafanız karışabilir o yüzden tekrar ederek anlamaya çalışalım...
// demo1.go sayfasındaki bilgileri main.go sayfasına yazdırmaya çalışıyoruz...
// Fonksiyonları daha sonra göreceğiz ama anlamaya çalışalım

package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya!"
	fmt.Println(metin)
	var kdv int = 10
	fmt.Println(kdv)
	fmt.Println(100 + (100 * 10 / 100))
	var kdv2 float32 = 0.1
	fmt.Println(kdv2)
	fmt.Println(100 + 100*kdv2)
	kdv3 := 25
	fmt.Println(kdv3)
	fmt.Printf("kdv3 değişkeninin veri tipi = %T", kdv3)
	fmt.Println()
	kdv3 = 12
	fmt.Println(kdv3)
}
