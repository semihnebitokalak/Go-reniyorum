// Bu dersimizde struct yapısını öğreneceğiz...
// Go dilinde struct, birden fazla veriyi tek bir veri yapısında saklamak için kullanılan kullanıcı tanımlı bir veri türüdür.
// struct ile farklı türdeki verileri bir araya getirip daha karmaşık veri yapıları oluşturabilirsiniz.
// struct lar, nesne tabanlı programlamadaki sınıfların (class) basit bir karşılığı olarak düşünülebilir.
// struct lar da fonksiyonlar gibi birer yapıdırlar.
package main

import "fmt"

// product(ürün) diye bir tane struct oluşturduk.
type product struct {
	name      string
	unitPrice float64
	// brand = marka
	brand string
}

func main() {
	// yazdırma işlemini böyle yazdırabiliyoruz.
	fmt.Println(product{"Bilgisayar", 40.999, "Macbook"})
	// Eğer isimli parametrelerle yazarsanız sırasıyla yazmanız gerekmez.
	fmt.Println(product{name: "Bilgisayar", brand: "Macbook", unitPrice: 55.499})
}
