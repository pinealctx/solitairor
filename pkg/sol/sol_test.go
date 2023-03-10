package sol

import "testing"

func TestPuzzle_Run1(t *testing.T) {
	var p = NewPuzzle(10000, 10000)
	tRunPuzzleAndLog(t, p, genSimpleState1(), &Record{})
}

func TestPuzzle_Run2(t *testing.T) {
	var p = NewPuzzle(10000, 10000)
	tRunPuzzleAndLog(t, p, genSimpleState2(), &Record{})
}

func TestPuzzle_RunRand(t *testing.T) {
	var p = NewPuzzle(100000, 100000)
	var cards = GenQRandCards()
	var r = &Record{InitCards: cards.Txt()}
	tRunPuzzleAndLog(t, p, NewGameStateFromLegacyCards(cards), r)
}

func TestShunCardSolver(t *testing.T) {
	var cs = GenCards()
	var bs = cs.ToBytes()
	testCardsSolverWithDebug(t, bs, 3012, true)
}

func TestPartSolver(t *testing.T) {
	var st = NewState()
	//deck:♠️9,♠️10,♠️J,♠️Q,♠️K,
	st.StockCardBits.AddCard(SpadesCards[9-1],
		SpadesCards[10-1], SpadesCards[JackRank-1], SpadesCards[QueenRank-1], SpadesCards[KingRank-1])
	//pile1:
	//pile2:♥️K*,♣️6,
	st.PileTable[2-1].AddCard(HeartsCards[KingRank-1].FaceDownCard(), ClubsCards[6-1])
	//pile3:♦️K,♣️Q,♦️J,
	st.PileTable[3-1].AddCard(DiamondsCards[KingRank-1], ClubsCards[QueenRank-1], DiamondsCards[JackRank-1])
	//pile4:♣️2*,♣️8*,♣️K*,♦️4,
	st.PileTable[4-1].AddCard(ClubsCards[2-1].FaceDownCard(), ClubsCards[8-1].FaceDownCard(),
		ClubsCards[KingRank-1].FaceDownCard(), DiamondsCards[4-1])
	//pile5:♣️4*,♣️10*,♦️2*,♦️6*,♦️9,
	st.PileTable[5-1].AddCard(ClubsCards[4-1].FaceDownCard(), ClubsCards[10-1].FaceDownCard(),
		DiamondsCards[2-1].FaceDownCard(), DiamondsCards[6-1].FaceDownCard(), DiamondsCards[9-1])
	//pile6:♣️3*,♣️9*,♦️A*,♦️5*,♦️8,♣️7,
	st.PileTable[6-1].AddCard(ClubsCards[3-1].FaceDownCard(), ClubsCards[9-1].FaceDownCard(),
		DiamondsCards[AceRank-1].FaceDownCard(), DiamondsCards[5-1].FaceDownCard(), DiamondsCards[8-1], ClubsCards[7-1])
	//pile7:♣️5*,♣️J*,♦️3*,♦️7*,♦️10*,♦️Q,
	st.PileTable[7-1].AddCard(ClubsCards[5-1].FaceDownCard(), ClubsCards[JackRank-1].FaceDownCard(),
		DiamondsCards[3-1].FaceDownCard(), DiamondsCards[7-1].FaceDownCard(), DiamondsCards[10-1].FaceDownCard(),
		DiamondsCards[QueenRank-1])
	testSolver(t, 100000000, st, "", true)
}

func TestSimpleCardSolver(t *testing.T) {
	var bs = []byte{122, 100, 87, 98, 89, 105, 92, 88, 106, 69, 82, 108, 91, 116, 104, 109, 125, 99, 97, 120, 114, 85, 113, 90, 39, 124, 68, 93, 71, 117, 84, 37, 65, 81, 119, 75, 118, 43, 121, 83, 66, 115, 3, 70, 74, 86, 12, 102, 72, 13, 73, 59}
	testCardsSolver(t, bs, 10000000)
	testCardsSolver(t, bs, 1000000)
	testCardsSolver(t, bs, 100000)
	testCardsSolver(t, bs, 10000)
	testCardsSolver(t, bs, 1000)
}

