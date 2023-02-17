package sol

func (s *StateM) moveFromDeck(i int) []*StateM {
	if s.PileTable[i].Empty() {
		// only king can move
		return s.moveFromDeckCards(i, SpadesKing, HeartsKing, ClubsKing, DiamondsKing)
	}
	return s.moveFromDeckCards(i, s.PileTable[i].Tail().FollowCards()...)
}

func (s *StateM) moveFromDeckCards(i int, cards ...Card) []*StateM {
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
