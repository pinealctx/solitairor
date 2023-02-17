package sol

const (
	endProcByMaxStack  = -1 // end process by max stack size limit
	endProcByMaxSearch = -2 // end process by max search size limit
)

// Road : solution road
type Road struct {
	ForwardStep int
	ReverseStep int
}

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
	endProcReason int

	// record
	roads []Road
}

func NewPuzzle(maxStackSize int, maxSearchSize int) *Puzzle {
	var p = &Puzzle{
		stack:         NewStack(maxStackSize),
		visitRec:      make(map[StateKey]*StateM),
		maxStackSize:  maxStackSize,
		maxSearchSize: maxSearchSize,
		roads:         nil,
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
		if count == endProcByMaxStack {
			p.endProcReason = endProcByMaxStack
			break
		} else if count == endProcByMaxSearch {
			p.endProcReason = endProcByMaxSearch
			break
		} else if count == 0 {
			// no child state, mark it
			var exist = p.visitRec[state.Key()]
			exist.ReverseStep = InfiniteStep
		}
	}
}

func (p *Puzzle) Road() []Road {
	return p.roads
}

func (p *Puzzle) push(childStates ...*StateM) int {
	var size = len(childStates)
	if size+p.stack.Size() >= p.maxStackSize {
		return endProcByMaxStack
	}
	if size+p.searchCount >= p.maxSearchSize {
		return endProcByMaxSearch
	}
	p.searchCount += size

	var count = 0
	for _, child := range childStates {
		if child.IsWin() {
			// already win, record it
			p.roads = append(p.roads, Road{
				ForwardStep: child.ForwardStep,
				ReverseStep: 0,
			})
			child.ReverseBroadcast()
		}

		child.PileTable.Sort()
		var key = child.Key()
		var exist, ok = p.visitRec[key]
		if ok {
			// the node be visited
			if exist.ReverseStep > 0 {
				// the node has a way, skip to search, record it.
				p.roads = append(p.roads, Road{
					ForwardStep: child.ForwardStep,
					ReverseStep: exist.ReverseStep,
				})
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
