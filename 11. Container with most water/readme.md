# 011 · Container With Most Water

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Array·TwoPointers·Greedy-6e40c9?style=flat-square&labelColor=0d1117)
![Time](https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117)
![Space](https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117)

</div>

---

## Задача

> You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `i`th line are `(i, 0)` and `(i, height[i])`. Find two lines that together with the x-axis forms a container, such that the container contains the most water.

Найти две вертикальные линии, которые вместе с осью X образуют контейнер с максимальным объёмом воды.

**Ссылка:** [leetcode.com/problems/container-with-most-water](https://leetcode.com/problems/container-with-most-water/)

---

## Инженерка

### Что за задача

Массив высот — это вертикальные стенки. Нужно выбрать две стенки так, чтобы площадь воды между ними была максимальной. Площадь = `ширина × min(левая, правая)`. Вода ограничена **короткой** стенкой — через длинную не перельётся.

### Первая мысль (наивное решение)

Два вложенных цикла — перебрать все пары `(i, j)`, посчитать площадь каждой. `O(n²)` — для `n = 10⁵` это ~10¹⁰ операций, не пройдёт.

### Оптимальное решение — два указателя навстречу

Начинаем с максимальной ширины: `left = 0`, `right = len - 1`. На каждом шаге:

1. Считаем площадь: `(right - left) * min(height[left], height[right])`
2. Обновляем максимум
3. **Двигаем указатель с меньшей высотой** внутрь

### Почему двигаем именно короткую стенку

Площадь = `ширина × min(высот)`. При сужении ширина **гарантированно** падает на 1. Единственный шанс увеличить площадь — поднять `min(высот)`.

- Если двигаем **короткую** стенку → `min` может вырасти → площадь может увеличиться
- Если двигаем **длинную** стенку → `min` останется прежним или упадёт → площадь **точно** не вырастет

Поэтому двигать длинную бессмысленно — мы гарантированно не пропустим оптимум.

---

## Решение

```go
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		area := (right - left) * min(height[left], height[right])
		maxArea = max(area, maxArea)
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
```

---

## Разбор на примерах

| Вход | Пара | Результат |
| ---- | ---- | --------- |
| `[1,8,6,2,5,4,8,3,7]` | `(8,7)` idx 1,8 | `49` |
| `[1,1]` | `(1,1)` idx 0,1 | `1` |

Трассировка для `[1, 8, 6, 2, 5, 4, 8, 3, 7]`:

| Шаг | `left` | `right` | `h[l]` | `h[r]` | Ширина | Площадь | `maxArea` | Действие |
| --- | ------ | ------- | ------ | ------ | ------ | ------- | --------- | -------- |
| 0 | 0 | 8 | 1 | 7 | 8 | 8 | 8 | `h[l] < h[r]` → left++ |
| 1 | 1 | 8 | 8 | 7 | 7 | **49** | **49** | `h[l] > h[r]` → right-- |
| 2 | 1 | 7 | 8 | 3 | 6 | 18 | 49 | `h[l] > h[r]` → right-- |
| 3 | 1 | 6 | 8 | 8 | 5 | 40 | 49 | `h[l] == h[r]` → left++ |
| 4 | 2 | 6 | 6 | 8 | 4 | 24 | 49 | `h[l] < h[r]` → left++ |
| 5 | 3 | 6 | 2 | 8 | 3 | 6 | 49 | `h[l] < h[r]` → left++ |
| 6 | 4 | 6 | 5 | 8 | 2 | 10 | 49 | `h[l] < h[r]` → left++ |
| 7 | 5 | 6 | 4 | 8 | 1 | 4 | 49 | `h[l] < h[r]` → left++ |
| — | 6 | 6 | — | — | — | — | — | `left == right` → стоп |

---

## Сложность

|            | Сложность | Объяснение |
| ---------- | --------- | ---------- |
| **Время**  | `O(n)` | Один проход двумя указателями навстречу |
| **Память** | `O(1)` | Три переменные: `left`, `right`, `maxArea` |

---

## Что применял

- **Два указателя** — начинаем с краёв, сужаем к центру. Работает потому что ширина монотонно убывает
- **Жадный выбор** — двигаем короткую стенку, потому что только так есть шанс увеличить `min(высот)`. Двигать длинную — гарантированно бесполезно
- **`left < right` вместо `left <= right`** — при `left == right` ширина = 0, площадь = 0 — лишняя итерация
- **Встроенные `min` / `max`** — Go 1.21+, без вспомогательных функций

---

## Темы
`Array` `Two Pointers` `Greedy`

---
