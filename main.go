package main

import (
	"fmt"

	"github.com/habibmrizki/minitask-day2/internals/minitasks1"
	minitask2 "github.com/habibmrizki/minitask-day2/internals/minitasks2"
)

func main() {
	fmt.Println("asdsad")
 
     // Membuat objek Person menggunakan "constructor"
    person := minitask2.NewPerson("Habib Muhammad Rizki", "Cilacap", "0878199393738")

    // Menampilkan data awal 
    fmt.Println("Data Awal:")
    fmt.Println(person.Print())

    // Menampilkan pesan salam 
    fmt.Println(person.Greet())

    // Mengubah nama menggunakan method setter 
    person.SetName("Habib Boboho")

    // Menampilkan pesan salam lagi untuk tes 
    fmt.Println("\nSetelah nama diubah:")
    fmt.Println(person.Greet())

    // Menampilkan data setelah diubah
    fmt.Println("Data Setelah Diubah:")
    fmt.Println(person.Print())

    // person := minitask2.NewPerson("Habib Muhammad Rizki", "Cilacao", "0878199393738")
    // fmt.Println(person.GetName())
    // fmt.Println(person.GetAddress())
    // fmt.Println(person.GetPhone())

	validPath := "text.txt" 
	invalidPath := "file.txt"
	directoryPath := "internals" 

	// Test 1: Jalankan program dengan input path yang valid
	fmt.Println("--- Test 1: Membaca file yang valid ---")
	minitasks1.RunReadFile(validPath)
	fmt.Println()

	// Test 2: Jalankan program dengan input path yang invalid (test error)
	fmt.Println("--- Test 2: Membaca file yang tidak ada (error) ---")
	minitasks1.RunReadFile(invalidPath)
	fmt.Println()

	// Test 3: Jalankan program dengan input path yang mengarah ke direktori (test panic)
	fmt.Println("--- Test 3: Membaca direktori (panic) ---")
	minitasks1.RunReadFile(directoryPath)
	fmt.Println()

}

