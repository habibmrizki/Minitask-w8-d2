package minitasks1

import (
	"fmt"
	"io"
	"os"
)

func ReadFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("gagal membuka file: %w", err)
	}

	content, err := io.ReadAll(file)
	if err != nil {
		// Panic jika terjadi error saat membaca
		panic(fmt.Sprintf("error membaca file: %v", err))
	}
	file.Close()

	return content, nil
}

func RunReadFile(path string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered dari panic:", r)
			fmt.Println("Proses berlanjut...")
			// Tutup file jika ada yang terbuka
			if file, err := os.Open(path); err == nil {
				file.Close()
			}
		}
	}()

	content, err := ReadFile(path)
	if err != nil {
		// Menampilkan error jika gagal membuka file
		fmt.Println("Error:", err)
		return
	}

	// Menampilkan konten file ketika tidak ada error dan panic
	fmt.Println("==> File berhasil dibaca")
	fmt.Println(string(content))
}	