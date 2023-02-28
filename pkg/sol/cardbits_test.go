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

func TestCardBits_AddCard(t *testing.T) {
	var x CardBits
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))

	x.AddCard(SpadesCards[0])
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))

	x.RemoveCard(SpadesCards[0])
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))

	x.AddCard(SpadesCards[0], HeartsCards[0])
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))

	x.RemoveCard(SpadesCards[0], HeartsCards[0])
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))

	x.AddCard(SpadesCards[0], HeartsCards[0])
	x.RemoveCard(HeartsCards[0])
	t.Log(x)
	t.Log(x.Empty())
	t.Log(x.Has(SpadesCards[0]))
	t.Log(x.Has(HeartsCards[0]))
}
