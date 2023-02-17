package sol

import "testing"

func TestCardBits_Set(t *testing.T) {
	var x CardBits
	x.Set(SpadesCards[0])
	x.Set(SpadesCards[3])

	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(SpadesCards[3]))
	t.Log(x.Has(SpadesCards[1]))

	x.Unset(SpadesCards[0])
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(SpadesCards[3]))
	t.Log(x.Has(SpadesCards[1]))

	x.Unset(SpadesCards[3])
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(SpadesCards[3]))
	t.Log(x.Has(SpadesCards[1]))
}
