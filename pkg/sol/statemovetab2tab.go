package sol

func (s *StateM) moveTab2OtherTab(src, target int) (*StateM, bool) {
	if s.PileTable[target].Empty() {
		// only king can move
		var i = s.PileTable[src].FoundUpKing()
		if i == -1 {
			return nil, false
		}
		var ns = s.Derive()
		ns.PileTable[src].MoveTail2Other(&ns.PileTable[target], i)
		return ns, true
	}
	var c = s.PileTable[target].Tail()
	var i = s.PileTable[src].FoundUpCanFollowSpecCard(c)
	if i == -1 {
		return nil, false
	}
	var ns = s.Derive()
	ns.PileTable[src].MoveTail2Other(&ns.PileTable[target], i)
	return ns, true
}
