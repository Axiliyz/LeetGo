# 242 · Valid Anagram

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-String·HashMap-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(k)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given two strings `s` and `t`, return `true` if `t` is an anagram of `s`, and `false` otherwise.

Определить, является ли строка `t` анаграммой строки `s` — то есть состоит ли она из тех же символов в том же количестве, но в другом порядке.

**Ссылка:** [leetcode.com/problems/valid-anagram](https://leetcode.com/problems/valid-anagram/)

---

## Инженерка

### Что за задача

Две строки — анаграммы, если содержат одинаковый набор символов с одинаковыми частотами. `"anagram"` и `"nagaram"` — анаграммы. `"rat"` и `"car"` — нет. Порядок не важен, важен только состав.

### Первая идея — сортировка

Отсортировать обе строки и сравнить. Если символы совпадают по составу, после сортировки строки станут идентичными.

```go
func isAnagram(s, t string) bool {
    a, b := []rune(s), []rune(t)
    sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
    sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
    return string(a) == string(b)
}
```

Работает и читается просто, но сортировка — это `O(n log n)`. Для строк это оверхед: нам не нужен порядок, нужны только частоты.

---

### Оптимальное решение — мапа частот

Посчитать, сколько раз встречается каждый символ в каждой строке, и сравнить две карты частот. Если карты равны — это анаграмма.

```
"anagram":  {a:3, n:1, g:1, r:1, m:1}
"nagaram":  {n:1, a:3, g:1, r:1, m:1}   → равны → true

"rat":      {r:1, a:1, t:1}
"car":      {c:1, a:1, r:1}             → не равны → false
```

`O(n)` времени (один проход по каждой строке), `O(k)` памяти, где `k` — число уникальных символов.

### Почему `rune`, а не `byte`

Если использовать `map[byte]int`, многобайтные UTF-8 символы (кириллица, эмодзи) будут считаться по отдельным байтам, и логика сломается. `range` по строке в Go итерирует именно по рунам (Unicode code points), поэтому `map[rune]int` корректно работает с любыми алфавитами — например, `"апельсин"` и `"спаниель"`.

---

## Решение

```go
package main

import (
	"fmt"
	"maps"
)

func buildMap(s string) map[rune]int {
	res := make(map[rune]int, len(s))
	for _, r := range s {
		res[r]++
	}
	return res
}

func isAnagram(s string, t string) bool {
	return maps.Equal(buildMap(s), buildMap(t))
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))     // true
	fmt.Println(isAnagram("rat", "car"))             // false
	fmt.Println(isAnagram("апельсин", "спаниель"))   // true
}
```

`maps.Equal` (Go 1.21+) сравнивает две карты по ключам и значениям — заодно ловит случай разной длины строк, так как при лишних символах карты просто не совпадут.

---

## Разбор на примерах

| Вход                       | Карта `s`              | Карта `t`              | Результат |
| -------------------------- | ---------------------- | ---------------------- | --------- |
| `"anagram"`, `"nagaram"`   | `{a:3,n:1,g:1,r:1,m:1}`| `{a:3,n:1,g:1,r:1,m:1}`| `true`    |
| `"rat"`, `"car"`           | `{r:1,a:1,t:1}`        | `{c:1,a:1,r:1}`        | `false`   |
| `"ab"`, `"a"`              | `{a:1,b:1}`            | `{a:1}`                | `false`   |
| `"апельсин"`, `"спаниель"` | совпадают по рунам     | совпадают по рунам     | `true`    |

---

## Сложность

|            | Сложность | Объяснение                                                        |
| ---------- | --------- | ----------------------------------------------------------------- |
| **Время**  | `O(n)`    | Проход по каждой строке + сравнение карт, операции `map` в среднем `O(1)` |
| **Память** | `O(k)`    | Две карты, где `k` — число уникальных символов (ограничено алфавитом) |

---

## Сравнение подходов

| Подход            | Время        | Память | Плюсы                       | Минусы                         |
| ----------------- | ------------ | ------ | --------------------------- | ------------------------------ |
| Сортировка        | `O(n log n)` | `O(n)` | Просто и наглядно           | Лишний оверхед на сортировку   |
| Две карты (наше)  | `O(n)`       | `O(k)` | Чисто, `maps.Equal` лаконично | Две аллокации карт             |

---

## Что применял

- **Карта частот** — `map[rune]int` подсчитывает, сколько раз встречается каждый символ; классический приём для задач «совпадают ли составы»
- **`range` по рунам** — итерация по строке в Go идёт по Unicode code points, а не байтам, поэтому решение работает с любым алфавитом
- **`maps.Equal`** — стандартная функция (Go 1.21+) для сравнения карт целиком, избавляет от ручного цикла
- **Предварительное выделение памяти** — `make(map[rune]int, len(s))` даёт хинт по размеру и снижает число реаллокаций

---

## Темы

`String` `Hash Map` `Hash Table` `Counting` `Unicode`

---
