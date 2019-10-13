package main

import "fmt"

type node struct {
	value string
	priopity int
	next *node
}

func PrintQueue(p *node) {
	fmt.Println()
	for z := p; z != nil; z = z.next {
		f := fmt.Sprintf("Значение: %s, приоритет: %d", z.value, z.priopity)
		fmt.Println(f)
	}
}

func Add(value string, priopity int, l *node) *node {
	e := &node{value, priopity, nil}
	if l == nil {
		l = e
		return l;
	}
	for p := l; p != nil; p = p.next {
		if p.next != nil && priopity == p.next.priopity {
			e.next = p.next
			p.next = e
			return l
		}
		if priopity > p.priopity { // in start
			e.next = p
			return e
		}
		if p.next == nil {
			p.next = e
			return l
		}
		if priopity < p.priopity && priopity > p.next.priopity {
			e.next = p.next
			p.next = e
			return l
		}
	}
	return l
}

func NewQueue(p *node) *node {
	var value string
	fmt.Print("Введи значение: ")
	fmt.Scanln(&value)
	var priopity int
	fmt.Print("Введи приоритет: ")
	fmt.Scanln(&priopity)
	return Add(value, priopity, p)
}

func SrchQueue(p *node) {
	var k int
	var value string
	var priopity int
	i := 1
	for i > 0 {
		fmt.Print("\n9 - Назад\n1 - Приоритет и значене\n2 - Значение")
		fmt.Print("\n3 - Приоритет\n ^__^ Выберите тип поиска:")
		fmt.Scanln(&k)
		switch k {
		case 1:
			fmt.Print("Введите значение: ")
			fmt.Scanln(&value)
			fmt.Print("Введите приоритет: ")
			fmt.Scanln(&priopity)
			Search(p, value, priopity)
		case 2:
			fmt.Print("Введите значение: ")
			fmt.Scanln(&value)
			SearchByValue(p, value)
		case 3:
			fmt.Print("Введите приоритет: ")
			fmt.Scanln(&priopity)
			SearchByPriority(p, priopity)
		case 9: return;
		}
	}
}

func Search(l *node, value string, priopity int) {
	for p := l; p != nil; p = p.next {
		if p.value == value && p.priopity == priopity {
			f := fmt.Sprintf("Найденый элемент имеет значение(%s) и приоритет(%d)", p.value, p.priopity)
			fmt.Println(f)
			return
		}
	}
	fmt.Println("Элемента с таким значенем не существует")
}

func SearchByValue(l *node, value string) {
	for p := l; p != nil; p = p.next {
		if p.value == value {
			f := fmt.Sprintf("Найденый элемент имеет значение(%s) и приоритет(%d)", p.value, p.priopity)
			fmt.Println(f)
			return
		}
	}
	fmt.Println("Элемента с таким значенем не существует")
}

func SearchByPriority(l *node, priopity int) {
	for p := l; p != nil; p = p.next {
		if p.priopity == priopity {
			f := fmt.Sprintf("Найденый элемент имеет значение(%s) и приоритет(%d)", p.value, p.priopity)
			fmt.Println(f)
			return
		}
	}
	fmt.Println("Элемента с таким значенем не существует")
}

func DelQueue(p *node) *node {
	if p == nil {
		return nil
	}
	f := fmt.Sprintf("Удален элемент %s с приоритетом %d", p.value, p.priopity)
	fmt.Println(f)
	return p.next
}

func Menu() int {
	var k int
	fmt.Print("\n9 - Выход\n1 - Добавить\n2 - Удалить\n")
	fmt.Print("3 - Найти элемент\n4 - Вывести на экран\n")
	fmt.Print(" ~_~ Сделай выбор:")
	fmt.Scanln(&k)
	fmt.Println()
	return k
}

func main() {
	var p *node
	lime := 0
	for lime < 1 {
		switch Menu() {
		case 1:
			p = NewQueue(p)
		case 2:
			p = DelQueue(p)
		case 3:
			SrchQueue(p)
		case 4:
			PrintQueue(p)
		case 9: return;
		}
	}
}
