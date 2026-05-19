# 121 · Best Time to Buy and Sell Stock

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Array·Greedy-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> You are given an array `prices` where `prices[i]` is the price of a given stock on the `iᵗʰ` day. You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock. Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return `0`.

Дан массив цен акции по дням. Нужно купить один раз и продать один раз позже — найти максимальную прибыль. Если прибыль невозможна — вернуть `0`.

**Ссылка:** [leetcode.com/problems/best-time-to-buy-and-sell-stock](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/)

---

## Инженерка

### Как додумался до решения

Прогнал на примере `[7, 1, 5, 3, 6, 4]` вручную.

Купить надо **дешевле**, продать **позже и дороже**. Брутфорс очевиден: перебрать все пары `(i, j)` где `i < j` и взять максимум `prices[j] - prices[i]`. Но это `O(n²)`.

Ключевой инсайт: **когда мы стоим на дне `j`, нас интересует только самый дешёвый день до него, не позже.** Значит можно идти слева направо и просто запоминать минимум на ходу.

- Если текущая цена выше минимума — считаем прибыль, обновляем максимум
- Если текущая цена ниже минимума — обновляем минимум (нашли лучший день для покупки)

Один проход — и задача решена.

---

### Первая попытка — брутфорс

Перебираем все пары: для каждого дня покупки пробуем все дни продажи после него.

```go
func maxProfit(prices []int) int {
    res := 0
    for i := 0; i < len(prices); i++ {
        for j := i + 1; j < len(prices); j++ {
            if prices[j]-prices[i] > res {
                res = prices[j] - prices[i]
            }
        }
    }
    return res
}
```

Работает, но `O(n²)` — на больших массивах TLE.

---

### Финальное решение — один проход с минимумом

Идём по массиву, держим два значения: `min` (лучший день покупки) и `res` (лучшая прибыль найденная пока).

```
prices = [7, 1, 5, 3, 6, 4]

день 0: цена=7, min=7, profit=0,  res=0
день 1: цена=1, min=1, profit=-6, res=0   ← обновили min
день 2: цена=5, min=1, profit=4,  res=4
день 3: цена=3, min=1, profit=2,  res=4
день 4: цена=6, min=1, profit=5,  res=5   ← лучшая прибыль
день 5: цена=4, min=1, profit=3,  res=5
```

`O(n)` времени, `O(1)` памяти.

---

## Решение

```go
package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0
	min := prices[0]
	for _, price := range prices {
		if price > min {
			res = max(res, price-min)
		} else {
			min = price
		}
	}
	return res
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{4, 3, 2, 1}))        // 0
}
```

---

## Разбор на примерах

| `prices`          | Купить | Продать | Прибыль |
| ----------------- | ------ | ------- | ------- |
| `[7,1,5,3,6,4]`  | день 1 (цена 1) | день 4 (цена 6) | `5` |
| `[1,2]`           | день 0 (цена 1) | день 1 (цена 2) | `1` |
| `[4,3,2,1]`       | нельзя          | нельзя          | `0` |
| `[2,4,1]`         | день 0 (цена 2) | день 1 (цена 4) | `2` |

---

## Сложность

|            | Сложность | Объяснение                                  |
| ---------- | --------- | ------------------------------------------- |
| **Время**  | `O(n)`    | Один проход по массиву                      |
| **Память** | `O(1)`    | Только две переменные `min` и `res`         |

---

## Эволюция решения

| Подход              | Время    | Память | Проблема                         |
| ------------------- | -------- | ------ | -------------------------------- |
| Брутфорс (все пары) | `O(n²)`  | `O(1)` | TLE на больших входных данных    |
| Один проход         | `O(n)`   | `O(1)` | ✓                                |

---

## Что применял

- **Greedy-инсайт** — не нужно помнить всю историю цен; достаточно знать минимум до текущего дня
- **Один проход** — совмещаем обновление минимума и подсчёт прибыли в одном цикле без дополнительных структур
- **Ранний выход через ветку** — вместо вычисления прибыли на каждом шаге обновляем `min` только когда цена падает ниже, иначе считаем profit; код читается как условие задачи

---

## Темы

`Array` `Dynamic Programming` `Greedy`

---
