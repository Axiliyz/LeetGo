# 001 · Two Sum

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-Hash_Map·Array-6e40c9?style=flat-square&labelColor=0d1117)
![Time](https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117)
![Space](https://img.shields.io/badge/Space-O(n)-f0a500?style=flat-square&labelColor=0d1117)

</div>

---

## Задача

> Given an array of integers `nums` and an integer `target`, return indices of the two numbers such that they add up to `target`.

**Ссылка:** [leetcode.com/problems/two-sum](https://leetcode.com/problems/two-sum/)

---

## Инженерка

### Что за задача
Найти два числа в массиве, сумма которых равна `target`. Вернуть их индексы.

### Первая мысль (наивное решение)
Два вложенных цикла — перебираем все пары. Работает, но `O(n²)` — на большом массиве это катастрофа.

### Оптимальное решение — Hash Map
За один проход по массиву:
1. Для каждого элемента `v` вычисляем `complement = target - v`
2. Проверяем — есть ли `complement` уже в map?
3. Если да — нашли пару, возвращаем индексы
4. Если нет — кладём `v` в map и идём дальше

### Почему именно так
- `map[int]int` — ключ это значение элемента, значение это его индекс
- Lookup в map = `O(1)` в среднем
- Один проход = `O(n)` итоговая сложность

---

## Решение

```go
func twoSum(nums []int, target int) []int {
    mapa := make(map[int]int, len(nums))

    for i, v := range nums {
        if j, ok := mapa[target-v]; ok {
            return []int{j, i}
        }
        mapa[v] = i
    }

    return []int{0, 0}
}
```

---

## Сложность

| | Сложность | Объяснение |
|---|---|---|
| **Time** | `O(n)` | Один проход по массиву |
| **Space** | `O(n)` | В худшем случае все элементы в map |

---

## Что применял

- **Hash Map как инструмент lookup** — когда нужно быстро проверить "видел ли я уже X", map это первый выбор
- **Complement trick** — вместо поиска пары `(a, b)`, ищем `b = target - a`
- **`make(map[int]int, len(nums))`** — предаллокация capacity ускоряет работу, Go не будет перевыделять память
- **Запись после проверки** — сначала проверяем map, потом пишем. Иначе элемент найдёт сам себя

---

## Темы
`Hash Map` `Array` `One Pass` `Complement Pattern`

---
