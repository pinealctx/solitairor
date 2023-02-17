package sol

// CardBits is a bitset of cards.
// now from bit 0-51
type CardBits uint64

// Set : set bit
func (x *CardBits) Set(c Card) *CardBits {
	*x |= 1 << c.Seq()
	return x
}

func (x *CardBits) Unset(c Card) *CardBits {
	*x &= ^(1 << c.Seq())
	return x
}

func (x *CardBits) AddCard(cards ...Card) *CardBits {
	for _, c := range cards {
		x.Set(c)
	}
	return x
}

func (x *CardBits) Has(c Card) bool {
	return *x&(1<<c.Seq()) != 0
}

func (x *CardBits) RemoveCard(cards ...Card) *CardBits {
	for _, c := range cards {
		x.Unset(c)
	}
	return x
}
