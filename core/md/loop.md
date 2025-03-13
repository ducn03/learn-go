# Hướng Dẫn Vòng Lặp For Trong Golang

## 1️⃣ Giới Thiệu
Vòng lặp `for` trong Golang là cách duy nhất để lặp qua các phần tử hoặc thực hiện một hành động nhiều lần. Go không có `while` hoặc `do-while`, nhưng `for` có thể thay thế chúng.

## 2️⃣ Cấu Trúc Cơ Bản Của `for`
### 🔹 Vòng lặp `for` cơ bản (Lặp với biến đếm)
```go
for i := 1; i <= 5; i++ {
    fmt.Println("Lần lặp thứ", i)
}
```
✅ **Chạy từ `1` đến `5`, tăng `i` sau mỗi lần lặp.**

---

### 🔹 `for` như `while` (Lặp theo điều kiện)
```go
j := 1
for j <= 5 {
    fmt.Println("Giá trị của j:", j)
    j++
}
```
✅ **Giống `while`, lặp khi `j <= 5`.**

---

### 🔹 `for` vô hạn (Cần `break` để thoát)
```go
count := 0
for {
    if count >= 3 {
        break
    }
    fmt.Println("Vòng lặp vô hạn, count =", count)
    count++
}
```
✅ **Chạy liên tục cho đến khi `break` được gọi.**

---

### 🔹 Lặp qua mảng (`range` với `for`)
```go
numbers := []int{10, 20, 30, 40, 50}
for index, value := range numbers {
    fmt.Printf("Phần tử %d: %d\n", index, value)
}
```
✅ **Duyệt từng phần tử của slice/mảng với `range`.**

---

## 3️⃣ Tổng Kết
✔ **Dùng `for` như vòng lặp đếm (`for i := 0; i < n; i++`).**  
✔ **Dùng `for` như `while` (`for điều kiện {}`).**  
✔ **Dùng `for` vô hạn (`for {}`) với `break`.**  
✔ **Dùng `for range` để duyệt slice, map.**  

