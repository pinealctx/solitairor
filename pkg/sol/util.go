package sol

// Stack : stack of state
type Stack struct {
	// stack of state
	states []*StateM
	count  int
}

// NewStack : create a new stack
func NewStack(size int) *Stack {
	return &Stack{states: make([]*StateM, 0, size)}
}

// Push : push a state to stack
func (s *Stack) Push(state *StateM) {
	s.states = append(s.states[:s.count], state)
	s.count++
}

// Pop : pop a state from stack
func (s *Stack) Pop() *StateM {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.states[s.count]
}

// Size : return the size of stack
func (s *Stack) Size() int {
	return s.count
}

func Min[T comparableT](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T comparableT](a, b T) T {
	if a > b {
		return a
	}
	return b
}

type comparableT interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64
}
