package main

import (
	"fmt"
)

type theQuestion struct {
	a int
	b int
}

func (i theQuestion) Cons() []int {
	return []int{i.a, i.b}
}

func (i theQuestion) Car() int {
	return i.a
}

func (i theQuestion) Cdr() int {
	return i.b
}

func main() {
	fmt.Println("vim-go")
}
