### Вопрос

Что выведет данная программа и почему?

```go
func main() {
    slice := []string{"a", "a"}
    
    func(slice []string) {
        slice = append(slice, "a")
        slice[0] = "b"
        slice[1] = "b"
        fmt.Print(slice)
    }(slice)
    fmt.Print(slice)
}
```

### Ответ

Вывод:

```go
[b b a][a a]
```

Объяснение:

Этот случай похож на предыдущий, однако перед изменением данных мы в слайс добавили
данные, произошла переаллокация данных и slice в функции стал скрывать в себе уже
другой массов, поэтому изменений в исходной функции мы не увидели.

Это можно увидеть, если изменить код следующим образом:

```go
func main() {
	//slice := []string{"a", "a"}
	slice := make([]string, 2, 3)
	slice[0] = "a"
	slice[1] = "a"

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
```

Вывод:

```go
[b b a][b b]
```

Здесь переаллокации не было и append вернул структуру с тем же скрытым массивом
внутри, так как места было достаточно для добавления нового элемента (len = 2, cap = 3).