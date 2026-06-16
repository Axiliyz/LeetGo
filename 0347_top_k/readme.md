# 347 · Top K Frequent Elements

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-HashMap·Sorting-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n·log·n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given an integer array `nums` and an integer `k`, return the `k` most frequent elements. You may return the answer in any order.

Вернуть `k` самых часто встречающихся элементов массива.

**Ссылка:** [leetcode.com/problems/top-k-frequent-elements](https://leetcode.com/problems/top-k-frequent-elements/)

---

## Инженерка

### Что за задача

Нужно найти `k` элементов с наибольшей частотой появления в массиве. Порядок в ответе не важен.

```
[1, 1, 1, 2, 2, 3], k=2  →  [1, 2]
 ↑ встречается 3 раза
        ↑ встречается 2 раза
```

### Шаг 1 — подсчёт частот

Первым делом нужна карта частот: сколько раз встречается каждый элемент.

```
{1: 3, 2: 2, 3: 1}
```

### Шаг 2 — выбор top-K

Дальше варианты расходятся по сложности и элегантности.

---

### Подход 1 — сортировка (наше решение)

Преобразуем карту в слайс пар `{значение, частота}`, сортируем по частоте убыванием, берём первые `k`.

```
pairs: [{1,3}, {2,2}, {3,1}]
sort:  [{1,3}, {2,2}, {3,1}]  ← уже отсортированы по .v
top-2: [1, 2]
```

Читается чисто, работает за `O(n log n)`. Единственный минус — сортируем весь слайс, хотя нам нужны только `k` первых элементов.

---

### Подход 2 — Bucket Sort `O(n)`

Если нужна линейная сложность: частота не может быть больше `n`, поэтому можно создать «корзины» — массив из `n+1` слотов, где индекс = частота. Кладём числа в нужные корзины, потом идём с конца и собираем `k` элементов.

```
nums = [1,1,1,2,2,3], n=6

bucket[1] = [3]
bucket[2] = [2]
bucket[3] = [1]

← читаем с конца: bucket[3]→1, bucket[2]→2 → ответ [1, 2]
```

`O(n)` времени и памяти — оптимально, но чуть сложнее в реализации.

---

## Решение

```go
package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}
	pairs := makePairs(counts)
	sorted := sortPairs(pairs)
	return printKPairs(sorted, k)
}

type pair struct {
	k int
	v int
}

func makePairs(counts map[int]int) []pair {
	pairs := make([]pair, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, pair{k, v})
	}
	return pairs
}

func sortPairs(pairs []pair) []pair {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].v > pairs[j].v
	})
	return pairs
}

func printKPairs(pairs []pair, k int) []int {
	res := make([]int, 0, k)
	for i := 0; i < k; i++ {
		res = append(res, pairs[i].k)
	}
	return res
}
```

---

## Разбор на примерах

| Вход                        | Частоты              | Отсортировано          | k=2      |
| --------------------------- | -------------------- | ---------------------- | -------- |
| `[1,1,1,2,2,3]`             | `{1:3, 2:2, 3:1}`    | `[(1,3),(2,2),(3,1)]`  | `[1, 2]` |
| `[1]`                       | `{1:1}`              | `[(1,1)]`              | `[1]`    |
| `[1,2]`                     | `{1:1, 2:1}`         | `[(1,1),(2,1)]`        | `[1, 2]` |

---

## Сложность

|            | Сложность      | Объяснение                                                        |
| ---------- | -------------- | ----------------------------------------------------------------- |
| **Время**  | `O(n log n)`   | Подсчёт частот `O(n)` + сортировка пар `O(n log n)`              |
| **Память** | `O(n)`         | Карта частот + слайс пар, оба линейны от числа уникальных элементов |

---

## Сравнение подходов

| Подход              | Время        | Память | Когда выбирать                              |
| ------------------- | ------------ | ------ | ------------------------------------------- |
| Сортировка (наше)   | `O(n log n)` | `O(n)` | Простота, читаемость                        |
| Bucket Sort         | `O(n)`       | `O(n)` | Когда важна линейная сложность              |
| Heap (min-heap k)   | `O(n log k)` | `O(k)` | Когда `k` мало, а `n` огромно               |

---

## Что применял

- **Карта частот** — `map[int]int` считает вхождения за один проход; стандартный первый шаг в задачах на «самые частые»
- **Struct pair** — упаковали ключ и его частоту в одну структуру, чтобы не потерять связь при сортировке
- **`sort.Slice` с компаратором** — сортировка убыванием по полю `.v`; компаратор `pairs[i].v > pairs[j].v` переворачивает стандартный порядок
- **Предаллокация** — `make([]pair, 0, len(counts))` и `make([]int, 0, k)` сразу резервируют нужную ёмкость

---

## Темы

`Hash Map` `Sorting` `Bucket Sort` `Top-K` `Counting`

---
