package main

import (
	"fmt"
)

type Respondent struct {
	fullName string
	age int
	gender string
	isSmoker bool
	cigarVariant []string
}


func main() {
	respondent := []Respondent{
		Respondent {
			fullName: "daffa",
			age: 26,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Gudang Garam Surya",
				"Sampoerna",
				"Pod",
			},
		},
		Respondent {
			fullName: "ari",
			age: 56,
			gender: "Laki-laki",
			isSmoker: true,
			cigarVariant: []string{
				"Gudang Garam Surya",
				"Sampoerna",
				"Donhill Hitam",
			},
		},
	}
	for i := 0; i < len(respondent); i++ {
		// fmt.Println("Responden", i+1)
		// fmt.Println("Full Name:", respondent[i].fullName)
		// fmt.Println("Age:", respondent[i].age)
		// fmt.Println("Gender:", respondent[i].gender)
		// fmt.Println("Is Respondent Smoker:", respondent[i].isSmoker)
		// fmt.Println("Cigar Variant:", respondent[i].cigarVariant)

	}
	var name string
	var age int
	var gender string
	var isSmoker bool
	var cigarVariant []string


	fmt.Print("Masukkan Nama Anda: ")
	fmt.Scanln(&name)
	fmt.Print("Masukkan Umur Anda: ")
	fmt.Scanln(&age)
	fmt.Print("Apa Jenis Kelamin anda: ")
	fmt.Scanln(&gender)
	fmt.Print("Apakah Anda Perokok?")
	fmt.Scanln(&isSmoker)
	fmt.Print("Rokok Apa Saja yang Pernah Anda Coba?")
	fmt.Scanln(&isSmoker)
	fmt.Printf("Nama :%s\n",name)
	fmt.Printf("Umur :%d\n",age)
	fmt.Printf("Jenis Kelamin :%s\n",gender)
	fmt.Printf("Anda Anda Perokok :%t\n",isSmoker)
	fmt.Printf("Anda Anda Perokok :%s\n",cigarVariant)
}