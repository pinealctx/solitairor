package sol

import "sort"

const (
	PileCount = 7
)

type PileList [PileCount]Pile

func (p *PileList) Len() int {
	return PileCount
}

func (p *PileList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PileList) Less(i, j int) bool {
	return p[i].Less(&p[j])
}

func (p *PileList) Sort() {
	sort.Sort(p)
}

func (p *PileList) SetPileCard(i int, cards ...Card) {
	p[i].AddCards(cards)
}

func (p *PileList) SetPileCards(i int, cards []Card) {
	p[i].AddCards(cards)
}

func (p *PileList) Clone() PileList {
	var o PileList
	for i := range p {
		o[i] = *(p[i].Clone())
	}
	return o
}

func (p *PileList) Equals(o *PileList) bool {
	for i := range p {
		if !p[i].Equals(&o[i]) {
			return false
		}
	}
	return true
}
