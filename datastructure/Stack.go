package datastructure

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {

	if v == nil {
		panic("nil is not allowd to push into Stack")
	}

	s.data = append(s.data, v)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) == 0 {
		panic("cannot pop from empty Stack")
	}
	n := len(s.data)
	v := s.data[n-1]
	s.data = s.data[:n-1]

	return v
}

func (s *Stack) Size() int {
	return len(s.data)
}

// import (
//	"fmt"
//	ds "github.com/luzcn/go-code/datastructure"
// )
//
// func main() {
//	stack := ds.Stack{}
//
//	stack.Push(1)
//
//	fmt.Println(stack.Pop())
//	fmt.Println(stack.Pop())
// }
