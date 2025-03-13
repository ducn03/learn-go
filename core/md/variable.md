# Hướng Dẫn Chi Tiết Về Biến Trong Golang

## 1️⃣ Giới Thiệu
Biến trong Golang được sử dụng để lưu trữ dữ liệu với nhiều loại kiểu khác nhau. Golang có các kiểu dữ liệu cơ bản như:
- **Số nguyên** (`int`, `uint`)
- **Số thực** (`float32`, `float64`)
- **Boolean** (`bool`)
- **Chuỗi** (`string`)
- **Mảng** (`array`)
- **Slice** (mảng động)
- **Map** (key-value)
- **Struct** (tương tự class trong OOP)
- **Interface** (định nghĩa hành vi chung)
- **Pointer** (tham chiếu đến địa chỉ bộ nhớ)
- **Hằng số** (`const`)

## 2️⃣ Cách Khai Báo Biến Trong Go
Golang hỗ trợ 3 cách khai báo biến:
- **Dùng `var`** (Có thể khai báo nhưng chưa cần gán giá trị).
- **Dùng `:=`** (Infer kiểu dữ liệu, chỉ dùng bên trong function).
- **Dùng `const`** (Hằng số, không thể thay đổi giá trị).

Ví dụ:
```go
var name string = "Golang"
age := 25
const PI = 3.14159
```

## 3️⃣ Các Kiểu Dữ Liệu Chính
### 🔹 3.1 Số Nguyên (`int`, `uint`)
- **`int`**: Có dấu, có thể chứa số âm và dương.
- **`uint`**: Không dấu, chỉ chứa số >= 0.

```go
var num int = -123
var unsignedNum uint = 456
numSecond := 1000 // Go tự động xác định kiểu dữ liệu
```

### 🔹 3.2 Số Thực (`float32`, `float64`)
- **`float32`**: Độ chính xác thấp hơn `float64`.
- **`float64`**: Chính xác cao hơn, thường được dùng mặc định.

```go
var pi float32 = 3.14
var price float64 = 99.99
```

### 🔹 3.3 Boolean (`bool`)
- Chỉ có **true** hoặc **false**.
```go
var isActive bool = true
```

### 🔹 3.4 Chuỗi (`string`)
- **Immutable** (Không thể thay đổi giá trị sau khi tạo).
```go
var str string = "Hello, Go!"
strSecond := "Golang is awesome!"
```

### 🔹 3.5 Mảng (`array`)
- **Có kích thước cố định, không thể thay đổi sau khi khai báo.**
```go
var arr [3]int = [3]int{1, 2, 3}
```

### 🔹 3.6 Slice (Mảng động)
- **Có thể thay đổi kích thước linh hoạt.**
```go
slice := []int{10, 20, 30}
```

### 🔹 3.7 Map (Key-Value, tương tự Dictionary trong Python)
- **Lưu trữ dữ liệu dạng cặp `key-value`.**
```go
user := map[string]int{"Alice": 25, "Bob": 30}
```

### 🔹 3.8 Struct (Tương tự class trong OOP)
- **Nhóm nhiều biến có liên quan vào một kiểu dữ liệu.**
```go
type Person struct {
    Name string
    Age  int
}
p := Person{Name: "Duc", Age: 25}
```

### 🔹 3.9 Interface (Định nghĩa hành vi chung)
- **Chỉ định những phương thức mà struct phải triển khai.**
```go
type Animal interface {
    Speak() string
}
```

### 🔹 3.10 Pointer (Tham chiếu đến địa chỉ bộ nhớ)
- **Lưu địa chỉ của một biến khác.**
```go
var x int = 10
var p *int = &x // Lấy địa chỉ của x
```

### 🔹 3.11 Hằng Số (`const`)
- **Không thể thay đổi giá trị sau khi khai báo.**
```go
const PI = 3.14159
```

## 4️⃣ Tổng Kết
✔ **Golang có hệ thống kiểu dữ liệu mạnh mẽ và linh hoạt.**  
✔ **Dùng `:=` để khai báo nhanh, `var` để khai báo biến toàn cục.**  
✔ **Dùng `const` để khai báo hằng số.**  
✔ **Mảng (`array`) cố định kích thước, còn `slice` linh hoạt hơn.**  
✔ **Dùng `map` khi cần lưu trữ dữ liệu dạng key-value.**  
✔ **Dùng `struct` để tạo kiểu dữ liệu tùy chỉnh.**  
✔ **Dùng `pointer` để tối ưu bộ nhớ khi làm việc với struct hoặc object lớn.**  

