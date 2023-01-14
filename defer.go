package main

import "fmt"

// Defer function adalah function yang bisa dijadwalkan untuk dieksekusi setelah function selesai dieksekusi,
// defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi

func logging() {
	fmt.Println("Selesai memanggil fungsi")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func main() {
	runApplication()
}
