package sol

func (s *StateM) MoveNext() []*StateM {
	var ret []*StateM
	// priority from low to high
	// 1. move from deck
	ret = s.handleMoveDeckToNoEmptyTab(ret)
	ret = s.handleMoveDeckToEmptyTab(ret)
	//ret = s.handleMoveDeckToTab(ret)
	ret = s.handleMoveDeckToFound(ret)

	// 2. move from table to table
	ret = s.handleMoveToOtherTab(ret)

	// 3. move from table to found
	ret = s.handleMoveToFound(ret)

	return ret
}

func (s *StateM) Fake() {
	_ = s.handleMoveDeckToTab
}

func (s *StateM) handleMoveDeckToTab(ret []*StateM) []*StateM {
	if !s.StockCardBits.Empty() {
		for i := 0; i < PileCount; i++ {
			var ns = s.moveDeckToTab(i)
			if len(ns) > 0 {
				ret = append(ret, ns...)
			}
		}
	}
	return ret
}

func (s *StateM) handleMoveDeckToNoEmptyTab(ret []*StateM) []*StateM {
	if !s.StockCardBits.Empty() {
		for i := 0; i < PileCount; i++ {
			var ns = s.moveDeckToNoEmptyTab(i)
			if len(ns) > 0 {
				ret = append(ret, ns...)
			}
		}
	}
	return ret
}

func (s *StateM) handleMoveDeckToEmptyTab(ret []*StateM) []*StateM {
	if !s.StockCardBits.Empty() {
		for i := 0; i < PileCount; i++ {
			var ns = s.moveDeckToEmptyTab(i)
			if len(ns) > 0 {
				ret = append(ret, ns...)
			}
		}
	}
	return ret
}

func (s *StateM) handleMoveDeckToFound(ret []*StateM) []*StateM {
	ret = s.moveDeckToSpadesFound(ret)
	ret = s.moveDeckToHeartsFound(ret)
	ret = s.moveDeckToClubsFound(ret)
	ret = s.moveDeckToDiamondsFound(ret)
	return ret
}

func (s *StateM) handleMoveToFound(ret []*StateM) []*StateM {
	for i := 0; i < PileCount; i++ {
		var n, ok = s.moveTab2Found(i)
		if ok {
			ret = append(ret, n)
		}
	}
	return ret
}

func (s *StateM) handleMoveToOtherTab(ret []*StateM) []*StateM {
	for i := PileCount - 1; i >= 0; i-- {
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
