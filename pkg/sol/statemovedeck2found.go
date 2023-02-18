package sol

func (s *StateM) moveDeckToSpadesFound(ret []*StateM) []*StateM {
	if s.SpadesFound == byte(KingRank) {
		return ret
	}
	var card = CardC(Spades, Rank(s.SpadesFound+1))
	if !s.StockCardBits.Has(card) {
		return ret
	}
	var ns = s.Derive()
	ns.SpadesFound++
	ns.StockCardBits.RemoveCard(card)
	ret = append(ret, ns)
	return ret
}

func (s *StateM) moveDeckToHeartsFound(ret []*StateM) []*StateM {
	if s.HeartsFound == byte(KingRank) {
		return ret
	}
	var card = CardC(Hearts, Rank(s.HeartsFound+1))
	if !s.StockCardBits.Has(card) {
		return ret
	}
	var ns = s.Derive()
	ns.HeartsFound++
	ns.StockCardBits.RemoveCard(card)
	ret = append(ret, ns)
	return ret
}

func (s *StateM) moveDeckToClubsFound(ret []*StateM) []*StateM {
	if s.ClubsFound == byte(KingRank) {
		return ret
	}
	var card = CardC(Clubs, Rank(s.ClubsFound+1))
	if !s.StockCardBits.Has(card) {
		return ret
	}
	var ns = s.Derive()
	ns.ClubsFound++
	ns.StockCardBits.RemoveCard(card)
	ret = append(ret, ns)
	return ret
}

func (s *StateM) moveDeckToDiamondsFound(ret []*StateM) []*StateM {
	if s.DiamondsFound == byte(KingRank) {
		return ret
	}
	var card = CardC(Diamonds, Rank(s.DiamondsFound+1))
	if !s.StockCardBits.Has(card) {
		return ret
	}
	var ns = s.Derive()
	ns.DiamondsFound++
	ns.StockCardBits.RemoveCard(card)
	ret = append(ret, ns)
	return ret
}
