package sol

var (
	cardMap = map[string]Card{
		"As": CardC(Spades, 1),
		"2s": CardC(Spades, 2),
		"3s": CardC(Spades, 3),
		"4s": CardC(Spades, 4),
		"5s": CardC(Spades, 5),
		"6s": CardC(Spades, 6),
		"7s": CardC(Spades, 7),
		"8s": CardC(Spades, 8),
		"9s": CardC(Spades, 9),
		"Ts": CardC(Spades, 10),
		"Js": CardC(Spades, 11),
		"Qs": CardC(Spades, 12),
		"Ks": CardC(Spades, 13),

		"Ah": CardC(Hearts, 1),
		"2h": CardC(Hearts, 2),
		"3h": CardC(Hearts, 3),
		"4h": CardC(Hearts, 4),
		"5h": CardC(Hearts, 5),
		"6h": CardC(Hearts, 6),
		"7h": CardC(Hearts, 7),
		"8h": CardC(Hearts, 8),
		"9h": CardC(Hearts, 9),
		"Th": CardC(Hearts, 10),
		"Jh": CardC(Hearts, 11),
		"Qh": CardC(Hearts, 12),
		"Kh": CardC(Hearts, 13),

		"Ac": CardC(Clubs, 1),
		"2c": CardC(Clubs, 2),
		"3c": CardC(Clubs, 3),
		"4c": CardC(Clubs, 4),
		"5c": CardC(Clubs, 5),
		"6c": CardC(Clubs, 6),
		"7c": CardC(Clubs, 7),
		"8c": CardC(Clubs, 8),
		"9c": CardC(Clubs, 9),
		"Tc": CardC(Clubs, 10),
		"Jc": CardC(Clubs, 11),
		"Qc": CardC(Clubs, 12),
		"Kc": CardC(Clubs, 13),

		"Ad": CardC(Diamonds, 1),
		"2d": CardC(Diamonds, 2),
		"3d": CardC(Diamonds, 3),
		"4d": CardC(Diamonds, 4),
		"5d": CardC(Diamonds, 5),
		"6d": CardC(Diamonds, 6),
		"7d": CardC(Diamonds, 7),
		"8d": CardC(Diamonds, 8),
		"9d": CardC(Diamonds, 9),
		"Td": CardC(Diamonds, 10),
		"Jd": CardC(Diamonds, 11),
		"Qd": CardC(Diamonds, 12),
		"Kd": CardC(Diamonds, 13),
	}
)

func ConvertCardFromString(s string) Card {
	var c, ok = cardMap[s]
	if !ok {
		panic(s)
	}
	return c
}

func ConvertCardsFromStringList(ss []string) Cards {
	var count = len(ss)
	var cards = make(Cards, count)
	for i := 0; i < count; i++ {
		cards[i] = ConvertCardFromString(ss[i])
	}
	return cards
}
