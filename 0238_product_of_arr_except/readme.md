# 238 · Product of Array Except Self

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Array·Prefix_Product-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n²)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given an integer array `nums`, return an array `answer` such that `answer[i]` is equal to the product of all the elements of `nums` except `nums[i]`. You must write an algorithm that runs in `O(n)` time and without using the division operation.

Для каждого индекса `i` вернуть произведение всех элементов массива, кроме `nums[i]`. Без деления, желательно за `O(n)`.

**Ссылка:** [leetcode.com/problems/product-of-array-except-self](https://leetcode.com/problems/product-of-array-except-self/)

---

## Инженерка

### Что за задача

```
nums   = [1, 2, 3, 4]
answer = [24, 12, 8, 6]

answer[0] = 2*3*4 = 24   (всё кроме nums[0])
answer[1] = 1*3*4 = 12   (всё кроме nums[1])
answer[2] = 1*2*4 = 8
answer[3] = 1*2*3 = 6
```

Деление запрещено условием — иначе можно было бы посчитать `total = произведение всех` и делить на `nums[i]`. Но это ещё и не сработало бы при наличии нулей в массиве (деление на ноль), так что условие защищает от хрупкого решения.

### Решение через слайсы (наше)

Для каждого `i` берём всё, что слева (`nums[:i]`) и всё, что справа (`nums[i+1:]`), перемножаем.

```go
prefix := nums[:i]   // всё до i
suffix := nums[i+1:] // всё после i
product := 1
// умножаем prefix и suffix друг на друга
```

Просто и наглядно, но для каждого `i` снова проходим почти весь массив — итого `O(n²)`. На входе в десятки тысяч элементов будет заметно медленно.

### Оптимальное решение — префикс × суффикс за один проход каждый

Идея: `answer[i] = (произведение всего слева от i) * (произведение всего справа от i)`. Эти произведения можно посчитать заранее за два линейных прохода, без пересчёта на каждой итерации.

```
nums    = [1, 2, 3, 4]

prefix[i] = произведение nums[0..i-1]:
prefix  = [1, 1, 2, 6]

suffix[i] = произведение nums[i+1..n-1]:
suffix  = [24, 12, 4, 1]

answer[i] = prefix[i] * suffix[i] = [24, 12, 8, 6]
```

```go
func productExceptSelf(nums []int) []int {
    n := len(nums)
    res := make([]int, n)

    prefix := 1
    for i := 0; i < n; i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    suffix := 1
    for i := n - 1; i >= 0; i-- {
        res[i] *= suffix
        suffix *= nums[i]
    }

    return res
}
```

Здесь `prefix` и `suffix` — не массивы, а просто бегущие переменные-аккумуляторы, которые накапливают произведение по ходу прохода. Память — только сам выходной массив, `O(1)` дополнительной.

---

## Решение

```go
package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
	fmt.Println(productExceptSelf([]int{-1, 1, 0, -3, 3}))
}

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		prefix := nums[:i]
		suffix := nums[i+1:]
		product := 1
		for _, num := range prefix {
			product *= num
		}
		for _, num := range suffix {
			product *= num
		}
		res[i] = product
	}
	return res
}
```

---

## Разбор на примерах

| `i` | `prefix`   | `suffix`   | Произведение | `res[i]` |
| --- | ---------- | ---------- | ------------ | -------- |
| 0   | `[]`       | `[2,3,4]`  | `1*2*3*4`    | `24`     |
| 1   | `[1]`      | `[3,4]`    | `1*3*4`      | `12`     |
| 2   | `[1,2]`    | `[4]`      | `1*2*4`      | `8`      |
| 3   | `[1,2,3]`  | `[]`       | `1*2*3`      | `6`      |

Случай с нулём: `[-1, 1, 0, -3, 3]` → `[0, 0, 9, 0, 0]` — везде где встречается ноль в произведении остальных, результат становится `0` (кроме позиции самого нуля).

---

## Сложность

|            | Сложность | Объяснение                                                          |
| ---------- | --------- | -------------------------------------------------------------------- |
| **Время**  | `O(n²)`   | Для каждого из `n` индексов проходим оставшиеся `~n` элементов      |
| **Память** | `O(n)`    | Слайсы `prefix`/`suffix` — это переиспользование исходного массива (без копий), плюс выходной `res` |

---

## Сравнение подходов

| Подход                  | Время    | Память | Плюсы                          | Минусы                              |
| ------------------------ | -------- | ------ | ------------------------------- | ------------------------------------ |
| Слайсы (наше)            | `O(n²)`  | `O(n)` | Очень наглядно, легко объяснить | Медленно на больших входах           |
| Префикс × суффикс         | `O(n)`   | `O(1)*` | Линейная сложность, оптимально  | Чуть менее очевидна на первый взгляд |

*без учёта самого выходного массива, который требуется по условию задачи

---

## Что применял

- **Слайсинг без копирования** — `nums[:i]` и `nums[i+1:]` в Go не копируют данные, а создают новый слайс-«вид» на тот же underlying массив; дёшево по памяти, но не ускоряет сам перебор
- **Отказ от деления** — задача явно запрещает делить на `nums[i]`, что подталкивает к разложению на префиксное и суффиксное произведение
- **Аккумулятор вместо массива** (в оптимальном варианте) — `prefix`/`suffix` можно держать как одну переменную, а не отдельный массив, экономя память

---

## Темы

`Array` `Prefix Product` `Suffix Product` `No Division`

---
