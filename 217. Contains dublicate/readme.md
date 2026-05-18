# 217 · Contains Duplicate

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Array·HashSet-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given an integer array `nums`, return `true` if any value appears at least twice in the array, and return `false` if every element is distinct.

Определить, есть ли в массиве дубликаты.

**Ссылка:** [leetcode.com/problems/contains-duplicate](https://leetcode.com/problems/contains-duplicate/)

---

## Инженерка

### Что за задача

Простой вопрос: в массиве целых чисел есть повторяющиеся элементы или нет? Если да — `true`, если все уникальны — `false`.

### Первая идея — вложенные циклы (наивно)

Для каждого элемента пройти по всем остальным и проверить, есть ли дубликат.

```go
for i := 0; i < len(nums); i++ {
    for j := i + 1; j < len(nums); j++ {
        if nums[i] == nums[j] {
            return true
        }
    }
}
return false
```

Работает, но `O(n²)` — на большых массивах медленно. Для каждого элемента мы «лишние» проверяем всё, что находится справа.

---

### Вторая идея — сортировка

Отсортировать массив — тогда дубликаты окажутся рядом. Проход один, сравниваем соседние элементы.

```go
sort.Ints(nums)
for i := 0; i < len(nums) - 1; i++ {
    if nums[i] == nums[i+1] {
        return true
    }
}
return false
```

`O(n log n)` времени на сортировку, `O(1)` дополнительной памяти (в зависимости от реализации sort). Быстрее, но сортировка — это оверхед.

---

### Оптимальное решение — мапа

Идём по массиву один раз. Для каждого элемента:
- Он уже в мапе? → дубликат, `return true`
- Нет → добавляем

```
Seen: {}
[1,2,3,1]
↓ 1 → Seen: {1}
↓ 2 → Seen: {1,2}
↓ 3 → Seen: {1,2,3}
↓ 1 → уже в {1,2,3} → return true
```

`O(n)` времени, `O(n)` в худшем случае памяти (когда все элементы уникальны, весь массив в set). Это оптимально для этой задачи.

---

## Решение

```go
package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool, len(nums))
	for _, val := range nums {
		if _, ok := seen[val]; ok {
			return true
		}
		seen[val] = true
	}

	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1,2,3,1}))        // true
	fmt.Println(containsDuplicate([]int{1,2,3,4}))        // false
	fmt.Println(containsDuplicate([]int{1,1,3,1,3,4,3,2,4,2})) // true
}
```

---

## Разбор на примерах

| Вход                | Трассировка                  | Результат |
| ------------------- | ---------------------------- | --------- |
| `[1,2,3,1]`         | `{1}` → `{1,2}` → `{1,2,3}` → 1 уже есть | `true` |
| `[1,2,3,4]`         | `{1}` → `{1,2}` → `{1,2,3}` → `{1,2,3,4}` конец | `false` |
| `[99]`              | `{99}` конец                 | `false` |
| `[0,0]`             | `{0}` → 0 уже есть            | `true` |

---

## Сложность

|            | Сложность | Объяснение                                                 |
| ---------- | --------- | ---------------------------------------------------------- |
| **Время**  | `O(n)`    | Один проход по массиву, операции `map` (вставка, поиск) в среднем `O(1)` |
| **Память** | `O(n)`    | В худшем случае (нет дубликатов) весь массив в `map`    |

---

## Сравнение подходов

| Подход             | Время    | Память | Плюсы                  | Минусы                     |
| ------------------ | -------- | ------ | ---------------------- | -------------------------- |
| Вложенные циклы    | `O(n²)`  | `O(1)` | Без доп. памяти         | Медленно на больших массивах |
| Сортировка         | `O(n log n)` | `O(1)` | Не требует памяти       | Модифицирует массив        |
| HashSet (наше)     | `O(n)`   | `O(n)` | **Оптимально по времени** | Требует `O(n)` памяти      |

---

## Что применял

- **Хеш-таблица как набор** — в Go это `map[T]bool`, где значение `true` просто маркер того, что ключ есть в наборе; идиоматический способ сделать set без специального типа
- **Ранний return true** — fail-fast подход: как только нашли дубликат, останавливаемся; это делает среднее время лучше, чем `O(n)`
- **Предварительное выделение памяти** — `make(map[int]bool, len(nums))` говорит Go заранее выделить место для `n` элементов, избегаем ненужных реаллокаций
- **`_, ok := map[key]`** — идиоматический Go способ проверить наличие ключа без дополнительного кода

---

## Темы

`Array` `Hash Set` `Hash Table`

---
