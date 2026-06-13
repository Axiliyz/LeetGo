# 049 · Group Anagrams

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-String·HashMap·Sorting-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n·k·log·k)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(n·k)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given an array of strings `strs`, group the anagrams together. You can return the answer in any order.

Сгруппировать строки-анаграммы вместе. `"eat"`, `"tea"`, `"ate"` — одна группа, потому что содержат одинаковый набор символов.

**Ссылка:** [leetcode.com/problems/group-anagrams](https://leetcode.com/problems/group-anagrams/)

---

## Инженерка

### Что за задача

Анаграммы — строки с одинаковым набором символов в одинаковых количествах, но в разном порядке. Нужно взять массив строк и разбить его на группы так, чтобы все анаграммы оказались вместе.

```
["eat","tea","tan","ate","nat","bat"]
→ [["eat","tea","ate"], ["tan","nat"], ["bat"]]
```

### Первая мысль — сравнивать попарно

Для каждой строки проверять все остальные — являются ли они анаграммами. Это `O(n² · k)`, где `k` — длина строки. На большом входе катастрофа.

### Оптимальное решение — ключ через сортировку

Главная идея: у всех анаграмм одной группы **отсортированные символы совпадают**.

```
"eat" → sort → "aet"
"tea" → sort → "aet"   ← тот же ключ
"ate" → sort → "aet"   ← тот же ключ

"tan" → sort → "ant"
"nat" → sort → "ant"   ← тот же ключ
```

Используем `map[string][]string`, где ключ — отсортированная строка, значение — группа анаграмм. Один проход по входу: для каждой строки вычисляем ключ и кладём строку в нужную группу.

### Почему `[]byte` а не `[]rune`

Задача на LeetCode гарантирует, что все символы — строчные латинские буквы (`a–z`). Это однобайтные ASCII символы, поэтому `[]byte` достаточно и быстрее, чем `[]rune`.

---

## Решение

```go
package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)

	for _, s := range strs {
		b := []byte(s)
		sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
		key := string(b)
		groups[key] = append(groups[key], s)
	}

	res := make([][]string, 0, len(groups))
	for _, v := range groups {
		res = append(res, v)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Println(groupAnagrams([]string{""}))
	fmt.Println(groupAnagrams([]string{"a"}))
}
```

---

## Разбор на примерах

| Строка | Ключ (sorted) | Группа в map               |
| ------ | ------------- | -------------------------- |
| `"eat"` | `"aet"`      | `["eat"]`                  |
| `"tea"` | `"aet"`      | `["eat", "tea"]`           |
| `"tan"` | `"ant"`      | `["tan"]`                  |
| `"ate"` | `"aet"`      | `["eat", "tea", "ate"]`    |
| `"nat"` | `"ant"`      | `["tan", "nat"]`           |
| `"bat"` | `"abt"`      | `["bat"]`                  |

Итог: `[["eat","tea","ate"], ["tan","nat"], ["bat"]]`

---

## Сложность

|            | Сложность        | Объяснение                                                                   |
| ---------- | ---------------- | ---------------------------------------------------------------------------- |
| **Время**  | `O(n · k log k)` | `n` строк, каждую сортируем за `O(k log k)`, где `k` — длина строки         |
| **Память** | `O(n · k)`       | Храним все строки в map + результирующий слайс                               |

---

## Что применял

- **Канонический ключ** — привести строку к «нормальной форме» (отсортировать символы), чтобы все анаграммы получили одинаковый ключ; классический приём для задач на группировку
- **`map[string][]string`** — ключ хранит сигнатуру группы, значение — все строки с этой сигнатурой; `append` к существующему слайсу в map работает через `groups[key] = append(groups[key], s)`
- **`[]byte` + `sort.Slice`** — конвертируем строку в байты, сортируем на месте, конвертируем обратно; дешевле чем `[]rune` для ASCII задач
- **`make([][]string, 0, len(groups))`** — предаллокируем результат по числу уникальных ключей, чтобы избежать лишних реаллокаций

---

## Темы

`String` `Hash Map` `Sorting` `Canonical Key` `Grouping`

---
