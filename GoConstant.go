package main

import "fmt"

//  Const lar bir defa yazılır ve değiştirlimez
/*
	ex:
	const a = "Merahba"
	fmt.Println(a) // Merhaba

	const a = "Selam" // Error daha önce tanımlı olduğu için değiştirilemez
*/

// CompanyName ile türeyen her şey stirn olmak zorunda
// const a CompanyName = "Selam" => Doğru
// const a CompanyName = 5 => Yanlış
// const a CompanyName = true => Yanlış

// CompanyName funclarda kullanılır
type CompanyName string

//  constlar () açarak çoklu kullanım olabilir
const (
	Facebook  CompanyName = "Facebook"
	Google    CompanyName = "Google"
	Apple     CompanyName = "Apple"
	Instagram CompanyName = "Instagram"
)

// Aşağıda nasıl kullanılacağını gördük
func PrintCompanyName(Company CompanyName) {
	fmt.Println(Company)
}

const thy = "turkish air lines"
const tai = "tai"

func main() {
	// Constant

	fmt.Println(thy, tai) // turkish air lines tai

	PrintCompanyName(Facebook)  // Facebook
	PrintCompanyName(Apple)     // Apple
	PrintCompanyName(Google)    // Google
	PrintCompanyName(Instagram) // Instagram

}
