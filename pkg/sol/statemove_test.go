package sol

import "testing"

func TestHandleMoveFromDeck1(t *testing.T) {
	var s = genSimpleState1()
	var ns = s.handleMoveFromDeck(nil)
	for _, v := range ns {
		t.Log(v)
	}
}

func TestHandleMoveToFound1(t *testing.T) {
	var s = genSimpleState1()
	var ns = s.handleMoveToFound(nil)
	for _, v := range ns {
		t.Log(v)
	}
}

func TestHandleMoveToOtherTab1(t *testing.T) {
	var s = genSimpleState1()
	var ns = s.handleMoveToOtherTab(nil)
	for _, v := range ns {
		t.Log(v)
	}
}

func TestHandleMoveFromDeck2(t *testing.T) {
	var s = genSimpleState2()
	var ns = s.handleMoveFromDeck(nil)
	for _, v := range ns {
		t.Log(v)
	}
}

func TestHandleMoveToFound2(t *testing.T) {
	var s = genSimpleState2()
	var ns = s.handleMoveToFound(nil)
	for _, v := range ns {
		t.Log(v)
	}
}

func TestHandleMoveToOtherTab2(t *testing.T) {
	var s = genSimpleState2()
	var ns = s.handleMoveToOtherTab(nil)
	for _, v := range ns {
		t.Log(v)
	}
}
