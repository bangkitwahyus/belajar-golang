package main

import "fmt"

//Panic function adalah fungsi yang bisa kita gunakan untuk menghentikan program
//Panic function biasanya digunakan disaat program kita terjadi error
//Saat panic fungsi dipanggil, program akan berhanti, namun defer fungsi akan tetap dieksekusi

// Recover function adalah function yang dapat kita gunakan untuk menangkap data panic
// Dengan recover proses panic akan terhenti, sehingga proses akan tetap berjalan

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error Message : ", message)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(true)
	fmt.Println("Doddy")
}
