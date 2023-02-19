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
	tRunPuzzleAndLog(t, p, NewGameState(cards), r)
}

func tRunPuzzleAndLog(t *testing.T, p *Puzzle, state *StateM, r *Record) {
	p.InitRoot(state)
	p.Run()
	if len(p.hit) > 0 {
		p.Record(r)
	}
	t.Log(r)
}
