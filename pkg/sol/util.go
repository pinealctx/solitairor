package sol

import (
	crd "crypto/rand"
	"encoding/binary"
	"io"
	"math/rand"
	"time"
)

// Swapper : swap two elements
type Swapper interface {
	Swap(i, j int)
	Len() int
}

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

// Shuffler : shuffle data
type Shuffler struct {
	//rand function between [min, max]
	randFunc func(min int, max int) int
}

func NewShuffler(randFunc func(min int, max int) int) Shuffler {
	return Shuffler{randFunc: randFunc}
}

// Shuffle : shuffle data
/*Reference from sort package, which has mentioned Fisher–Yates shuffle.
https://en.wikipedia.org/wiki/Fisher%E2%80%93Yates_shuffle
*/
func (s Shuffler) Shuffle(ds Swapper) {
	var size = ds.Len()
	for i := size - 1; i >= 0; i-- {
		var j = s.randFunc(0, i)
		ds.Swap(i, j)
	}
}

/*
math.Rand实际上是伪随机数，如果seed够随机，它也能做出足够随机的数。
这里会实现两种随机函数，一种是读取计算机中的随机墒，另一种是取时间的NanoUnix。
*/

// RandBetween : generate rand between [a, b], return int
func RandBetween(min int, max int) int {
	if min > max {
		panic("invalid min and max of RandBetween")
	}
	if min == max {
		//no choice
		return min
	}
	/* #nosec */
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

// QRandBetween : quick generate rand between [a, b], use same global rand source: rand.Rand
func QRandBetween(min int, max int) int {
	if min > max {
		panic("invalid min and max of QRandBetween")
	}
	if min == max {
		//no choice
		return min
	}
	/* #nosec */
	rand.Seed(time.Now().UnixNano())
	/* #nosec */
	return rand.Intn(max-min+1) + min
}

// SecRandBetween : generate rand between [a, b], return int
func SecRandBetween(min int, max int) int {
	if min > max {
		panic("invalid min and max of SecRandBetween")
	}
	if min == max {
		//no choice
		return min
	}
	var buf [8]byte
	var _, err = io.ReadFull(crd.Reader, buf[:])
	if err != nil {
		return RandBetween(min, max)
	}
	var seed = int64(binary.LittleEndian.Uint64(buf[:]))
	/* #nosec */
	var r = rand.New(rand.NewSource(seed))
	return r.Intn(max-min+1) + min
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
