package sol

func (s *StateM) moveDeckToTab(i int) []*StateM {
	if s.PileTable[i].Empty() {
		// only king can move
		return s.moveDeckCardsToTab(i, SpadesKing, HeartsKing, ClubsKing, DiamondsKing)
	}
	return s.moveDeckCardsToTab(i, s.PileTable[i].Tail().FollowCards()...)
}

func (s *StateM) moveDeckToEmptyTab(i int) []*StateM {
	if s.PileTable[i].Empty() {
		// only king can move
		return s.moveDeckCardsToTab(i, SpadesKing, HeartsKing, ClubsKing, DiamondsKing)
	}
	return nil
}

func (s *StateM) moveDeckToNoEmptyTab(i int) []*StateM {
	if s.PileTable[i].Empty() {
		return nil
	}
	return s.moveDeckCardsToTab(i, s.PileTable[i].Tail().FollowCards()...)
}

func (s *StateM) moveDeckCardsToTab(i int, cards ...Card) []*StateM {
	var ret []*StateM
	for j := range cards {
		if s.StockCardBits.Has(cards[j]) {
			var ns = s.Derive()
			ns.StockCardBits.RemoveCard(cards[j])
			ns.PileTable[i].AddCard(cards[j])
			ret = append(ret, ns)
		}
	}
	return ret
}
