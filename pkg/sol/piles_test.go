package sol

import "testing"

func TestPileList_Sort(t *testing.T) {
	var p PileList

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

	p.SetPileCards(0, cards0)
	p.SetPileCards(1, cards1)
	p.SetPileCards(2, cards2)
	p.SetPileCards(3, cards3)
	p.SetPileCards(4, cards4)
	p.SetPileCards(5, cards5)

	testLogPileList(t, &p)

	t.Log("")
	t.Log("")

	p.Sort()
	testLogPileList(t, &p)
}

func testLogPileList(t *testing.T, p *PileList) {
	for i := range p {
		t.Log(p[i])
	}
}
