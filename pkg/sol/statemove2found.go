package sol

func (s *StateM) moveTab2Found(i int) (*StateM, bool) {
	var c = s.PileTable[i].Tail()
	if c.Null() {
		return nil, false
	}
	var suit = c.Suit()
	switch suit {
	case Spades:
		if s.SpadesFound == byte(c.Rank()) {
			var ns = s.Derive()
			ns.SpadesFound++
			ns.PileTable[i].RemoveTail()
			return ns, true
		}
	case Hearts:
		if s.HeartsFound == byte(c.Rank()) {
			var ns = s.Derive()
			ns.HeartsFound++
			ns.PileTable[i].RemoveTail()
			return ns, true
		}
	case Clubs:
		if s.ClubsFound == byte(c.Rank()) {
			var ns = s.Derive()
			ns.ClubsFound++
			ns.PileTable[i].RemoveTail()
			return ns, true
		}
	case Diamonds:
		if s.DiamondsFound == byte(c.Rank()) {
			var ns = s.Derive()
			ns.DiamondsFound++
			ns.PileTable[i].RemoveTail()
			return ns, true
		}
	}
	return nil, false
}
