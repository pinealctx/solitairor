package sol

import "bytes"

const (
	Spades   Suit = 0
	Hearts   Suit = 1
	Clubs    Suit = 2
	Diamonds Suit = 3
)

const (
	AceRank        = 1
	KingRank  Rank = 13
	QueenRank      = 12
	JackRank       = 11
)

const (
	CardCountOfColor = 13
)

var (
	SpadesCards   [CardCountOfColor]Card
	HeartsCards   [CardCountOfColor]Card
	ClubsCards    [CardCountOfColor]Card
	DiamondsCards [CardCountOfColor]Card

	SpadesDownCards   [CardCountOfColor]Card
	HeartsDownCards   [CardCountOfColor]Card
	ClubsDownCards    [CardCountOfColor]Card
	DiamondsDownCards [CardCountOfColor]Card

	SpadesKing   = CardC(Spades, KingRank)
	HeartsKing   = CardC(Hearts, KingRank)
	ClubsKing    = CardC(Clubs, KingRank)
	DiamondsKing = CardC(Diamonds, KingRank)
)

// Suit : card suit (Spades, Hearts,  Clubs, Diamonds) 黑红梅方
type Suit byte

func (x Suit) String() string {
	switch x {
	case Spades:
		return "♠️"
	case Hearts:
		return "♥️"
	case Clubs:
		return "♣️"
	case Diamonds:
		return "♦️"
	}
	return "?"
}

func (x Suit) CanFollow(y Suit) bool {
	switch x {
	case Spades:
		return y == Hearts || y == Diamonds
	case Hearts:
		return y == Spades || y == Clubs
	case Clubs:
		return y == Hearts || y == Diamonds
	case Diamonds:
		return y == Spades || y == Clubs
	}
	return false
}

// Rank : card rank (1-13)
type Rank byte

func (x Rank) String() string {
	switch x {
	case AceRank:
		return "A"
	case 10:
		return "10"
	case JackRank:
		return "J"
	case QueenRank:
		return "Q"
	case KingRank:
		return "K"
	}
	return string('0' + x)
}

// Card : card define
// 0-3 : rank
// 4-5 : suit
// 6: face down
type Card byte

func (x Card) Rank() Rank {
	return Rank(x & 0b00001111)
}

func (x Card) Suit() Suit {
	return Suit((x >> 4) & 0b00000011)
}

func (x Card) FaceDown() bool {
	return x&0b01000000 != 0
}

func (x Card) IsRed() bool {
	return x.Suit() == Hearts || x.Suit() == Diamonds
}

func (x Card) Seq() int {
	var suit = x.Suit()
	return int(suit)*CardCountOfColor + int(x.Rank()) - 1
}

func (x Card) Null() bool {
	return x == 0
}

func (x Card) CanFollow(y Card) bool {
	if x.Rank() != y.Rank()-1 {
		return false
	}
	return x.Suit().CanFollow(y.Suit())
}

func (x Card) FollowCards() []Card {
	var rank = x.Rank()
	if rank == AceRank || x.Null() {
		return nil
	}
	rank--
	switch x.Suit() {
	case Spades:
		return []Card{CardC(Hearts, rank), CardC(Diamonds, rank)}
	case Hearts:
		return []Card{CardC(Spades, rank), CardC(Clubs, rank)}
	case Clubs:
		return []Card{CardC(Hearts, rank), CardC(Diamonds, rank)}
	case Diamonds:
		return []Card{CardC(Spades, rank), CardC(Clubs, rank)}
	}
	return nil
}

func (x Card) String() string {
	if x.FaceDown() {
		return x.Suit().String() + x.Rank().String() + "*"
	}
	return x.Suit().String() + x.Rank().String()
}

func (x *Card) SetFaceDown() {
	*x |= 0b01000000
}

func (x *Card) SetFaceUp() {
	*x &^= 0b01000000
}

func (x *Card) SetRank(r Rank) {
	*x = (*x & 0b11110000) | Card(r)
}

func (x *Card) SetSuit(s Suit) {
	*x = (*x & 0b00001111) | Card(s)<<4
}

func CardC(s Suit, r Rank) Card {
	return Card(s)<<4 | Card(r)
}

func CardCD(s Suit, r Rank) Card {
	return Card(s)<<4 | Card(r) | 0b01000000
}

func GenCards() Cards {
	var cards = make([]Card, CardCountOfColor*4)
	for i := 0; i < CardCountOfColor; i++ {
		cards[i] = SpadesCards[i]
		cards[i+CardCountOfColor] = HeartsCards[i]
		cards[i+CardCountOfColor*2] = ClubsCards[i]
		cards[i+CardCountOfColor*3] = DiamondsCards[i]
	}
	return cards
}

func GenQRandCards() Cards {
	var cards = GenCards()
	QShuffle.Shuffle(cards)
	return cards
}

type Cards []Card

func (x Cards) Len() int {
	return len(x)
}

func (x Cards) Less(i, j int) bool {
	return x[i].Seq() < x[j].Seq()
}

func (x Cards) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x Cards) JTxt() string {
	var l = len(x)
	if l == 0 {
		return "[]"
	}
	var buf = bytes.NewBuffer(make([]byte, 768))
	buf.Reset()
	buf.WriteByte('[')
	for i := 0; i < l-1; i++ {
		buf.WriteByte('"')
		buf.WriteString(x[i].String())
		buf.WriteByte('"')
		buf.WriteByte(',')
	}
	buf.WriteByte('"')
	buf.WriteString(x[l-1].String())
	buf.WriteByte('"')
	buf.WriteByte(']')
	return buf.String()
}

func (x Cards) Txt() string {
	var l = len(x)
	if l == 0 {
		return ""
	}
	var buf = bytes.NewBuffer(make([]byte, 512))
	buf.Reset()
	for i := 0; i < l-1; i++ {
		buf.WriteString(x[i].String())
		buf.WriteByte(',')
	}
	buf.WriteString(x[l-1].String())
	return buf.String()
}

func init() {
	for i := 0; i < CardCountOfColor; i++ {
		SpadesCards[i] = CardC(Spades, Rank(i+1))
		HeartsCards[i] = CardC(Hearts, Rank(i+1))
		ClubsCards[i] = CardC(Clubs, Rank(i+1))
		DiamondsCards[i] = CardC(Diamonds, Rank(i+1))

		SpadesDownCards[i] = CardCD(Spades, Rank(i+1))
		HeartsDownCards[i] = CardCD(Hearts, Rank(i+1))
		ClubsDownCards[i] = CardCD(Clubs, Rank(i+1))
		DiamondsDownCards[i] = CardCD(Diamonds, Rank(i+1))
	}
}
