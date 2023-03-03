package pile

import (
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"strconv"
	"strings"
)

const (
	CardCount = 52
)

type Suit uint

const (
	SuitClub Suit = iota + 1
	SuitDiamond
	SuitHeart
	SuitSpade
)

func NewSuit(suit sol.Suit) Suit {
	switch suit {
	case sol.Spades:
		return SuitSpade
	case sol.Diamonds:
		return SuitDiamond
	case sol.Hearts:
		return SuitHeart
	case sol.Clubs:
		return SuitClub
	default:
		return 255
	}
}

var (
	SuitOrder = []Suit{SuitSpade, SuitHeart, SuitClub, SuitDiamond}
)

func (s Suit) String() string {
	switch s {
	case SuitSpade:
		return "♠"
	case SuitHeart:
		return "♥"
	case SuitClub:
		return "♣"
	case SuitDiamond:
		return "♦"
	default:
		return "*"
	}
}

func (s Suit) ToSol() sol.Suit {
	switch s {
	case SuitSpade:
		return sol.Spades
	case SuitHeart:
		return sol.Hearts
	case SuitClub:
		return sol.Clubs
	case SuitDiamond:
		return sol.Diamonds
	default:
		return 255
	}
}

type Card struct {
	Suit  Suit
	Value byte
}

func NewCard(card sol.Card) *Card {
	return &Card{
		Suit:  NewSuit(card.Suit()),
		Value: byte(card.Rank()),
	}
}

func (c *Card) String() string {
	var v = strconv.Itoa(int(c.Value))
	if len(v) == 1 {
		v += " "
	}
	return fmt.Sprintf("%s%s", c.Suit.String(), v)
}

func (c *Card) ToSol() sol.Card {
	return sol.CardC(c.Suit.ToSol(), sol.Rank(c.Value))
}

type Cards []*Card

func NewCards(cards sol.Cards) Cards {
	var dist = make(Cards, len(cards))
	for i, it := range cards {
		dist[i] = NewCard(it)
	}
	return dist
}

func (c Cards) String() string {
	var b strings.Builder
	for _, cc := range c {
		b.WriteString(cc.String())
		b.WriteRune(' ')
	}
	return b.String()
}

func (c Cards) Reverse() Cards {
	var l = len(c)
	var dist = make(Cards, l)
	for i, j := 0, l-1; i < l; i, j = i+1, j-1 {
		dist[i] = c[j]
	}
	return dist
}

func (c Cards) ShowKlondikes() string {
	var cards = c.Reverse()
	var b strings.Builder
	var index = -1
	for i := 0; i < 7; i++ {
		b.WriteRune('\n')
		for ii := 0; ii < i; ii++ {
			b.WriteString("   ")
			b.WriteRune(' ')
		}
		for j := 0; j < 7-i; j++ {
			index++
			b.WriteString(cards[index].String())
			b.WriteRune(' ')
		}
	}
	index++
	b.WriteRune('\n')
	for ; index < CardCount; index++ {
		b.WriteString(cards[index].String())
		b.WriteRune(' ')
	}
	return b.String()
}

func (c Cards) ShowTriPeaks() string {
	var cards = c.Reverse()
	var b strings.Builder
	var maxIndex = 18
	var layer1 = []int{3, 9, 15}
	var layer2 = []int{2, 4, 8, 10, 14, 16}
	var layer3 = []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	var layer4 = []int{0, 2, 4, 6, 8, 10, 12, 14, 16, 18}
	var layers = [][]int{layer1, layer2, layer3, layer4}

	var cardIndex = -1
	for _, layer := range layers {
		b.WriteRune('\n')
		for i := 0; i <= maxIndex; i++ {
			var found = false
			for _, index := range layer {
				if i == index {
					found = true
					break
				}
			}
			if found {
				cardIndex++
				b.WriteString(cards[cardIndex].String())
				b.WriteRune(' ')
			} else {
				b.WriteString("   ")
				b.WriteRune(' ')
			}
		}
	}
	cardIndex++
	b.WriteRune('\n')
	for ; cardIndex < CardCount; cardIndex++ {
		b.WriteString(cards[cardIndex].String())
		b.WriteRune(' ')
	}
	return b.String()
}

func (c Cards) ToSol() sol.Cards {
	var dist = make(sol.Cards, len(c))
	for i, it := range c {
		dist[i] = it.ToSol()
	}
	return dist
}

func NewFreshCards() Cards {
	var cards = make(Cards, 0, CardCount)
	for _, suit := range SuitOrder {
		for i := byte(1); i <= 13; i++ {
			cards = append(cards, &Card{
				Suit:  suit,
				Value: i,
			})
		}
	}
	return cards
}
