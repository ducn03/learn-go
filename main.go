package main

import (
	"fmt"
	"learn-go/core"
)

func main() {
	var choice int

	for {
		// Hiển thị menu lựa chọn
		fmt.Println("\nChọn một chức năng:")
		fmt.Println("1. Chạy HelloWorld")
		fmt.Println("2. Chạy Variable")
		fmt.Println("3. Thoát chương trình")
		fmt.Print("Nhập lựa chọn: ")

		fmt.Scanln(&choice)

		fmt.Println("")
		// Xử lý lựa chọn
		switch choice {
		case 1:
			core.Helloworld()
		case 2:
			core.Variable()
		case 3:
			fmt.Println("Thoát chương trình...")
			return // Dừng chương trình
		default:
			fmt.Println("Lựa chọn không hợp lệ! Hãy nhập 1, 2 hoặc 3, ...")
		}
	}
}
