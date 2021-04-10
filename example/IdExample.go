package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Veri alma
	UserName := bufio.NewReader(os.Stdin)
	LastName := bufio.NewReader(os.Stdin)
	Age := bufio.NewReader(os.Stdin)

	// \n bir alt satıra geçmeyi sağlar...
	fmt.Print("Your Name: ")
	UserNamestr, _ := UserName.ReadString('\n')

	fmt.Print("Your LastName: ")
	LastNamestr, _ := LastName.ReadString('\n')

	fmt.Print("Your Age: ")

	// Verileri aldık
	Agestr, _ := Age.ReadString('\n')

	// Verileri ekrana yazdık
	fmt.Println("Name ", UserNamestr)
	fmt.Println("LastName ", LastNamestr)
	fmt.Println("Age ", Agestr)

}
