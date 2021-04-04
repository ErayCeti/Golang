package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Convert

	var MyString = "1"
	var x = 10
	var f float32 = 2.3

	fmt.Println(MyString, x, f) // 1 10 23

	// String den int
	number, _ := strconv.Atoi(MyString) // strconv.Atoi() func => içine string değer alır dışarı int değer döndürür, hatayı ikince paremtre alır
	fmt.Println(number)                 // 1 olarak geldi ama int type ile

	// eğerki var kullanmadıysan : kullan var niyete geçsin
	result := number + 2
	fmt.Println(result) // 3

	// int den string

	// var kuşşandığımız için : kullanmadık

	// go da text ile number toplanmaz
	/*
		var num = 5
		fmt.Println("Number: " + num) // Error
	*/
	var Mynumber = 5
	text := strconv.Itoa(Mynumber) // strconv.Itoa() func  => içine string değer alır int'e çevirir

	fmt.Println("Sonuc: " + text) // Sonuc: 5

	// convert casting

	var i int = 5
	var float float64 = float64(i + 1)
	var u uint = uint(float + 1)

	fmt.Println(i, float, u) // 5, 5, 5

}
