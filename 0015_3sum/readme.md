# 015 · 3Sum

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Array·TwoPointers·Sorting-6e40c9?style=flat-square&labelColor=0d1117)
![Time](https://img.shields.io/badge/Time-O(n²)-f0a500?style=flat-square&labelColor=0d1117)
![Space](https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117)

</div>

---

## Задача

> Given an integer array `nums`, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`. Notice that the solution set must not contain duplicate triplets.

Найти все уникальные тройки чисел в массиве, сумма которых равна нулю.

**Ссылка:** [leetcode.com/problems/3sum](https://leetcode.com/problems/3sum/)

---

## Инженерка

### Что за задача

По сути — расширение Two Sum. Но есть два усложнения:
1. Нужно найти **все** тройки, а не одну пару
2. Тройки не должны **дублироваться** — `[-1, 0, 1]` и `[-1, 0, 1]` это одна и та же тройка

### Первая мысль (наивное решение)

Три вложенных цикла — перебрать все комбинации `(i, j, k)`. `O(n³)` — для `n = 3000` это ~2.7 × 10¹⁰ операций, не пройдёт по времени. Плюс нужно как-то убирать дубликаты — это ещё боль с set'ами.

### Оптимальное решение — сортировка + два указателя

Сводим задачу к Two Sum для отсортированного массива:

1. **Сортируем** массив — это позволит и пропускать дубликаты, и использовать два указателя
2. **Фиксируем** первый элемент `nums[i]` в цикле `i = 0..n-3`
3. Для оставшейся части ставим **два указателя**: `left = i+1`, `right = n-1`
4. Двигаем указатели навстречу:
   - `sum == 0` → нашли тройку, пропускаем дубликаты, сдвигаем оба указателя
   - `sum < 0` → нужно больше, двигаем `left` вправо
   - `sum > 0` → нужно меньше, двигаем `right` влево
5. **Пропуск дубликатов** на трёх уровнях:
   - Для `i` — если `nums[i] == nums[i-1]`, skip
   - Для `left` — после нахождения тройки, пропускаем одинаковые `nums[left]`
   - Для `right` — аналогично

### Почему именно сортировка

- Позволяет **детерминированно** пропускать дубликаты — одинаковые элементы стоят рядом
- Делает массив **монотонным** — два указателя навстречу гарантированно покрывают все пары
- Стоимость сортировки `O(n log n)` тонет в `O(n²)` основного алгоритма

---

## Решение

```go
func threeSum(nums []int) [][]int {
	res := make([][]int, 0, len(nums)/3)
	nums = QuickSort(nums)

	for i := range len(nums) - 2 {
		left := i + 1
		right := len(nums) - 1

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			summa := nums[i] + nums[left] + nums[right]
			if summa == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if summa < 0 {
				left++
			} else if summa > 0 {
				right--
			}
		}
	}
	return res
}
```

Сортировка — собственная реализация QuickSort (3-way partition):

```go
func QuickSort[T cmp.Ordered](arr []T) []T {
	if len(arr) < 2 {
		return arr
	}
	left := make([]T, 0, len(arr)/2)
	right := make([]T, 0, len(arr)/2)
	eq := make([]T, 0, len(arr)/2)
	pivot := arr[len(arr)/2]
	for i := 0; i < len(arr); i++ {
		if arr[i] < pivot {
			left = append(left, arr[i])
		} else if arr[i] > pivot {
			right = append(right, arr[i])
		} else {
			eq = append(eq, arr[i])
		}
	}
	left = QuickSort(left)
	right = QuickSort(right)
	res := append(left, eq...)
	res = append(res, right...)
	return res
}
```

---

## Разбор на примерах

| Вход | Результат |
| ---- | --------- |
| `[-1, 0, 1, 2, -1, -4]` | `[[-1, -1, 2], [-1, 0, 1]]` |
| `[0, 1, 1]` | `[]` |
| `[0, 0, 0]` | `[[0, 0, 0]]` |

Трассировка для `[-1, 0, 1, 2, -1, -4]` → после сортировки: `[-4, -1, -1, 0, 1, 2]`:

| `i` | `nums[i]` | `left` | `right` | `nums[l]` | `nums[r]` | `sum` | Действие |
| --- | --------- | ------ | ------- | --------- | --------- | ----- | -------- |
| 0 | -4 | 1 | 5 | -1 | 2 | -3 | `sum < 0` → left++ |
| 0 | -4 | 2 | 5 | -1 | 2 | -3 | `sum < 0` → left++ |
| 0 | -4 | 3 | 5 | 0 | 2 | -2 | `sum < 0` → left++ |
| 0 | -4 | 4 | 5 | 1 | 2 | -1 | `sum < 0` → left++ |
| 0 | -4 | 5 | 5 | — | — | — | `left == right` → конец |
| 1 | -1 | 2 | 5 | -1 | 2 | **0** | ✅ `[-1, -1, 2]` → skip dups, left++, right-- |
| 1 | -1 | 3 | 4 | 0 | 1 | **0** | ✅ `[-1, 0, 1]` → skip dups, left++, right-- |
| 1 | -1 | 4 | 3 | — | — | — | `left >= right` → конец |
| 2 | -1 | — | — | — | — | — | `nums[2] == nums[1]` → skip |
| 3 | 0 | 4 | 5 | 1 | 2 | 3 | `sum > 0` → right-- |
| 3 | 0 | 4 | 4 | — | — | — | `left == right` → конец |

Итого: `[[-1, -1, 2], [-1, 0, 1]]` ✅

---

## Сложность

|            | Сложность | Объяснение |
| ---------- | --------- | ---------- |
| **Время**  | `O(n²)` | Внешний цикл `O(n)` × внутренний два указателя `O(n)`. Сортировка `O(n log n)` тонет |
| **Память** | `O(n)` | QuickSort создаёт вспомогательные слайсы. Без учёта результата — `O(n)` |

---

## Что применял

- **Сведение к Two Sum** — фиксируем один элемент, ищем пару в оставшейся части. Классический приём декомпозиции
- **Два указателя навстречу** — работает на отсортированном массиве. `sum < 0` → двигаем левый, `sum > 0` → правый
- **Пропуск дубликатов на 3 уровнях** — для `i`, `left`, `right`. Сортировка делает дубликаты соседними → достаточно сравнения с предыдущим
- **`range len(nums) - 2`** — Go 1.22+ range over integer, заменяет `for i := 0; i < len(nums)-2; i++`
- **QuickSort с 3-way partition** — собственная реализация с дженериками `cmp.Ordered`. Разделение на `< pivot`, `== pivot`, `> pivot` эффективно для массивов с повторами
- **`make([][]int, 0, len(nums)/3)`** — предаллокация capacity для результата, чтобы избежать частых реаллокаций

---

## Темы
`Array` `Two Pointers` `Sorting` `QuickSort`

---
