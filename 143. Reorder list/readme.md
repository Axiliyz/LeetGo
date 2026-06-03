# 143 · Reorder List

<div align="left">

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-f0a500?style=flat-square&labelColor=0d1117)
![Language](https://img.shields.io/badge/Lang-Go-00ADD8?style=flat-square&logo=go&logoColor=white&labelColor=0d1117)
![Topics](https://img.shields.io/badge/Topics-LinkedList·TwoPointers-6e40c9?style=flat-square&labelColor=0d1117)
![Time](<https://img.shields.io/badge/Time-O(n)-f0a500?style=flat-square&labelColor=0d1117>)
![Space](<https://img.shields.io/badge/Space-O(1)-f0a500?style=flat-square&labelColor=0d1117>)

</div>

---

## Задача

> Given the `head` of a singly linked list `L0 → L1 → … → Ln-1 → Ln`, reorder it to:
> `L0 → Ln → L1 → Ln-1 → L2 → Ln-2 → …`
> You may not modify the values, only the nodes themselves.

Нужно «сложить» список пополам: первый узел, потом последний, потом второй, потом предпоследний и так далее, чередуя края к центру.
**Ссылка:** [leetcode.com/problems/reorder-list](https://leetcode.com/problems/reorder-list/)

---

## Инженерка

### Как додумался до решения

Прогнал на примере `[1, 3, 5, 7, 9]`. Нужно получить `[1, 9, 3, 7, 5]`.

Первая попытка — пройти список один раз и вставлять хвост между узлами. Провалилась: в **односвязном** списке нельзя дойти до предпоследнего элемента, не пройдя весь список заново. Указатель `last` стоит на месте, а нужно двигать его назад. В итоге рвутся связи и образуется цикл — `Print` зависает навсегда.

Классический способ для односвязного — три шага: найти середину (slow/fast), развернуть вторую половину, слить поочерёдно. Работает, но геморно.

Ключевая мысль: **если список двусвязный, reorder делается двумя указателями навстречу.** Один с головы (`left`), другой с хвоста (`right`). На каждом шаге пришиваем `right` сразу после `left`, потом сдвигаем `left` вперёд, а `right` — назад через `Prev`. Платим памятью за обратные ссылки, но логика читается на раз.

- `left` — идёт с головы вперёд
- `right` — идёт с хвоста назад (благодаря `Prev`)
- встретились или стали соседями → стоп

---

### Решение — два указателя навстречу

Сохраняем соседей (`leftNext`, `rightPrev`) **до** перезаписи стрелок, иначе потеряем остаток списка. Пришиваем `right` после `left`, двигаем окно к центру. Останавливаемся, когда указатели встретились (нечётная длина) или стали соседями (чётная).

```
было:  1 → 3 → 5 → 7 → 9

шаг 0:  left=1, right=9
        1 → 9   3 → 7 → 5(центр)
        leftNext=3, rightPrev=7

шаг 1:  left=3, right=7
        1 → 9 → 3 → 7   5(центр)
        left=5, right=5 → встретились, стоп

обрыв:  5.Next = nil

итог:  1 → 9 → 3 → 7 → 5
```

Условие цикла — `left != right && left.Next != right`:
- `left != right` ловит **нечётную** длину (встретились в центре).
- `left.Next != right` ловит **чётную** (стали соседями). Без неё получишь `right.Next = right` — самозацикливание.

После перестройки `Next` один финальный проход чинит все `Prev` и `Tail`, чтобы список остался валидным двусвязным.

`O(n)` времени, `O(1)` памяти.

---

## Решение

```go
func (l *LinkedList[T]) Reorder() error {
	if l.Length < 1 {
		return errors.New("List is empty")
	} else if l.Length < 3 {
		return nil
	}

	left, right := l.Head, l.Tail
	for left != right && left.Next != right {
		leftNext := left.Next
		rightPrev := right.Prev

		left.Next = right
		right.Next = leftNext
		left = leftNext
		right = rightPrev
	}
	left.Next = nil

	// финальный проход: чиним Prev у всех узлов и Tail
	var prev *ListNode[T]
	cur := l.Head
	for cur != nil {
		cur.Prev = prev
		prev = cur
		cur = cur.Next
	}
	l.Tail = prev

	return nil
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	l := NewLinkedList[int]()
	for _, v := range arr {
		l.Append(v)
	}
	l.Reorder()
	l.Print() // 1->9->3->7->5
}
```

> Узел двусвязный: `ListNode { Val, Next, Prev }`, а `Append` проставляет `node.Prev = l.Tail`.

---

## Разбор на примерах

| Список            | После reorder       |
| ----------------- | ------------------- |
| `[1, 3, 5, 7, 9]` | `[1, 9, 3, 7, 5]`   |
| `[1, 2, 3, 4]`    | `[1, 4, 2, 3]`      |
| `[1, 2, 3]`       | `[1, 3, 2]`         |
| `[1, 2]`          | `[1, 2]`            |
| `[1]`             | `[1]`               |

---

## Сложность

|            | Сложность | Объяснение                                                       |
| ---------- | --------- | ---------------------------------------------------------------- |
| **Время**  | `O(n)`    | Проход указателей навстречу + один проход на починку `Prev`      |
| **Память** | `O(1)`    | Два указателя, не зависят от длины (`Prev` уже хранится в узлах)  |

---

## Что применял

- **Два указателя навстречу** — `left` с головы, `right` с хвоста схлопываются к центру
- **Двусвязность (`Prev`)** — даёт шаг назад от хвоста за `O(1)`, чего нет в односвязном списке
- **Сохранение соседей перед перезаписью** — `leftNext`/`rightPrev` спасают остаток списка от потери
- **Двойное условие выхода** — `left != right` и `left.Next != right` корректно обрабатывают чётную и нечётную длину
- **Финальный проход на `Prev`/`Tail`** — после перестройки `Next` восстанавливаем целостность двусвязного списка

---

## Темы

`Linked List` `Two Pointers`

---
