package pile

import (
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

var (
	CardSeed = decimal.NewFromInt32(1)
)

func init() {
	for i := int32(1); i <= CardCount; i++ {
		CardSeed = CardSeed.Mul(decimal.NewFromInt32(i))
	}
}

func Seed(seed string) (Cards, error) {
	seed = strings.TrimRight(seed, "\\r")
	var e = make([]int, CardCount)
	var n = decimal.Zero
	var r = decimal.NewFromInt(CardCount)
	var o = CardSeed
	var t, err = decimal.NewFromString(seed)
	if err != nil {
		return nil, fmt.Errorf("invalid seed: %w", err)
	}
	t = t.Mod(o)
	for a := CardCount; a > 0; a-- {
		n = t.Mod(r)
		var i = n.IntPart()
		e[a-1] = int(i)
		t = t.Div(r)
		r = r.Sub(decimal.NewFromInt(1))
	}
	var cards = NewFreshCards()
	for a := CardCount - 1; a >= 1; a-- {
		if a != e[a] {
			cards[a], cards[e[a]] = cards[e[a]], cards[a]
		}
	}
	return cards, nil
}
