// Bu dersimizde struct yapıları ile defer kullanımını işleyeceğiz.
package main

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi : ", p.productName)
}
func main() {
	p := product{productName: "Laptop", unitPrice: 5000}
	defer p.save()
	fmt.Println("İşlem başarili")

}
