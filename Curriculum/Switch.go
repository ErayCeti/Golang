package main

import "fmt"

func main() {

	// switch mantığı
	/*
		switch içine alınan değeri case ile değerlendirir

		örneğin:

		data := "Eray"
		switch data{
		case "Mehmet": // data içideki değer eğer mehmet ise bu girdi çalışacak
			fmt.Println("Mehmet")
		case "Yaren": // data içindeki değer eğer yaren ise bu girdi çalışacak
			fmt.Println("Yaren")
		default: // Eğer hiçbir case çalışmaz ise default olarak bu çalışacak
			fmt.Println("Bulamadım")
		}

		switch içinde de atama yapılabilir

		örneğin:
		switch  a := 1; a {
		case 1:
			fmt.Println("1")
		case 2:
			fmt.Println("2")
		case 3:
			fmt.Println("3")
		}

		switch caseleri içinde kontorl yapılabilir

		örneğin:
		switch a:=2; a {
		case a==1:
			fmt.Println("1")
		case a==2:
			fmt.Println("2")
		case a>2: // eğer sayı 2 den büyükse bu case çalışacak
			fmt.Println("2den çok büyük")
		}
	*/

	data := "Eray"

	switch data {
	case "Mehmet":
		fmt.Println("Girdiniz Mehmet")
	case "Yaren":
		fmt.Println("Yaren")
	default:
		fmt.Println("Bulamadım")
	}

	switch a := 1; a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("Bulamadım")
	}

}
