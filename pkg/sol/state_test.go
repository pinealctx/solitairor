package sol

import "testing"

type VisitRec map[StateKey]struct{}

func NewVisitRec() VisitRec {
	return make(VisitRec)
}

func (v VisitRec) RecordIfNotVisited(s *StateM) bool {
	var k = s.Key()
	if _, ok := v[k]; ok {
		return false
	}
	v[k] = struct{}{}
	return true
}

func TestVisitRec_RecordIfNotVisited(t *testing.T) {
	var v = NewVisitRec()
	var s = NewState()
	var ok = v.RecordIfNotVisited(s)
	if !ok {
		t.Fatal("should be ok")
	}
	ok = v.RecordIfNotVisited(s)
	if ok {
		t.Fatal("should not be ok")
	}

	var ns = NewStateFrom(1, 2, 3, 4, 5, Pile{Cards: []Card{
		CardC(Spades, 1),
	}})
	ok = v.RecordIfNotVisited(ns)
	if !ok {
		t.Fatal("should be ok")
	}
	ok = v.RecordIfNotVisited(ns)
	if ok {
		t.Fatal("should not be ok")
	}
}

func TestStateM_Key(t *testing.T) {
	var cards0 = []Card{
		SpadesCards[12],
		HeartsCards[1],
		ClubsCards[2],
		DiamondsCards[3],
	}
	var cards1 = []Card{
		SpadesCards[1],
		ClubsCards[1],
		HeartsCards[2],
		SpadesCards[9],
	}
	var cards2 = []Card{
		SpadesCards[11],
		ClubsCards[3],
		DiamondsCards[1],
	}
	var cards3 = []Card{
		SpadesCards[2],
	}
	var cards4 = []Card{
		DiamondsCards[2],
	}
	var cards5 = []Card{
		SpadesCards[10],
	}

	var p1 PileList
	p1.SetPileCards(0, cards0)
	p1.SetPileCards(1, cards1)
	p1.SetPileCards(2, cards2)
	p1.SetPileCards(3, cards3)
	p1.SetPileCards(4, cards4)
	p1.SetPileCards(5, cards5)
	var n1 = NewStateFrom(0, 0, 0, 0, 0, p1[:]...)

	var p2 PileList
	p2.SetPileCards(6, cards0)
	p2.SetPileCards(5, cards1)
	p2.SetPileCards(4, cards2)
	p2.SetPileCards(3, cards3)
	p2.SetPileCards(2, cards4)
	p2.SetPileCards(1, cards5)
	var n2 = NewStateFrom(0, 0, 0, 0, 0, p2[:]...)

	n1.PileTable.Sort()
	n2.PileTable.Sort()
	t.Log(n1.Key() == n2.Key())
}

func TestNewQRandState(t *testing.T) {
	var s = NewQRandState()
	for i := 0; i < PileCount; i++ {
		t.Log(s.PileTable[i].Cards)
	}
}

/*
Deck             Spades    Hearts    Clubs   Diamonds
[]                [Q]       [Q]       [Q]      [Q]

	Tile1    Tile2    Tile3    Tile4   Tile5   Tile6   Tile7
	  ♠K      ♦K       ♣K      ♥K
*/
func genSimpleState1() *StateM {
	var tile1 = PileFromCards([]Card{
		SpadesCards[KingRank-1],
	})

	var tile2 = PileFromCards([]Card{
		DiamondsCards[KingRank-1],
	})

	var tile3 = PileFromCards([]Card{
		ClubsCards[KingRank-1],
	})

	var tile4 = PileFromCards([]Card{
		HeartsCards[KingRank-1],
	})

	return NewStateFrom(QueenRank, QueenRank, QueenRank, QueenRank, 0, tile1, tile2, tile3, tile4)
}

/*
Deck             Spades    Hearts    Clubs   Diamonds
[]                [9]       [9]       [9]      [9]

	Tile1    Tile2    Tile3    Tile4   Tile5   Tile6   Tile7
	  ♠K      ♦K       ♣K      ♥K
	  ♦Q      ♣Q       ♥Q      ♠Q
	  ♣J      ♥J       ♠J      ♦J
	  ♠10     ♦10      ♣10     ♥10
*/
func genSimpleState2() *StateM {
	var tile1 = PileFromCards([]Card{
		SpadesCards[KingRank-1],
		DiamondsCards[QueenRank-1],
		ClubsCards[JackRank-1],
		HeartsCards[10-1],
	})

	var tile2 = PileFromCards([]Card{
		DiamondsCards[KingRank-1],
		ClubsCards[QueenRank-1],
		HeartsCards[JackRank-1],
		SpadesCards[10-1],
	})

	var tile3 = PileFromCards([]Card{
		ClubsCards[KingRank-1],
		HeartsCards[QueenRank-1],
		SpadesCards[JackRank-1],
		DiamondsCards[10-1],
	})

	var tile4 = PileFromCards([]Card{
		HeartsCards[KingRank-1],
		SpadesCards[QueenRank-1],
		DiamondsCards[JackRank-1],
		ClubsCards[10-1],
	})

	return NewStateFrom(9, 9, 9, 9, 0, tile1, tile2, tile3, tile4)
}
