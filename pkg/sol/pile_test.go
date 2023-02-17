package sol

import "testing"

func TestNewPile(t *testing.T) {
	var p = NewPile()
	if len(p.Cards) != 0 {
		t.Fatal("NewPile() failed")
		return
	}
	p = NewPile(CardC(Hearts, 1), CardC(Diamonds, 11), CardC(Clubs, 10))
	if len(p.Cards) != 3 {
		t.Fatal("NewPile() failed")
		return
	}
}

func TestPile_RemoveTailFromIndex(t *testing.T) {
	testPileRemoveTailFromIndex(t, 0)
	testPileRemoveTailFromIndex(t, 1)
	testPileRemoveTailFromIndex(t, 2)
}

func TestPile_Encode(t *testing.T) {
	var p = NewPile()
	testPileEncode(t, p)
	p = NewPile(CardC(Hearts, 1), CardC(Diamonds, 11), CardC(Clubs, 10))
	testPileEncode(t, p)
	p = NewPile(CardCD(Hearts, 1), CardC(Diamonds, 11), CardC(Clubs, 10))
	testPileEncode(t, p)
}

func TestPile_MoveTail2Other(t *testing.T) {
	testPileMoveTail2Other(t, 0)
	testPileMoveTail2Other(t, 1)
	testPileMoveTail2Other(t, 2)
}

func testPileEncode(t *testing.T, p *Pile) {
	var s = p.Encode()
	var q = NewPileFromString(s)
	if !p.Equals(q) {
		t.Fatal("Pile.Encode() failed")
		return
	}
	t.Log(q)
}

func testPileRemoveTailFromIndex(t *testing.T, index int) {
	var p = NewPile(CardC(Hearts, 1), CardC(Diamonds, 11), CardC(Clubs, 10))
	t.Log(p.Cards)
	p.RemoveTailFromIndex(index)
	if len(p.Cards) != index {
		t.Fatal("RemoveTailFromIndex() failed")
		return
	}
	t.Log(p.Cards)
}

func testPileMoveTail2Other(t *testing.T, index int) {
	var p = NewPile(CardC(Hearts, 1), CardC(Diamonds, 11), CardC(Clubs, 10))
	var q = NewPile()
	p.MoveTail2Other(q, index)
	if len(p.Cards) != index || len(q.Cards) != 3-index {
		t.Fatal("MoveTail2Other() failed")
		return
	}
	t.Log(p.Cards, q.Cards)
}
