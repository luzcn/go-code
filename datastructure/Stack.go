package datastructure

type Stack interface {
	Push(v interface{})
	Pop() interface{}
	Size() int
}

type stack struct {
	data []interface{}
}

func (s *stack) Push(v interface{}) {

	if v == nil {
		panic("nil is not allowd to push into stack")
	}

	s.data = append(s.data, v)
}

func (s *stack) Pop() interface{} {
	if len(s.data) == 0 {
		panic("cannot pop from empty stack")
	}
	n := len(s.data)
	v := s.data[n-1]
	s.data = s.data[:n-1]

	return v
}

func (s *stack) Size() int {
	return len(s.data)
}

func NewStack() Stack {
	return &stack{}
}

// import (
//	"fmt"
//	ds "github.com/luzcn/go-code/datastructure"
// )
//
// func main() {
//	stack := ds.stack{}
//
//	stack.Push(1)
//
//	fmt.Println(stack.Pop())
//	fmt.Println(stack.Pop())
// }
