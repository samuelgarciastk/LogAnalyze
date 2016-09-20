package common

//Stack is a stack implement, but not thread safe.
type Stack []interface{}

//Push adds value v into stack s.
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}

//Pop removes the value on the top of stack s and return this value.
//if stack s is empty, it will return nil and false.
func (s *Stack) Pop() (interface{}, bool) {
	l := len(*s)
	if l < 1 {
		return nil, false
	}
	defer func() {
		*s = (*s)[:l-1]
	}()
	return (*s)[l-1], true
}
