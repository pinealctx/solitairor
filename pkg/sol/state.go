package sol

import (
	"bytes"
	"fmt"
)

const (
	DeckCount = 24
)

var (
	QShuffle   = NewShuffler(QRandBetween)
	Shuffle    = NewShuffler(RandBetween)
	SecShuffle = NewShuffler(SecRandBetween)
)

type StateKey struct {
	// found foundations, spades, hearts, clubs, diamonds
	// spades found count
	// 0 means no card in spades foundation, 1 means A in foundation, 2 means A and 2 in foundation, ...
	SpadesFound   byte
	HeartsFound   byte
	ClubsFound    byte
	DiamondsFound byte
	StockCardBits CardBits

	Tile1Key string
	Tile2Key string
	Tile3Key string
	Tile4Key string
	Tile5Key string
	Tile6Key string
	Tile7Key string
}

type StateM struct {
	// found foundations, spades, hearts, clubs, diamonds
	// spades found count
	// 0 means no card in spades foundation, 1 means A in foundation, 2 means A and 2 in foundation, ...
	SpadesFound   byte
	HeartsFound   byte
	ClubsFound    byte
	DiamondsFound byte
	StockCardBits CardBits

	PileTable PileList

	Parent      *StateM
	ForwardStep int
	ReverseStep int
}

func NewState() *StateM {
	return &StateM{}
}

func NewQRandState() *StateM {
	var cards = GenQRandCards()
	return NewGameStateFromCards(cards)
}

func NewGameStateFromLegacyCards(cards []Card) *StateM {
	var s = NewState()
	for i := 0; i < DeckCount; i++ {
		s.StockCardBits.AddCard(cards[i])
	}
	var cursor = DeckCount
	for i := 0; i < PileCount; i++ {
		for j := 0; j < i+1; j++ {
			if j != i {
				var c = cards[cursor]
				c.SetFaceDown()
				s.PileTable[i].AddCard(c)
			} else {
				s.PileTable[i].AddCard(cards[cursor])
			}
			cursor++
		}
	}
	return s
}

func NewGameStateFromCards(cards []Card) *StateM {
	var s = NewState()
	for i := 0; i < DeckCount; i++ {
		s.StockCardBits.AddCard(cards[i])
	}
	var cursor = DeckCount
	for j := 0; j < PileCount; j++ {
		for i := j; i < PileCount; i++ {
			var c = cards[cursor]
			if i != j {
				c.SetFaceDown()
			}
			s.PileTable[i].AddCard(c)
			cursor++
		}
	}
	return s
}

func NewStateFrom(sf, hf, cf, df byte, stock CardBits, piles ...Pile) *StateM {
	var s = NewState()
	s.SpadesFound = sf
	s.HeartsFound = hf
	s.ClubsFound = cf
	s.DiamondsFound = df
	s.StockCardBits = stock

	var l = len(piles)
	var min = Min[int](l, PileCount)
	for i := 0; i < min; i++ {
		piles[i].CloneTo(&s.PileTable[i])
	}

	return s
}

func (s *StateM) Derive() *StateM {
	var ns = *s
	ns.ForwardStep++
	ns.Parent = s
	ns.ReverseStep = 0
	ns.PileTable = s.PileTable.Clone()
	return &ns
}

func (s *StateM) Key() StateKey {
	var k StateKey
	k.SpadesFound = s.SpadesFound
	k.HeartsFound = s.HeartsFound
	k.ClubsFound = s.ClubsFound
	k.DiamondsFound = s.DiamondsFound
	k.StockCardBits = s.StockCardBits

	for i := 0; i < PileCount; i++ {
		var p = &s.PileTable[i]
		switch i {
		case 0:
			k.Tile1Key = p.Encode()
		case 1:
			k.Tile2Key = p.Encode()
		case 2:
			k.Tile3Key = p.Encode()
		case 3:
			k.Tile4Key = p.Encode()
		case 4:
			k.Tile5Key = p.Encode()
		case 5:
			k.Tile6Key = p.Encode()
		case 6:
			k.Tile7Key = p.Encode()
		}
	}
	return k
}

func (s *StateM) IsWin() bool {
	for i := 0; i < PileCount; i++ {
		var l = s.PileTable[i].Len()
		if l > 0 {
			for j := 0; j < l; j++ {
				if s.PileTable[i].Cards[j].FaceDown() {
					return false
				}
			}
		}
	}
	return true
}

func (s *StateM) SamePiles(o *StateM) bool {
	for i := 0; i < PileCount; i++ {
		if !s.PileTable[i].Equals(&o.PileTable[i]) {
			return false
		}
	}
	return true
}

func (s *StateM) ReverseBroadcast() {
	var p = s
	// reverse broadcast
	for {
		var pp = p.Parent
		if pp == nil {
			break
		}
		if pp.ReverseStep == 0 || pp.ReverseStep > p.ReverseStep+1 {
			pp.ReverseStep = p.ReverseStep + 1
		}
		p = pp
	}
}

func (s *StateM) ReverseBroadcastWithLog() {
	var p = s
	fmt.Println(p)
	// reverse broadcast
	for {
		var pp = p.Parent
		if pp == nil {
			break
		}
		fmt.Println(pp)
		if pp.ReverseStep == 0 || pp.ReverseStep > p.ReverseStep+1 {
			pp.ReverseStep = p.ReverseStep + 1
		}
		p = pp
	}
}

func (s *StateM) String() string {
	var buf = bytes.NewBuffer(make([]byte, 4096))
	buf.Reset()

	buf.WriteString("foundations:")
	buf.WriteString(fmt.Sprintf("spades:%d", s.SpadesFound))
	buf.WriteString(fmt.Sprintf(",hearts:%d", s.HeartsFound))
	buf.WriteString(fmt.Sprintf(",clubs:%d", s.ClubsFound))
	buf.WriteString(fmt.Sprintf(",diamonds:%d", s.DiamondsFound))
	buf.WriteByte('\n')

	buf.WriteString("deck:")
	var cs = GenCards()
	for _, c := range cs {
		if s.StockCardBits.Has(c) {
			buf.WriteString(c.String())
			buf.WriteByte(',')
		}
	}
	buf.WriteByte('\n')

	for i := 0; i < PileCount; i++ {
		buf.WriteString(fmt.Sprintf("pile%d:", i+1))
		var p = &s.PileTable[i]
		for j := 0; j < p.Len(); j++ {
			buf.WriteString(p.Cards[j].String())
			buf.WriteByte(',')
		}
		buf.WriteByte('\n')
	}
	return buf.String()
}
