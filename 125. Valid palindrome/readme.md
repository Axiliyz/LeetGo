# 125 · Valid Palindrome

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-TwoPointers·String-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> A phrase is a **palindrome** if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers. Given a string `s`, return `true` if it is a palindrome, or `false` otherwise.

Строка — палиндром, если после удаления всех не-буквенно-цифровых символов и приведения к нижнему регистру она читается одинаково в обе стороны. Вернуть `true` или `false`.

**Ссылка:** [leetcode.com/problems/valid-palindrome](https://leetcode.com/problems/valid-palindrome/)

---

## Инженерка

### Как додумался до решения

Прогнал вручную на `"A man, a plan, a canal: Panama"`.

Первая мысль — убрать мусор заранее: пройтись по строке, оставить только буквы и цифры в нижнем регистре, получить чистую строку, а потом проверить её на палиндром двумя указателями. Это работает и читается просто.

```go
// prepareString оставляет только alphanumeric в нижнем регистре
func prepareString(s string) string {
    return strings.Map(func(r rune) rune {
        if unicode.IsLetter(r) || unicode.IsDigit(r) {
            return unicode.ToLower(r)
        }
        return -1
    }, s)
}
```

Но тут тратим `O(n)` памяти на новую строку. Можно не создавать её вообще: поставить два указателя по краям оригинальной строки и просто **перепрыгивать** не-alphanumeric символы прямо в цикле. Память `O(1)`, логика та же.

---

### Первая попытка — чистим строку, потом проверяем

Создаём отфильтрованную строку, затем сравниваем с двух концов.

```go
func isPalindrome(s string) bool {
    str := prepareString(s)
    left, right := 0, len(str)-1
    for left < right {
        if str[left] != str[right] {
            return false
        }
        left++
        right--
    }
    return true
}
```

Читаемо, но `O(n)` памяти на промежуточную строку.

---

### Финальное решение — два указателя без предобработки

Идём с двух концов и пропускаем «мусорные» символы на лету. Никакой новой строки — сравниваем прямо в исходной.

```
s = "A man, a plan, a canal: Panama"
      ^                           ^
      L                           R

пропускаем ' ', ',', ':', итд прямо внутри цикла
сравниваем ToLower(s[L]) == ToLower(s[R])
```

`O(n)` времени, `O(1)` памяти.

---

## Решение

```go
package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(str string) bool {
	left := 0
	right := len(str) - 1
	for left < right {
		for !unicode.IsLetter(rune(str[left])) && !unicode.IsDigit(rune(str[left])) && left < right {
			left++
		}
		for !unicode.IsLetter(rune(str[right])) && !unicode.IsDigit(rune(str[right])) && left < right {
			right--
		}
		if unicode.ToLower(rune(str[left])) != unicode.ToLower(rune(str[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama")) // true
	fmt.Println(isPalindrome("race a car"))                     // false
	fmt.Println(isPalindrome(" "))                              // true
}
```

---

## Разбор на примерах

| Строка                            | После очистки          | Палиндром? |
| --------------------------------- | ---------------------- | ---------- |
| `"A man, a plan, a canal: Panama"`| `amanaplanacanalpanama` | `true`     |
| `"race a car"`                    | `raceacar`             | `false`    |
| `" "`                             | `""`                   | `true`     |
| `"0P"`                            | `0p`                   | `false`    |

---

## Сложность

|            | Сложность | Объяснение                                                   |
| ---------- | --------- | ------------------------------------------------------------ |
| **Время**  | `O(n)`    | Каждый символ посещается не более одного раза                |
| **Память** | `O(1)`    | Только два указателя, новая строка не создаётся              |

---

## Эволюция решения

| Подход                         | Время  | Память | Проблема                          |
| ------------------------------ | ------ | ------ | --------------------------------- |
| Предобработка + два указателя  | `O(n)` | `O(n)` | Лишняя строка в памяти            |
| Два указателя in-place         | `O(n)` | `O(1)` | ✓                                 |

---

## Что применял

- **Два указателя** — классика для задач на палиндромы: левый и правый сходятся к центру, при несовпадении сразу `false`
- **Пропуск невалидных символов внутри цикла** — вложенные `for` с условием позволяют двигать указатель до следующего alphanumeric прямо на месте, без предварительного прохода
- **`unicode.IsLetter` / `unicode.IsDigit`** — Go-пакет `unicode` корректно обрабатывает не только ASCII, что делает решение устойчивым к любым входным данным
- **`unicode.ToLower`** — приведение к нижнему регистру при сравнении вместо мутации строки

---

## Темы

`Two Pointers` `String`

---
