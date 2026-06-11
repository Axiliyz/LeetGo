# 704 · Binary Search

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Binary_Search·Array-6e40c9?style=flat-square&labelColor=0d1117)
![Time](https://img.shields.io/badge/Time-O(log_n)-f0a500?style=flat-square&labelColor=0d1117)
![Space](https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117)

</div>

---

## Задача

> Given an array of integers `nums` which is sorted in ascending order, and an integer `target`, write a function to search `target` in `nums`. If `target` exists, then return its index. Otherwise, return `-1`.

Найти индекс элемента `target` в отсортированном массиве. Если элемента нет — вернуть `-1`.

**Ссылка:** [leetcode.com/problems/binary-search](https://leetcode.com/problems/binary-search/)

---

## Инженерка

### Что за задача
Классический бинарный поиск — фундаментальный алгоритм, который ищет элемент в отсортированном массиве за логарифмическое время. Массив уже отсортирован — это ключевое условие, без него бинарный поиск не работает.

### Первая мысль (наивное решение)
Линейный поиск — пробегаем массив от начала до конца, сравниваем каждый элемент с `target`. Работает за `O(n)`, но мы вообще не используем тот факт, что массив отсортирован.

### Оптимальное решение — Binary Search
Идея: на каждом шаге делим пространство поиска пополам.

1. Ставим два указателя: `left = 0`, `right = len(nums) - 1`
2. Считаем середину: `mid = (left + right) / 2`
3. Если `nums[mid] == target` — нашли, возвращаем `mid`
4. Если `nums[mid] < target` — target правее, двигаем `left = mid + 1`
5. Если `nums[mid] > target` — target левее, двигаем `right = mid - 1`
6. Повторяем, пока `left <= right`

### Почему именно так
- Каждая итерация отбрасывает половину оставшихся элементов
- За `log₂(n)` шагов пространство поиска сжимается до одного элемента
- Никакой дополнительной памяти — только три переменные (`left`, `right`, `mid`)

---

## Решение

```go
func Binary(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
```

---

## Разбор на примерах

**Массив:** `[-1, 0, 3, 5, 9, 12]`, **target:** `9`

| Шаг | `left` | `right` | `mid` | `nums[mid]` | Действие            |
| --- | ------ | ------- | ----- | ----------- | ------------------- |
| 1   | 0      | 5       | 2     | `3`         | `3 < 9` → left = 3 |
| 2   | 3      | 5       | 4     | `9`         | `9 == 9` → return 4 |

**Массив:** `[-1, 0, 3, 5, 9, 12]`, **target:** `2`

| Шаг | `left` | `right` | `mid` | `nums[mid]` | Действие              |
| --- | ------ | ------- | ----- | ----------- | --------------------- |
| 1   | 0      | 5       | 2     | `3`         | `3 > 2` → right = 1  |
| 2   | 0      | 1       | 0     | `-1`        | `-1 < 2` → left = 1  |
| 3   | 1      | 1       | 1     | `0`         | `0 < 2` → left = 2   |
| 4   | 2      | 1       | —     | —           | `left > right` → -1  |

---

## Сложность

|            | Сложность   | Объяснение                                    |
| ---------- | ----------- | --------------------------------------------- |
| **Время**  | `O(log n)`  | Каждый шаг делит пространство поиска пополам  |
| **Память** | `O(1)`      | Только три переменные, никаких доп. структур  |

---

## Что применял

- **Два указателя (left/right)** — классический паттерн для бинарного поиска, определяют текущее окно поиска
- **Деление пополам** — `mid = (left + right) / 2` делит пространство на две части; в Go нет проблемы переполнения для разумных размеров массива, но в продакшене безопаснее `left + (right - left) / 2`
- **Условие `left <= right`** — именно `<=`, а не `<`, иначе пропустим случай когда `left == right` и элемент находится ровно в этой единственной позиции
- **Сдвиг на `mid ± 1`** — важно не включать `mid` в новый диапазон, иначе цикл может зависнуть навсегда

---

## Темы
`Binary Search` `Array` `Two Pointers` `Divide and Conquer`

---
