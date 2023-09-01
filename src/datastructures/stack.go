package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	values []interface{} //Any or Object
}

func (stack Stack) Size() int {
	return len(stack.values)
}

func (stack Stack) IsEmpty() bool {
	return stack.Size() == 0
}
func (stack *Stack) PileUp(value interface{}) {
	stack.values = append(stack.values, value);
}
func (stack *Stack) Unstack() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("Empty stack!")
	}
	removed := stack.values[stack.Size() - 1]
	stack.values = stack.values[:stack.Size() - 1]
	return removed, nil
}

func main() {
	stack1 := Stack{}
	fmt.Println("Stack created with size: ", stack1.Size())
	stack1.PileUp("Go")
	stack1.PileUp(2009)
	stack1.PileUp(3.14)
	stack1.PileUp("FIM")
	stack1.PileUp("FIM2")
	fmt.Printf("New size from stack : %d\n", stack1.Size())
	removed, _ := stack1.Unstack()
	fmt.Println(removed)
}