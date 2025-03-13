package core

import "fmt"

/*
Operators - Các phép toán trong Golang
- Toán tử số học: +, -, *, /, %
- Toán tử so sánh: ==, !=, >, <, >=, <=
- Toán tử logic: &&, ||, !
- Toán tử gán: =, +=, -=, *=, /=, %=
- Toán tử bitwise (ít dùng trong cơ bản): &, |, ^, <<, >>
*/
func Operators() {
	a, b := 10, 3

	// Toán tử số học (Arithmetic Operators)
	fmt.Println("----- TOÁN TỬ SỐ HỌC -----")
	fmt.Println("a + b =", a+b) // Cộng
	fmt.Println("a - b =", a-b) // Trừ
	fmt.Println("a * b =", a*b) // Nhân
	fmt.Println("a / b =", a/b) // Chia lấy phần nguyên
	fmt.Println("a % b =", a%b) // Chia lấy phần dư

	// Toán tử so sánh (Comparison Operators)
	fmt.Println("\n----- TOÁN TỬ SO SÁNH -----")
	fmt.Println("a == b:", a == b) // So sánh bằng
	fmt.Println("a != b:", a != b) // So sánh khác
	fmt.Println("a > b:", a > b)   // Lớn hơn
	fmt.Println("a < b:", a < b)   // Nhỏ hơn
	fmt.Println("a >= b:", a >= b) // Lớn hơn hoặc bằng
	fmt.Println("a <= b:", a <= b) // Nhỏ hơn hoặc bằng

	// Toán tử logic (Logical Operators)
	fmt.Println("\n----- TOÁN TỬ LOGIC -----")
	fmt.Println("true && false:", true && false) // AND logic
	fmt.Println("true || false:", true || false) // OR logic
	fmt.Println("!true:", !true)                 // NOT logic

	// Toán tử gán (Assignment Operators)
	fmt.Println("\n----- TOÁN TỬ GÁN -----")
	x := 5
	fmt.Println("Giá trị ban đầu của x:", x)
	x += 2 // x = x + 2
	fmt.Println("x += 2 ->", x)
	x -= 1 // x = x - 1
	fmt.Println("x -= 1 ->", x)
	x *= 3 // x = x * 3
	fmt.Println("x *= 3 ->", x)
	x /= 2 // x = x / 2
	fmt.Println("x /= 2 ->", x)
	x %= 3 // x = x % 3
	fmt.Println("x %= 3 ->", x)
}
