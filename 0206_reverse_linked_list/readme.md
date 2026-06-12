# 206 · Reverse Linked List

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-LinkedList·Recursion-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given the `head` of a singly linked list, reverse the list, and return the reversed list.

Дан односвязный список. Нужно развернуть его — последний узел становится головой, первый хвостом.
**Ссылка:** [leetcode.com/problems/reverse-linked-list](https://leetcode.com/problems/reverse-linked-list/)

---

## Инженерка

### Как додумался до решения

Прогнал на примере `[3, 2, 0, -4]`. Нужно получить `[-4, 0, 2, 3]`.

Первая мысль — собрать все значения в слайс, развернуть его и пересобрать список. Работает, но это `O(n)` по памяти и две лишних прогулки по списку.

Ключевая мысль: **разворачивать список можно на лету, переворачивая одну стрелку за раз.** Для этого достаточно трёх указателей — `prev`, `cur`, `next`. На каждом шаге запоминаем `next`, разворачиваем `cur.Next` назад на `prev`, и сдвигаем окно вперёд.

- `prev` — уже развёрнутая часть (изначально `nil`)
- `cur` — текущий узел, который разворачиваем
- `next` — буфер, чтобы не потерять хвост

Дополнительно — старый `Head` становится новым `Tail`, новый `Head` это последний `prev`.

---

### Решение — три указателя

Идём по списку один раз. На каждом шаге переворачиваем стрелку `cur.Next` назад на `prev`, предварительно сохранив `next`, чтобы не потерять остаток списка. В конце `prev` указывает на бывший хвост — это и есть новая голова.

```
было:  3 → 2 → 0 → -4 → nil

шаг 0:  prev=nil, cur=3
        nil ← 3   2 → 0 → -4 → nil

шаг 1:  prev=3,   cur=2
        nil ← 3 ← 2   0 → -4 → nil

шаг 2:  prev=2,   cur=0
        nil ← 3 ← 2 ← 0   -4 → nil

шаг 3:  prev=0,   cur=-4
        nil ← 3 ← 2 ← 0 ← -4

итог: Head=-4, Tail=3
```

`O(n)` времени, `O(1)` памяти — только три указателя.

---

## Решение

```go
func (l *LinkedList[T]) Reverse() error {
	if l.Length < 2 {
		return nil
	}

	var prev *ListNode[T]
	cur := l.Head
	l.Tail = l.Head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	l.Head = prev
	return nil
}

func main() {
	arr := []int{3, 2, 0, -4}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	l.Reverse()
	a := l.Head
	for a != nil {
		fmt.Println(a.Val) // -4, 0, 2, 3
		a = a.Next
	}
}
```

---

## Разбор на примерах

| Список            | После reverse     |
| ----------------- | ----------------- |
| `[3, 2, 0, -4]`   | `[-4, 0, 2, 3]`   |
| `[1, 2]`          | `[2, 1]`          |
| `[1]`             | `[1]`             |
| `[]`              | `[]`              |

---

## Сложность

|            | Сложность | Объяснение                                                   |
| ---------- | --------- | ------------------------------------------------------------ |
| **Время**  | `O(n)`    | Один проход по списку, разворачиваем по одной стрелке за шаг |
| **Память** | `O(1)`    | Три указателя `prev`, `cur`, `next` — не зависят от длины    |

---

## Что применял

- **Три указателя** — `prev`, `cur`, `next` позволяют развернуть список in-place за один проход
- **Сохранение `next` перед перезаписью** — без буфера `cur.Next = prev` потеряет связь с остатком списка
- **Обновление `Tail` до цикла** — бывший `Head` становится новым `Tail`, фиксируем заранее, пока он ещё доступен
- **Ранний выход по длине** — для 0–1 узла разворачивать нечего

---

## Темы

`Linked List` `Recursion` `Two Pointers`

---
