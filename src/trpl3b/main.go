package main

import "fmt"

func main() {

	fmt.Println("Hello world")

	var mobil string = "BMW 2020"
	cars, motor := "Mercy 2020", "Ini motor honda"
	fmt.Println(mobil)
	fmt.Println(cars, " ", motor)

	//definisi variabel singkat
	huruf, angka := "satu", 1

	fmt.Println(huruf, "=", angka)

	//bolean
	var status bool
	status = true

	fmt.Println(status)
	//konstanta
	const nilai string = "123"
	fmt.Println("nilai")

	i := 1
	j := 5
	var hasil1 string
	if i < j {
		hasil1 = "i lebih kecil dari j"
	} else if i > j {
		hasil1 = "i lebih besar dari j"
	}

	fmt.Println(hasil1)

	var newCars [4]string
	newCars[0] = "BMW"
	newCars[1] = "Toyota"
	newCars[2] = "marcy"

	for i := 0; i <= 2; i++ {
		fmt.Println(newCars[i])
	}

}
