package main

import "fmt"

func main() {
	var sehirler [5]string
	sehirler[0] = "İstanbul"
	sehirler[1] = "Ankara"
	sehirler[2] = "Sakarya"
	sehirler[3] = "İzmir"
	sehirler[4] = "Kayseri"
	fmt.Println(sehirler)
	for i := 0; i < 5; i++ {
		fmt.Println(sehirler[i])
	}
}
