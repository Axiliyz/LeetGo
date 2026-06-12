# 141 · Linked List Cycle

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-00d4aa?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-LinkedList·TwoPointers-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given `head`, the head of a linked list, determine if the linked list has a cycle in it. There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

Дан связный список. Нужно определить, есть ли в нём цикл — то есть существует ли узел, до которого можно вернуться, идя по `next`. Память — желательно `O(1)`.

**Ссылка:** [leetcode.com/problems/linked-list-cycle](https://leetcode.com/problems/linked-list-cycle/)

---

## Инженерка

### Как додумался до решения

Прогнал на примере `[3, 2, 0, -4]` с замыканием хвоста на голову.

Первая мысль — сложить адреса посещённых узлов в `map[*ListNode]bool`. Идём по списку: если узел уже видели — цикл; если дошли до `nil` — цикла нет. Работает за `O(n)` по времени, но тратит `O(n)` памяти.

Ключевая мысль: **если в списке есть цикл, то быстрый указатель рано или поздно догонит медленный изнутри петли.** Если цикла нет — быстрый просто упрётся в `nil`. Никакой памяти не нужно — только два указателя.

- Медленный движется на 1 шаг
- Быстрый — на 2 шага
- Встретились → цикл есть
- Быстрый дошёл до `nil` → цикла нет

Это алгоритм **Флойда (черепаха и заяц)**.

---

### Первая попытка — через мапу посещённых

Идём по списку и складываем указатели на узлы в множество. На каждом шаге проверяем, не встречали ли этот узел раньше.

```go
func hasCycle(head *ListNode) bool {
    seen := map[*ListNode]bool{}
    for cur := head; cur != nil; cur = cur.Next {
        if seen[cur] {
            return true
        }
        seen[cur] = true
    }
    return false
}
```

Работает за `O(n)` времени, но память `O(n)` — для больших списков жирно.

---

### Финальное решение — два указателя (Флойд)

Запускаем `slow` и `fast` от головы. `slow` шагает по одному узлу, `fast` — по два. Если в списке есть цикл, разрыв между указателями сокращается на 1 на каждой итерации внутри петли — значит они обязательно встретятся.

```
list: 3 → 2 → 0 → -4 ─┐
          ↑___________│

шаг 0:  slow=3,  fast=2
шаг 1:  slow=2,  fast=-4
шаг 2:  slow=0,  fast=2
шаг 3:  slow=-4, fast=-4   ← встретились → цикл
```

Если цикла нет — `fast` или `fast.Next` рано или поздно станут `nil`, и мы выйдем из цикла.

`O(n)` времени, `O(1)` памяти.

---

## Решение

```go
package main

import "fmt"

func (l *LinkedList[T]) hasCycle() bool {
	if l.Length < 2 {
		return false
	}
	slow := l.Head
	fast := l.Head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{3, 2, 0, -4}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	fmt.Println(l.hasCycle()) // false
	l.Tail.Next = l.Head
	fmt.Println(l.hasCycle()) // true
}
```

---

## Разбор на примерах

| Список                  | Структура                    | Результат |
| ----------------------- | ---------------------------- | --------- |
| `[3, 2, 0, -4]`         | хвост → голова (цикл)        | `true`    |
| `[1, 2]`                | хвост → голова (цикл)        | `true`    |
| `[3, 2, 0, -4]`         | обычный список               | `false`   |
| `[1]`                   | один узел, без цикла         | `false`   |
| `[]`                    | пустой список                | `false`   |

---

## Сложность

|            | Сложность | Объяснение                                                   |
| ---------- | --------- | ------------------------------------------------------------ |
| **Время**  | `O(n)`    | `fast` пройдёт максимум 2n узлов до встречи или до `nil`     |
| **Память** | `O(1)`    | Только два указателя `slow` и `fast`                         |

---

## Эволюция решения

| Подход                  | Время  | Память |
| ----------------------- | ------ | ------ |
| Мапа посещённых         | `O(n)` | `O(n)` |
| Два указателя (Флойд)   | `O(n)` | `O(1)` |

---

## Что применял

- **Алгоритм Флойда** — классический трюк «черепаха и заяц»: при наличии цикла быстрый указатель гарантированно догонит медленный
- **Два указателя** — заменяют хранение истории посещений и снижают память с `O(n)` до `O(1)`
- **Ранний выход по длине** — для списка из 0–1 узла цикл невозможен, проверяем сразу
- **Старт `fast = head.Next`** — даёт разные стартовые позиции, чтобы условие `slow == fast` не сработало на нулевом шаге

---

## Темы

`Linked List` `Two Pointers` `Hash Table`

---
