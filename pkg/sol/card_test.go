package sol

import "testing"

func TestCard_SpadesCards_String(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		t.Log(SpadesCards[i])
		t.Log(SpadesDownCards[i])
	}
}

func TestCard_HeartsCards_String(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		t.Log(HeartsCards[i])
		t.Log(HeartsDownCards[i])
	}
}

func TestCard_ClubsCards_String(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		t.Log(ClubsCards[i])
		t.Log(ClubsDownCards[i])
	}
}

func TestCard_DiamondsCards_String(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		t.Log(DiamondsCards[i])
		t.Log(DiamondsDownCards[i])
	}
}

func TestCard_FaceDown(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		var x = SpadesCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceDown()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = HeartsCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceDown()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = ClubsCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceDown()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = DiamondsCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceDown()
		t.Log(x)
	}
}

func TestCard_SetFaceUp(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		var x = SpadesDownCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceUp()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = HeartsDownCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceUp()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = ClubsDownCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceUp()
		t.Log(x)
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = DiamondsDownCards[i]
		t.Log(x)
		x.SetFaceDown()
		x.SetFaceUp()
		t.Log(x)
	}
}

func TestCard_Rank(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		var x = SpadesCards[i]
		t.Log(x)
		t.Log(x.Rank())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = HeartsCards[i]
		t.Log(x)
		t.Log(x.Rank())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = ClubsCards[i]
		t.Log(x)
		t.Log(x.Rank())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = DiamondsCards[i]
		t.Log(x)
		t.Log(x.Rank())
	}
}

func TestCard_Suit(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		var x = SpadesCards[i]
		t.Log(x)
		t.Log(x.Suit())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = HeartsCards[i]
		t.Log(x)
		t.Log(x.Suit())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = ClubsCards[i]
		t.Log(x)
		t.Log(x.Suit())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = DiamondsCards[i]
		t.Log(x)
		t.Log(x.Suit())
	}
}

func TestCard_IsRed(t *testing.T) {
	for i := 0; i < CardCountOfColor; i++ {
		var x = SpadesCards[i]
		t.Log(x.IsRed())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = HeartsCards[i]
		t.Log(x.IsRed())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = ClubsCards[i]
		t.Log(x.IsRed())
	}

	for i := 0; i < CardCountOfColor; i++ {
		var x = DiamondsCards[i]
		t.Log(x.IsRed())
	}
}

func TestCardZero(t *testing.T) {
	var card Card
	t.Log(card.Suit(), card.Rank(), card.Seq(), card.FaceDown())
}

func TestCards_Txt(t *testing.T) {
	var r = GenQRandCards()
	t.Log(r.Txt())
	t.Log(r.JTxt())
}
