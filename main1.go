// http://algolist.manual.ru/ds/s_has.php
package main

import "fmt"

type link struct {
	value string
	next  *link
}

var dict [5]*link

// подсчет хеша
func Hash(v string) int32 {
	length := len(v)
	str := []rune(v)
	var sum int32 = 0
	for i := 0; i < length; i++ {
		sum += str[i]
	}
	var res int32 = sum % 5
	fmt.Println("Value: ", v, " | Hash: ", res, " | ✓ OK")
	return int32(res)
}

// выовд меню
func Menu() int {
	var k int
	fmt.Print("\n1 - Добавить\n2 - Удалить\n3 - Найти элемент\n")
	fmt.Print("4 - Вывести хеш таблицу\n9 - Выход\n\n  ~_~  Сделай выбор: ")
	fmt.Scanln(&k)
	fmt.Println()
	return k
}

// создание значения
func Add() {
	var value string
	fmt.Print("Введи элемент: ")
	fmt.Scanln(&value)
	var z int32 = Hash(value)
	dict[z] = Append(&link{value, nil}, dict[z])
}

// удаление значения
func Del() {
	var value string
	fmt.Print("Введи элемент для удаления: ")
	fmt.Scanln(&value)
	var z int32 = Hash(value)
	if Srch(value, dict[z]) {
		fmt.Println("Элемент удален: ✓ OK")
		dict[z] = Rm(value, dict[z])
	} else {
		fmt.Println("Ничего не найдено")
	}
}

// поиск данных
func Search() {
	var value string
	fmt.Print("Введи элемент который вы ищете: ")
	fmt.Scanln(&value)
	var z int32 = Hash(value)
	if Srch(value, dict[z]) {
		fmt.Println("Элемент найден")
	} else {
		fmt.Println("Ничего не найдено")
	}
}

// отображение хеш таблицы
func Show() {
	var empty bool = false
	for i := 0; i < 5; i++ {
		if dict[i] != nil {
			empty = true
			fmt.Print("dict #", i, "  /  -->")
			var cc int = 0
			for p := dict[i]; p != nil; p = p.next {
				cc++
				fmt.Print("  #", cc, " data: ", p)
			}
			fmt.Println()
		}
	}
	if !empty {
		fmt.Println("Хеш таблица пуста")
	}
}

// функция создания или добавления нового элемента списка
func Append(node, l *link) *link {
	if l == nil {
		return node
	}
	for i := l; i != nil; i = i.next {
		if i.next == nil {
			i.next = node
			return l
		}
	}
	return l
}

// функция поиска данных в списке
func Srch(v string, l *link) bool {
	if l != nil {
		for p := l; p != nil; p = p.next {
			if p.value == v {
				return true
			}
		}
	}
	return false
}

// функция верного удаления данных
func Rm(value string, l *link) *link {
	var r *link
	if l.next == nil {
		return nil
	}
	var is_found bool = false
	for i := l; i != nil; i = i.next {
		if i.value == value && !is_found {
			is_found = true
			continue
		}
		r = Append(&link{i.value, nil}, r)
	}
	return r
}

func main() {
	var k int = 1
	for k > 0 {
		switch Menu() {
		case 1:
			Add()
		case 2:
			Del()
		case 3:
			Search()
		case 4:
			Show()
		case 9:
			return
		}
	}
}
