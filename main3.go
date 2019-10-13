package main

import (
	"fmt"
	"strings"
)

type stack struct {
	data string
	next *stack
}

func Priority(v string) int {
	switch v {
	case "^":
		return 4
	case "*":
		return 3
	case "/":
		return 3
	case "-":
		return 2
	case "+":
		return 2
	case "(":
		return 1
	default:
		return 0
	}
}

func AddStack(data string, p *stack) *stack {
	return &stack{data, p}
}

func AddList(data string, p *stack) {
	element := &stack{data, nil}
	if p.data == "|" {
		p.data = data
		p.next = nil
	} else {
		for z := p; z != nil; z = z.next {
			if z.next == nil {
				z.next = element
				return
			}
		}
	}
}

func PrintStack(p *stack) {
	for z := p; z != nil; z = z.next {
		if z.data != "|" {
			fmt.Print(z.data)
		}
	}
	fmt.Println()
}

func PrintStack2(p *stack) {
	line := ""
	for z := p; z != nil; z = z.next {
		if z.data != "|" {
			line += z.data
		}
	}
	res := Revers(line)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])
	}
	fmt.Println()
}

func DelStack(p *stack) *stack {
	if p.next == nil {
		p.data = "|"
		return p
	}
	var element *stack
	element = p.next
	return element
}

func Postfix(inect []string) *stack {
	list := &stack{"|", nil}
	ste := &stack{"|", nil}

	for i := 0; i != len(inect); i++ {
		if inect[i] >= "a" && inect[i] <= "z" {
			AddList(inect[i], list)
		} else if Priority(inect[i]) == 1 {
			ste = AddStack(inect[i], ste)
		} else if Priority(inect[i]) != 0 && Priority(inect[i]) != 1 {
			if Priority(ste.data) < Priority(inect[i]) || ste.next == nil {
				ste = AddStack(inect[i], ste)
			} else {
				for Priority(ste.data) >= Priority(inect[i]) {
					AddList(ste.data, list)
					ste = DelStack(ste)
				}
				ste = AddStack(inect[i], ste)
			}
		}
		if inect[i] == ")" {
			for ste.data != "(" {
				AddList(ste.data, list)
				ste = DelStack(ste)
			}
			ste = DelStack(ste)
		}
	}
	for ste.next != nil {
		if ste.data != "(" {
			AddList(ste.data, list)
		}
		ste = DelStack(ste)
	}
	return list
}

func Prefix(inect []string) *stack {
	list := &stack{"|", nil}
	blum := make([]string, len(inect))
	buff := make([]string, len(inect))
	copy(blum, inect)
	copy(buff, inect)
	for i := len(blum) - 1; i >= 0; i-- {
		if blum[i] == "(" {
			buff[len(blum)-i-1] = ")"
		} else {
			if blum[i] == ")" {
				buff[len(blum)-i-1] = "("
			} else {
				buff[len(blum)-i-1] = blum[i]
			}
		}
	}
	list = Postfix(buff)
	l := &stack{"|", nil}
	for z := list; z != nil; z = z.next {
		l = AddStack(z.data, l)
	}
	return l
}

func main() {
	head := &stack{"|", nil}
	var inect string
	fmt.Print("Введите формулу в инфиксе: ")
	fmt.Scanln(&inect)
	rower1 := strings.Split(inect, "")
	head = Postfix(rower1)
	fmt.Print("Постфиксное: ")
	PrintStack(head)
	fmt.Print("Префиксное: ")
	head = Prefix(rower1)
	PrintStack(head)
}

func Revers(s string) []string {
	rower := strings.Split(s, "")
	result := strings.Split(s, "")
	x := 0
	for i := len(rower) - 1; i >= 0; i-- {
		if rower[i] == "(" {
			rower[i] = ")"
		} else if rower[i] == ")" {
			rower[i] = "("
		}
		result[x] = rower[i]
		x++
	}
	return result
}
