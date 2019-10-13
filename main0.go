package main

import "fmt"
import "math"

type list struct {
  cof int
  pow float64
  next *list
}

func printFormual(list *list, s string, res int) {
  fmt.Print(s)
  fmt.Print(" = ")
  for p := list; p != nil; p = p.next {
    prefix := " + "
    if p.next == nil {
      prefix = ""
    }
    f := fmt.Sprintf("(%d * x^%d)%s", p.cof, int(p.pow), prefix);
    fmt.Print(f)
  }
  ff := fmt.Sprintf(" = %d", res)
  fmt.Print(ff)
  fmt.Println()
}

func showList(list *list) {
  fmt.Println()
  fmt.Println()
  for p := list; p != nil; p = p.next {
    fmt.Print(p)
  }
  fmt.Println();
}

func AddItem(newNode, list *list) *list {
  if list == nil {
    return list
  }
  for p := list; p != nil; p = p.next {
    if p.next == nil {
      p.next = newNode
      return list
    }
   }
  return list
}

func Add(p, q, result *list) *list {
  for i := p; i != nil; i = i.next {
    result = AddItem(&list{i.cof, i.pow, nil}, result)
    if q != nil {
      result = AddItem(&list{q.cof, q.pow, nil}, result)
    }
    q = q.next
  }
  return result.next
}

func Meaning(list *list, x *float64) int {
  res := 0
  for p := list; p != nil; p = p.next {
    temp := math.Pow(*x, p.pow)
    res += p.cof * int(temp)
  }
  return res;
}

func Equality(qList *list, pList *list) {
  res := true

  q := qList
  for p := pList; p != nil; p = p.next {
    if p.cof != q.cof && p.pow != q.pow {
      res = false
      break
    }
    q = q.next
  }

  if res {
    fmt.Println("Многочлены равны")
  } else {
    fmt.Println("Многочлены не равны")
  }
}

func main() {
  P1 := &list{6, 3, nil}
  P2 := &list{2, 2, nil}
  P3 := &list{-1, 1, nil}
  P4 := &list{-9, 0, nil}

  p := P1

  var x float64
  fmt.Print("Enter x: ")
  fmt.Scanln(&x)

  p = AddItem(P2, p)
  p = AddItem(P3, p)
  p = AddItem(P4, p)

  Q1 := &list{-7, 6, nil}
  Q2 := &list{-4, 4, nil}
  Q3 := &list{1, 1, nil}
  Q4 := &list{2, 0, nil}

  q := Q1
  q = AddItem(Q2, q)
  q = AddItem(Q3, q)
  q = AddItem(Q4, q)

  r := &list{0, 0, nil}
  r = Add(p, q, r)

  printFormual(p, "P(x)", Meaning(p, &x))
  printFormual(q, "Q(x)", Meaning(q, &x))
  Equality(q, p)
  printFormual(r, "R(x)", Meaning(r, &x))

  showList(p)
  showList(q)
  showList(r)
}
