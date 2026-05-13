# 020 · Valid Parentheses

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Stack·String-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given a string `s` containing just the characters `'('`, `')'`, `'{'`, `'}'`, `'['` and `']'`, determine if the input string is valid.

Строка валидна если:

1. Открывающие скобки закрываются скобкой **того же типа**
2. Открывающие скобки закрываются в **правильном порядке**
3. Каждая закрывающая скобка имеет **соответствующую открывающую**

**Ссылка:** [leetcode.com/problems/valid-parentheses](https://leetcode.com/problems/valid-parentheses/)

---

## Инженерка

### Что за задача

Проверить, правильно ли расставлены скобки трёх видов. Задача на вложенность — последняя открытая скобка должна закрыться первой.

### Первая мысль (наивное решение)

Считать количество открытых и закрытых скобок каждого типа. Не работает — `([)]` дало бы счётчики 1:1, но строка невалидна. Порядок имеет значение.

### Оптимальное решение — Stack

Стек как пачка прингл, последняя положенная чипсина достаётся первой(LIFO).

1. Встречаем открывающую `( [ {` — кладём на стек
2. Встречаем закрывающую `) ] }` — снимаем верхушку стека
3. Проверяем: верхушка — это нужная открывающая для текущей закрывающей?
4. Нет совпадения или стек пуст — `false`
5. Прошли всю строку — стек должен быть пуст

### Почему именно стек

Скобки — это рекурсивная вложенность. Стек моделирует её напрямую: открываем → запоминаем, закрываем → проверяем последнее запомненное. Аналог — стек вызовов функций: `main()` вызвал `foo()`, `foo()` вызвал `bar()` — `bar()` должен завершиться первым.

### Lookup Map вместо switch

`pairs` map связывает закрывающую скобку с ожидаемой открывающей. Это чище, чем цепочка `if/else` или `switch`, и легко расширяется.

---

## Решение

```go
package main

import (
	"errors"
	"fmt"
)

func isValid(s string) bool {
	stack := &Stack{}
	pairs := map[string]string {
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, ch := range s {
		char := string(ch)

		if char == "(" || char == "[" || char == "{" {
			stack.add(char)
		} else {
			top, err := stack.remove()
			if err != nil {
				return false
			}
			if pairs[char] != top {
				return false
			}
		}
	}

	return len(stack.data) == 0
}

type Stack struct {
	data []string
}

func (s *Stack) add(item string) {
	s.data = append(s.data, item)
}

func (s *Stack) remove() (string, error) {
	if len(s.data) < 1 {
		return "", errors.New("Wrong operation, size < 1")
	}
	elem := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return elem, nil
}

func (s *Stack) peek() (string, error) {
	if len(s.data) < 1 {
		return "", errors.New("Wrong operation, size < 1")
	}
	return s.data[len(s.data)-1], nil
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("([])"))
	fmt.Println(isValid("([)]"))
}

```

---

## Разбор на примерах

| Вход       | Трассировка стека                                         | Результат |
| ---------- | --------------------------------------------------------- | --------- |
| `"()"`     | push `(` → pop `(`, пара ✓ → стек пуст                    | `true`    |
| `"()[]{}"` | каждая пара закрывается сразу                             | `true`    |
| `"([])"`   | push `(` → push `[` → pop `[` ✓ → pop `(` ✓               | `true`    |
| `"(]"`     | push `(` → pop `(`, ожидали `[` → несовпадение            | `false`   |
| `"([)]"`   | push `(` → push `[` → pop `[`, ожидали `)` → несовпадение | `false`   |
| `"{"`      | push `{` → конец строки, стек не пуст                     | `false`   |
| `"}"`      | стек пуст при pop                                         | `false`   |

---

## Сложность

|            | Сложность | Объяснение                                                         |
| ---------- | --------- | ------------------------------------------------------------------ |
| **Время**  | `O(n)`    | Один проход по строке, каждый символ обрабатывается ровно один раз |
| **Память** | `O(n)`    | В худшем случае (все открывающие) весь стек заполнен               |

---

## Что применял

- **Stack (LIFO)** — ключевая структура для задач с вложенностью и парностью; последнее открытое должно закрыться первым
- **Lookup Map как диспетчер** — `pairs` заменяет громоздкий switch и делает логику декларативной: "закрывающая X соответствует открывающей Y"
- **Проверка пустого стека в конце** — `len(stack.data) == 0` ловит случай незакрытых скобок, когда строка прошла без ошибок, но открытых осталось больше
- **Ранний return false** — fail-fast подход: как только нашли невалидность, выходим, не тратим время на остаток строки
- **`peek()` как логическое завершение структуры стека** — метод добавлен в структуру для полноты самописной реализации

---

## Темы

`Stack` `String` `Bracket Matching` `Map`

---
