package main

import (
	"fmt"
	"learn-go/core"
)

// Danh sách chức năng (số -> function & tên hiển thị)
var functions = map[int]struct {
	name string
	fn   func()
}{
	1: {"HelloWorld", core.HelloWorld},
	2: {"Variable", core.Variable},
	3: {"Operators", core.Operators},
	4: {"If Else", core.IfElse},
	5: {"Switch Case", core.SwitchCase},
	6: {"Loop", core.ForLoop},
}

func main() {
	for {
		// Hiển thị menu tự động với tên function
		fmt.Println("\nChọn lấy 1 cái:")
		for key, value := range functions {
			fmt.Printf("%d. %s\n", key, value.name)
		}
		fmt.Println("99. Thoát chương trình")
		fmt.Print("Nhập lựa chọn: ")

		var choice int
		fmt.Scanln(&choice)

		fmt.Println("")
		// Xử lý lựa chọn
		if choice == 99 {
			fmt.Println("Thoát chương trình...")
			return
		}

		if function, exists := functions[choice]; exists {
			function.fn() // Gọi function tương ứng
		} else {
			fmt.Println("Lựa chọn không hợp lệ! Hãy nhập số từ menu.")
		}
	}
}
