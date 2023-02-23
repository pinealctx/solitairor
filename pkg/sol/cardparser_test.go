package sol

import "testing"

func TestStringSlice(t *testing.T) {
	var ss = "123"
	t.Log(ss)
	t.Log(ss[1:])
	t.Log(ss[2:])
	t.Log(ss[3:])
}

func TestParseCard(t *testing.T) {
	t.Log(ParseCard("♠️"))
	t.Log(ParseCard("♥️️"))
	t.Log(ParseCard("♣️"))
	t.Log(ParseCard("♦️️"))

	t.Log(ParseCard("🫀A"))

	t.Log(ParseCard("♠️A"))
	t.Log(ParseCard("♥️J"))
	t.Log(ParseCard("♣️K"))
	t.Log(ParseCard("♦️Q"))

	t.Log(ParseCard("♠️1"))
	t.Log(ParseCard("♥️2"))
	t.Log(ParseCard("♣️3"))
	t.Log(ParseCard("♦️4"))
	t.Log(ParseCard("♠️5"))
	t.Log(ParseCard("♥️6"))
	t.Log(ParseCard("♣️7"))
	t.Log(ParseCard("♦️8"))
	t.Log(ParseCard("♦️9"))
	t.Log(ParseCard("♦️10"))
	t.Log(ParseCard("♠️10"))
}

func TestParseCards(t *testing.T) {
	var cs = "♣️10"
	t.Log(ParseCards(cs))
	t.Log([]rune(cs))
	cs = "♣️10,♦️7,♣️K,♣️4,♣️8,♥️7,♥️4,♠️5,♦️J,♦️8,♠️2,♥️J,♣️J,♣️7,♦️2,♥️2,♥️K,♦️3,♠️7,♥️8,♠️10,♠️Q,♣️Q,♠️9,♦️4,♦️6,♣️2,♠️A,♠️3,♥️9,♠️8,♦️K,♣️9,♦️10,♦️9,♥️A,♠️J,♣️5,♥️3,♥️5,♦️Q,♣️6,♠️6,♥️Q,♣️3,♦️A,♣️A,♥️6,♠️4,♠️K,♦️5,♥️10"
	t.Log(ParseCards(cs))
	cs = "♦️10,♣️4,♥️7,♣️2,♥️9,♣️9,♥️Q,♥️8,♣️10,♠️5,♥️2,♣️Q,♥️J,♦️4,♣️8,♣️K,♦️K,♣️3,♣️A,♦️8,♦️2,♥️5,♦️A,♥️10,♣️7,♦️Q,♣️5,♠️4,♠️A,♣️J,♥️K,♥️A,♦️9,♠️3,♠️7,♦️7,♥️3,♠️6,♠️Q,♦️5,♠️J,♠️2,♠️10,♣️6,♠️K,♥️4,♦️6,♦️3,♥️6,♠️8,♠️9,♦️J"
	t.Log(ParseCards(cs))
}
