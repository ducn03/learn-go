# Hướng Dẫn Mảng (Array) Trong Golang

## 1️⃣ Giới Thiệu
Mảng (`array`) trong Go là một tập hợp các phần tử có cùng kiểu dữ liệu và có kích thước cố định. Khác với `slice`, mảng không thể thay đổi kích thước sau khi khai báo.

---
## 2️⃣ Khai Báo Mảng Trong Golang
```go
var numbers [5]int  // Mảng có 5 phần tử kiểu int
numbers[0] = 10      // Gán giá trị vào mảng
numbers[1] = 20
```
Hoặc khai báo với giá trị ban đầu:
```go
fruits := [3]string{"Táo", "Cam", "Xoài"}
```
✅ **Mảng `numbers` có 5 phần tử kiểu `int`, `fruits` có 3 phần tử kiểu `string`.**

---
## 3️⃣ Truy Cập Phần Tử Mảng
```go
fmt.Println("Phần tử đầu tiên:", numbers[0])
fmt.Println("Phần tử thứ hai:", numbers[1])
```
✅ **Dùng `numbers[index]` để truy cập phần tử theo chỉ mục (`index`).**

---
## 4️⃣ Duyệt Mảng Bằng Vòng Lặp `for`
### 🔹 Cách 1: Dùng `for` với `len()`
```go
for i := 0; i < len(numbers); i++ {
    fmt.Println("Phần tử", i, ":", numbers[i])
}
```
### 🔹 Cách 2: Dùng `range`
```go
for index, value := range fruits {
    fmt.Printf("Phần tử %d: %s\n", index, value)
}
```
✅ **Dùng `range` giúp code ngắn gọn hơn khi duyệt mảng.**

---
## 5️⃣ So Sánh Mảng Trong Golang
```go
a := [3]int{1, 2, 3}
b := [3]int{1, 2, 3}
c := [3]int{4, 5, 6}
fmt.Println("a == b:", a == b) // true vì các phần tử giống nhau
fmt.Println("a == c:", a == c) // false vì giá trị khác nhau
```
✅ **Go cho phép so sánh mảng cùng kích thước bằng `==`.**

---
## 6️⃣ Tổng Kết
✔ **Mảng có kích thước cố định, không thể thay đổi sau khi khai báo.**  
✔ **Truy cập phần tử bằng `array[index]`.**  
✔ **Dùng `for` hoặc `range` để duyệt qua mảng.**  
✔ **Mảng có thể so sánh nếu có cùng kích thước.**  
