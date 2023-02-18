package sol

import "testing"

func TestPuzzle_Run1(t *testing.T) {
	var p = NewPuzzle(10000, 10000)
	p.InitRoot(genSimpleState1())
	p.Run()
	t.Log(p.Road())
	t.Log(p.searchCount)
}

func TestPuzzle_Run2(t *testing.T) {
	var p = NewPuzzle(10000, 10000)
	p.InitRoot(genSimpleState2())
	p.Run()
	t.Log(p.Road())
	t.Log(p.searchCount)
}
