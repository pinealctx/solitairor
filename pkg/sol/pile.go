package sol

import (
	"bytes"
)

type Pile struct {
	// Cards is the cards in the pile.
	Cards []Card
}

func PileFromCard(cards ...Card) Pile {
	return PileFromCards(cards)
}

func PileFromCards(cards []Card) Pile {
	var p = Pile{}
	p.AddCards(cards)
	return p
}

func NewPile(cards ...Card) *Pile {
	return NewPileFromCards(cards)
}

func NewPileFromCards(cards []Card) *Pile {
	var p = &Pile{}
	return p.AddCards(cards)
}

func NewPileFromString(s string) *Pile {
	var p = &Pile{}
	p.Decode(s)
	return p
}

func (p *Pile) AddCard(cards ...Card) *Pile {
	return p.AddCards(cards)
}

func (p *Pile) AddCards(cards []Card) *Pile {
	if len(cards) > 0 {
		p.Cards = append(p.Cards, cards...)
	}
	return p
}

func (p *Pile) RemoveTailFromIndex(i int) *Pile {
	p.Cards = p.Cards[:i]
	if i-1 >= 0 {
		p.Cards[i-1].SetFaceUp()
	}
	return p
}

func (p *Pile) RemoveTail() *Pile {
	var l = len(p.Cards)
	return p.RemoveTailFromIndex(l - 1)
}

func (p *Pile) MoveTail2Other(o *Pile, i int) {
	var removes = p.Cards[i:]
	p.RemoveTailFromIndex(i)
	o.AddCard(removes...)
}

func (p *Pile) Encode() string {
	var buf = bytes.NewBuffer(make([]byte, 20))
	buf.Reset()
	for i := range p.Cards {
		buf.WriteByte(byte(p.Cards[i]))
	}
	return buf.String()
}

func (p *Pile) Decode(s string) {
	var l = len(s)
	if l == 0 {
		return
	}
	p.Cards = make([]Card, l)
	for i := 0; i < l; i++ {
		p.Cards[i] = Card(s[i])
	}
}

func (p *Pile) Equals(o *Pile) bool {
	if len(p.Cards) != len(o.Cards) {
		return false
	}
	for i := range p.Cards {
		if p.Cards[i] != o.Cards[i] {
			return false
		}
	}
	return true
}

func (p *Pile) Less(o *Pile) bool {
	var l1 = len(p.Cards)
	var l2 = len(o.Cards)

	if l1 < l2 {
		return true
	} else if l1 > l2 {
		return false
	} else { // l1 == l2
		if l1 == 0 {
			return false
		}
		return p.Cards[0] < o.Cards[0]
	}
}

func (p *Pile) Len() int {
	return len(p.Cards)
}

func (p *Pile) Clone() *Pile {
	var cards = make([]Card, len(p.Cards))
	copy(cards, p.Cards)
	return &Pile{Cards: cards}
}

func (p *Pile) CloneTo(o *Pile) {
	o.Cards = make([]Card, len(p.Cards))
	copy(o.Cards, p.Cards)
}

func (p *Pile) Tail() Card {
	var l = len(p.Cards)
	if l == 0 {
		return 0
	}
	return p.Cards[l-1]
}

func (p *Pile) FoundUpKing() int {
	var l = len(p.Cards)
	if l == 0 {
		return -1
	}
	for i := l - 1; i >= 0; i-- {
		if p.Cards[i].FaceDown() {
			return -1
		}
		if p.Cards[i].Rank() == KingRank {
			return i
		}
	}
	return -1
}

func (p *Pile) FoundUpCanFollowSpecCard(c Card) int {
	var l = len(p.Cards)
	if l == 0 {
		return -1
	}
	for i := l - 1; i >= 0; i-- {
		if p.Cards[i].FaceDown() {
			return -1
		}
		if p.Cards[i].CanFollow(c) {
			return i
		}
	}
	return -1
}

func (p *Pile) Empty() bool {
	return len(p.Cards) == 0
}

func (p *Pile) Count() int {
	return len(p.Cards)
}