func TestMixCardSolver1(t *testing.T) {
	var bs = []byte{106, 119, 109, 100, 104, 87, 84, 69, 123, 120, 66, 91, 107, 103, 114, 82, 93, 115, 71, 88, 74, 76, 108, 73, 52, 118, 65, 72, 121, 85, 113, 34, 67, 125, 81, 124, 97, 25, 105, 75, 102, 86, 58, 101, 70, 68, 19, 92, 77, 35, 117, 26}
	testCardsSolver(t, bs, 10000000)
	testCardsSolver(t, bs, 1000000)
	testCardsSolver(t, bs, 100000)
	testCardsSolver(t, bs, 10000)
	testCardsSolver(t, bs, 1000)
}

func TestMixCardSolver2(t *testing.T) {
	var bs = []byte{106, 119, 109, 100, 104, 87, 84, 69, 123, 120, 66, 91, 107, 103, 114, 82, 93, 115, 71, 88, 74, 76, 108, 73, 52, 118, 65, 72, 121, 85, 113, 34, 67, 125, 81, 124, 97, 25, 105, 75, 102, 86, 58, 101, 70, 68, 19, 92, 77, 35, 117, 26}
	var cs = MakeCardsFromBytes(bs)
	t.Log(cs)
	var st = NewGameStateFromCards(cs)
	var p = NewPuzzle(100000, 100000000)
	p.InitRoot(st)
	var s1 = NewState()
	s1.PileTable[2].AddCard(DiamondsCards[9-1])

	s1.PileTable[3].AddCard(
		SpadesCards[8-1],
		DiamondsCards[7-1],
		ClubsCards[6-1],
		HeartsCards[5-1],
	)

	s1.PileTable[4].AddCard(
		ClubsCards[KingRank-1],
		HeartsCards[QueenRank-1],
		ClubsCards[JackRank-1],
		HeartsCards[10-1],
		ClubsCards[9-1],
		DiamondsCards[8-1],
		ClubsCards[7-1],
		HeartsCards[6-1],
		ClubsCards[5-1],
		HeartsCards[4-1],
	)

	s1.PileTable[5].AddCard(
		SpadesCards[KingRank-1],
		DiamondsCards[QueenRank-1],
		SpadesCards[JackRank-1],
		DiamondsCards[10-1],
		SpadesCards[9-1],
		HeartsCards[8-1],
		SpadesCards[7-1],
		DiamondsCards[6-1],
		SpadesCards[5-1],
		DiamondsCards[4-1],
		SpadesCards[3-1],
	)

	s1.PileTable[6].AddCard(
		DiamondsCards[KingRank-1],
		ClubsCards[QueenRank-1],
		HeartsCards[JackRank-1],
		ClubsCards[10-1],
		HeartsCards[9-1],
		ClubsCards[8-1],
		HeartsCards[7-1],
		SpadesCards[6-1],
		DiamondsCards[5-1],
		SpadesCards[4-1],
		HeartsCards[3-1],
	)
	s1.PileTable.Sort()
	p.findFunc = func(o *StateM) bool {
		return s1.SamePiles(o)
	}

	p.Run()
	var r = &Record{InitCards: cs.Txt()}
	p.Record(r)
	t.Log(r)
}

func testCardsSolver(t *testing.T, bs []byte, maxSearchSize int) {
	testCardsSolverWithDebug(t, bs, maxSearchSize, false)
}

func testCardsSolverWithDebug(t *testing.T, bs []byte, maxSearchSize int, debugMode bool) {
	var cs = MakeCardsFromBytes(bs)
	t.Log(cs)
	var st = NewGameStateFromCards(cs)
	testSolver(t, maxSearchSize, st, cs.Txt(), debugMode)
}

func testSolver(t *testing.T, maxSearchSize int, st *StateM, cardsTxt string, debugMode bool) {
	var p = NewPuzzle(10000000, maxSearchSize)
	p.debugMode = debugMode
	p.InitRoot(st)
	p.Run()
	var r = &Record{InitCards: cardsTxt}
	p.Record(r)
	t.Log(r)
}

func tRunPuzzleAndLog(t *testing.T, p *Puzzle, state *StateM, r *Record) {
	p.InitRoot(state)
	p.Run()
	if len(p.hit) > 0 {
		p.Record(r)
	}
	t.Log(r)
}
