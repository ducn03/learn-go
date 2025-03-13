# Hướng Dẫn Toán Tử Trong Golang

## 1️⃣ Giới Thiệu
Toán tử trong Go giúp thực hiện các phép tính toán và so sánh. Golang hỗ trợ nhiều loại toán tử:
- **Toán tử số học** (`+`, `-`, `*`, `/`, `%`)
- **Toán tử so sánh** (`==`, `!=`, `>`, `<`, `>=`, `<=`)
- **Toán tử logic** (`&&`, `||`, `!`)
- **Toán tử gán** (`=`, `+=`, `-=`, `*=`, `/=`, `%=`, `&=`, `|=`, `^=`)
- **Toán tử bitwise** (`&`, `|`, `^`, `<<`, `>>`)

## 2️⃣ Toán Tử Số Học (`+`, `-`, `*`, `/`, `%`)
Dùng để thực hiện các phép toán cơ bản:
```go
package main
import "fmt"

func main() {
    a, b := 10, 3
    fmt.Println("a + b =", a+b)
    fmt.Println("a - b =", a-b)
    fmt.Println("a * b =", a*b)
    fmt.Println("a / b =", a/b)  // Chia lấy phần nguyên
    fmt.Println("a % b =", a%b)  // Chia lấy phần dư
}
```
✅ **Output:**
```
a + b = 13
a - b = 7
a * b = 30
a / b = 3
a % b = 1
```

## 3️⃣ Toán Tử So Sánh (`==`, `!=`, `>`, `<`, `>=`, `<=`)
Dùng để so sánh hai giá trị:
```go
fmt.Println("a == b:", a == b) // Bằng
fmt.Println("a != b:", a != b) // Khác
fmt.Println("a > b:", a > b)   // Lớn hơn
fmt.Println("a < b:", a < b)   // Nhỏ hơn
fmt.Println("a >= b:", a >= b) // Lớn hơn hoặc bằng
fmt.Println("a <= b:", a <= b) // Nhỏ hơn hoặc bằng
```
✅ **Output:**
```
a == b: false
a != b: true
a > b: true
a < b: false
a >= b: true
a <= b: false
```

## 4️⃣ Toán Tử Logic (`&&`, `||`, `!`)
Dùng để kiểm tra điều kiện logic:
```go
fmt.Println("true && false:", true && false) // AND logic
fmt.Println("true || false:", true || false) // OR logic
fmt.Println("!true:", !true)                 // NOT logic
```
✅ **Output:**
```
true && false: false
true || false: true
!true: false
```

## 5️⃣ Toán Tử Gán (`=`, `+=`, `-=`, `*=`, `/=`, `%=`)
Dùng để gán và cập nhật giá trị biến:
```go
x := 5
x += 2  // x = x + 2
x -= 1  // x = x - 1
x *= 3  // x = x * 3
x /= 2  // x = x / 2
x %= 3  // x = x % 3
```
✅ **Output:**
```
x += 2 -> 7
x -= 1 -> 6
x *= 3 -> 18
x /= 2 -> 9
x %= 3 -> 0
```

## 6️⃣ Tổng Kết
✔ **Toán tử số học giúp thực hiện phép tính.**  
✔ **Toán tử so sánh dùng để kiểm tra giá trị.**  
✔ **Toán tử logic hỗ trợ điều kiện `if-else`.**  
✔ **Toán tử gán giúp cập nhật giá trị nhanh hơn.**  
