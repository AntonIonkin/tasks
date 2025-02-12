package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func main() {
	//SwapTwoNumbers
	a := 3
	b := 5
	fmt.Printf("SwapTwoNumbers\na=%d b=%d\ta=%p b=%p\n", a, b, &a, &b)
	SwapTwoNumbers(&a, &b)
	fmt.Printf("a=%d b=%d\ta=%p b=%p\n", a, b, &a, &b)

	//ModifySliceCopy
	sliceInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("ModifySliceCopy\n%v\n%v\n", sliceInt, ModifySliceCopy(sliceInt))

	//minElement
	c := []int{5, 3, 8, 2, 7}
	cNil := []int{}
	fmt.Println("minElement")
	fmt.Println(minElement(c))
	fmt.Println(minElement(cNil))

	//createList
	fmt.Println("createList")
	fmt.Println(createList(5))

	//minPointer
	x1 := 2
	x2 := 3
	q := minPointer(&x1, &x2)
	fmt.Println("minPointer")
	fmt.Printf("%v=%p", *q, q)
}

func SwapTwoNumbers(a, b *int) {
	*a, *b = *b, *a
}

func ModifySliceCopy(s []int) []int {
	sliceIntNew := make([]int, 0, cap(s))
	for i, v := range s {
		switch v % 2 {
		case 0:
			sliceIntNew = append(sliceIntNew, s[i]+1)
		case 1:
			sliceIntNew = append(sliceIntNew, s[i]-1)
		}

	}
	return sliceIntNew
}

func minElement(slice []int) *int {
	if len(slice) == 0 {
		return nil
	}

	min := slice[0]

	for i := 0; i < len(slice); i++ {
		if slice[i] < min {
			min = slice[i]
		}
	}

	return &min
}

func createList(n int) *Node {
	return &Node{Value: n}
}

func minPointer(a, b *int) *int {
	switch {
	case a == nil && b == nil:
		return nil
	case a == nil:
		return b
	case b == nil:
		return a
	case *a > *b:
		return b
	case *a < *b:
		return a
	}

	return nil
}
