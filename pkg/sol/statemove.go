package sol

func (s *StateM) MoveNext() []*StateM {
	var ret []*StateM
	// priority from low to high
	// 1. move from deck
	for i := 0; i < PileCount; i++ {
		var ns = s.moveFromDeck(i)
		if len(ns) > 0 {
			ret = append(ret, ns...)
		}
	}

	// 2. move from table to found
	for i := 0; i < PileCount; i++ {
		var n, ok = s.moveTab2Found(i)
		if ok {
			ret = append(ret, n)
		}
	}

	// 3. move from table to table
	for i := 0; i < PileCount; i++ {
		for j := 0; j < PileCount; j++ {
			if i == j {
				continue
			}
			var n, ok = s.moveTab2OtherTab(i, j)
			if ok {
				ret = append(ret, n)
			}
		}
	}

	return ret
}
