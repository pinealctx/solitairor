package sol

import "math"

type Puzzle struct {
	stack    *Stack
	visitRec map[StateKey]*StateM

	// max stack size
	maxStackSize int

	// max search size
	maxSearchSize int
	// search count
	searchCount int

	// end proc reason (0-normal -1 -- max stack size -2 -- max search size)
	endProcReason EndProcReason

	// record hit
	// key is the step, value is counter.
	hit map[int]int
	// findFunc
	findFunc func(*StateM) bool
}

func NewPuzzle(maxStackSize int, maxSearchSize int) *Puzzle {
	var p = &Puzzle{
		stack:         NewStack(maxStackSize),
		visitRec:      make(map[StateKey]*StateM),
		maxStackSize:  maxStackSize,
		maxSearchSize: maxSearchSize,
		hit:           make(map[int]int),
		findFunc: func(state *StateM) bool {
			return state.IsWin()
		},
	}
	return p
}

func (p *Puzzle) InitRoot(root *StateM) {
	p.push(root)
}

func (p *Puzzle) Run() {
	for p.stack.Size() > 0 {
		var state = p.stack.Pop()
		var childStates = state.MoveNext()
		var count = p.push(childStates...)
		if count == EndProcByMaxStack {
			p.endProcReason = EndProcByMaxStack
			break
		} else if count == EndProcByMaxSearch {
			p.endProcReason = EndProcByMaxSearch
			break
		}
		// Actually, no need to set node reverse step to infinite.
		/* else if count == 0 {
			// no child state, mark it
			var exist = p.visitRec[state.Key()]
			exist.ReverseStep = InfiniteStep
		}*/
	}
}

func (p *Puzzle) SolutionCount() int {
	return len(p.hit)
}

func (p *Puzzle) Record(r *Record) {
	r.SolutionCount = len(p.hit)
	r.MaxStackSize = p.maxStackSize
	r.MaxSearchSize = p.maxSearchSize
	r.SearchCount = p.searchCount
	r.EndProcReason = int(p.endProcReason)

	if r.SolutionCount == 0 {
		return
	}

	var minStep = math.MaxInt
	var maxStep = 0
	var sumStep = 0
	var sumCount = 0
	var sumDiff = 0
	var averageStep = 0

	for step, count := range p.hit {
		if step < minStep {
			minStep = step
		}
		if step > maxStep {
			maxStep = step
		}
		sumStep += step * count
		sumCount += count
	}
	if sumStep > 0 {
		averageStep = sumStep / sumCount
	}
	for step, count := range p.hit {
		var diff = step - averageStep
		if diff < 0 {
			diff = -diff
		}
		sumDiff += diff * count
	}

	r.AverageStep = averageStep
	r.DiffStep = sumDiff / sumCount
	r.MinStep = minStep
	r.MaxStep = maxStep
}

func (p *Puzzle) push(childStates ...*StateM) EndProcReason {
	var size = len(childStates)
	if size+p.stack.Size() >= p.maxStackSize {
		return EndProcByMaxStack
	}
	if size+p.searchCount >= p.maxSearchSize {
		return EndProcByMaxSearch
	}
	p.searchCount += size

	var count EndProcReason
	for _, child := range childStates {
		if p.findFunc(child) {
			// already win, record it
			p.hit[child.ForwardStep]++
			child.ReverseBroadcast()
			// no need to push into stack
			continue
		}

		child.PileTable.Sort()
		var key = child.Key()
		var exist, ok = p.visitRec[key]
		if ok {
			// the node be visited
			if exist.ReverseStep > 0 {
				// the node has a way, skip to search, record it.
				p.hit[child.ForwardStep+exist.ReverseStep]++
			}
			// skip to push into stack
			continue
		}

		p.stack.Push(child)
		p.visitRec[key] = child
		count++
	}
	return count
}
