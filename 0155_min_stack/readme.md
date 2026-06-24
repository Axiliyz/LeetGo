# 155 · Min Stack

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Stack·Design-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(1)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Design a stack that supports `push`, `pop`, `top`, and retrieving the minimum element in constant time.

Нужно реализовать структуру данных со следующими операциями, каждая за `O(1)`:

1. `push(val)` — кладёт элемент на стек
2. `pop()` — снимает верхний элемент
3. `top()` — возвращает верхний элемент
4. `getMin()` — возвращает минимальный элемент во всём стеке

**Ссылка:** [leetcode.com/problems/min-stack](https://leetcode.com/problems/min-stack/)

---

## Инженерка

### Что за задача

Обычный стек хранит только верхушку, и узнать минимум можно только пройдя весь стек — `O(n)`. Задача — добавить к стеку "память" о минимуме, не теряя `O(1)` ни на одной операции.

### Первая мысль (наивное решение)

Хранить только `data` и при каждом `getMin()` пробегать по всему массиву в поиске минимума. Работает, но `getMin()` стоит `O(n)`, что не подходит под условие задачи.

Другой вариант — держать одну переменную `min` и обновлять её при `push`. Но при `pop()` минимума мы не знаем, каким он был до этого — переменная не хранит историю.

### Оптимальное решение — два стека в связке

Завести второй стек `mins`, который растёт и схлопывается **синхронно** с основным:

1. `push(val)` — кладём `val` в `data`. Если `mins` пуст или `val <= mins.top()` — кладём `val` и в `mins` тоже
2. `pop()` — снимаем верхушку `data`. Если она равна верхушке `mins` — снимаем и из `mins`
3. `top()` — просто верхушка `data`
4. `getMin()` — просто верхушка `mins`

`mins` хранит не «глобальный минимум один раз», а минимум **на каждый момент времени** — то есть минимум среди элементов, которые ещё лежат в стеке. Когда минимальный элемент снимается, `mins` откатывается к предыдущему минимуму, который был актуален до его появления.

### Почему именно так

Ключевая проблема наивного варианта с одной переменной — `pop()` не умеет "вспомнить" предыдущий минимум. Стек `mins` решает это явно: каждый раз, когда появляется новый минимум, старый не выбрасывается, а остаётся под ним в стеке и всплывёт обратно, когда новый минимум будет снят. По сути это история минимумов, а не одно число.

Условие `val <= mins.top()` (а не строгое `<`) важно при дублирующихся минимумах: если положить `1, 1` и затем `pop()` один раз, в `mins` должна остаться вторая единица, иначе `getMin()` после первого `pop` сломается.

---

## Решение

```go
package main

import (
	"cmp"
	"errors"
	"fmt"
)

var ErrEmptyStack = errors.New("Stack is empty")

type MinStack[T cmp.Ordered] struct {
	data []T
	mins []T
}

func NewMinStack[T cmp.Ordered]() *MinStack[T] {
	return &MinStack[T]{}
}

func (s *MinStack[T]) Push(val T) {
	s.data = append(s.data, val)

	if len(s.mins) == 0 || val <= s.mins[len(s.mins)-1] {
		s.mins = append(s.mins, val)
	}
}

func (s *MinStack[T]) Pop() error {
	if s.IsEmpty() {
		return ErrEmptyStack
	}

	topId := len(s.data) - 1
	val := s.data[topId]

	if val == s.mins[len(s.mins)-1] {
		s.mins = s.mins[:len(s.mins)-1]
	}

	var zero T
	s.data[topId] = zero
	s.data = s.data[:topId]
	return nil
}

func (s *MinStack[T]) Top() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, ErrEmptyStack
	}
	return s.data[len(s.data)-1], nil
}

func (s *MinStack[T]) GetMin() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, ErrEmptyStack
	}
	return s.mins[len(s.mins)-1], nil
}

func (s *MinStack[T]) IsEmpty() bool {
	return len(s.data) < 1
}

func main() {
	s := NewMinStack[int]()
	s.Push(8)
	s.Push(1)
	s.Push(2)
	fmt.Println(s.GetMin())
	fmt.Println(s.Pop())
	fmt.Println(s.IsEmpty())
	fmt.Println(s.Top())
}
```

---

## Разбор на примерах

| Операция   | `data`    | `mins`  | Результат |
| ---------- | --------- | ------- | --------- |
| `push(8)`  | `[8]`     | `[8]`   | —         |
| `push(1)`  | `[8,1]`   | `[8,1]` | —         |
| `push(2)`  | `[8,1,2]` | `[8,1]` | —         |
| `getMin()` | `[8,1,2]` | `[8,1]` | `1`       |
| `pop()`    | `[8,1]`   | `[8,1]` | снят `2`, в `mins` он не лежал — не трогаем |
| `getMin()` | `[8,1]`   | `[8,1]` | `1`       |
| `pop()`    | `[8]`     | `[8]`   | снят `1`, он был верхушкой `mins` — снимаем и там |
| `getMin()` | `[8]`     | `[8]`   | `8`       |

---

## Сложность

|            | Сложность | Объяснение                                                          |
| ---------- | --------- | -------------------------------------------------------------------- |
| **Время**  | `O(1)`    | Каждая операция работает только с верхушками двух стеков             |
| **Память** | `O(n)`    | В худшем случае (строго убывающая последовательность) `mins` растёт вместе с `data` |

---

## Что применял

- **Generics (`cmp.Ordered`)** — стек работает с любым сравнимым типом (`int`, `float64`, `string`), а не только с одним конкретным
- **Вспомогательный стек как история минимумов** — вместо одной переменной хранится стек минимумов, что позволяет «откатываться» к предыдущему минимуму при `pop()`
- **Нестрогое сравнение `<=` при push в `mins`** — критично для корректной работы с повторяющимися минимальными значениями
- **Явный `var zero T` перед обрезкой слайса** — обнуляет ссылку на удалённый элемент, чтобы избежать утечки памяти через GC при работе с указателями/строками
- **Sentinel-ошибка `ErrEmptyStack`** — единая переиспользуемая ошибка вместо `errors.New` в каждом методе

---

## Темы

`Stack` `Design` `Generics`

---
