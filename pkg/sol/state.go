package sol

const (
	InfiniteStep = -1
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
		if s.PileTable[i].Len() > 0 {
			return false
		}
	}
	return true
}

func (s *StateM) ReverseBroadcast() {
	var p = s
	if s.ReverseStep == InfiniteStep {
		return
	}

	// reverse broadcast
	for {
		var pp = p.Parent
		if pp == nil {
			break
		}
		pp.ReverseStep = p.ReverseStep + 1
		p = pp
	}
}
