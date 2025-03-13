package core

import "fmt"

func HelloWorld() {
	// In ra và tự động xuống dòng
	fmt.Println("Hello World!")

	// So sánh các cách in ra màn hình
	fmt.Println("Println - Tự động xuống dòng")
	fmt.Printf("Printf - Không xuống dòng, format: %s %d %.2f\n", "Chuỗi", 100, 3.14)
	fmt.Print("Print - Không xuống dòng\n")

	// fmt.Sprintf() - Trả về chuỗi chứ không in ra màn hình
	message := fmt.Sprintf("Sprintf - Trả về chuỗi: %s %d %.2f", "Chuỗi", 100, 3.14)
	fmt.Println(message)
}
