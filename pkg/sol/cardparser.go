package sol

import (
	"fmt"
	"strings"
)

const (
	SpadesStr   = "♠️"
	HeartsStr   = "♥️"
	ClubsStr    = "♣️"
	DiamondsStr = "♦️"

	SpadesStrLen   = len(SpadesStr)
	HeartsStrLen   = len(HeartsStr)
	ClubsStrLen    = len(ClubsStr)
	DiamondsStrLen = len(DiamondsStr)
)

func ParseCards(ss string) (Cards, error) {
	var ts = strings.Split(ss, ",")
	var size = len(ts)
	if size == 0 {
		return nil, nil
	}
	var cards = make([]Card, size)
	var err error
	for i := 0; i < size; i++ {
		cards[i], err = ParseCard(ts[i])
		if err != nil {
			return nil, err
		}
	}
	return cards, nil
}

func ParseCard(ss string) (Card, error) {
	var suit, rankStr, err = splitCardStr(ss)
	if err != nil {
		return 0, err
	}
	var rank Rank
	rank, err = parseRank(rankStr)
	if err != nil {
		return 0, err
	}

	return CardC(suit, rank), nil
}

func ConvertLegacyCards(cards Cards) Cards {
	var stateM = NewGameStateFromLegacyCards(cards)
	var size = len(cards)
	var clientCards = make([]Card, size)

	for i := 0; i < DeckCount; i++ {
		clientCards[i] = cards[i]
		clientCards[i].SetFaceDown()
	}
	var index = DeckCount

	for j := 0; j < PileCount; j++ {
		for i := j; i < PileCount; i++ {
			clientCards[index] = stateM.PileTable[i].Cards[j]
			index++
		}
	}
	return clientCards
}

func splitCardStr(s string) (Suit, string, error) {
	if strings.HasPrefix(s, SpadesStr) {
		return Spades, s[SpadesStrLen:], nil
	}
	if strings.HasPrefix(s, HeartsStr) {
		return Hearts, s[HeartsStrLen:], nil
	}
	if strings.HasPrefix(s, ClubsStr) {
		return Clubs, s[ClubsStrLen:], nil
	}
	if strings.HasPrefix(s, DiamondsStr) {
		return Diamonds, s[DiamondsStrLen:], nil
	}
	return 0, "", fmt.Errorf("invalid suit: %+v", s)
}

func parseRank(s string) (Rank, error) {
	switch s {
	case "A":
		return AceRank, nil
	case "2":
		return Rank(2), nil
	case "3":
		return Rank(3), nil
	case "4":
		return Rank(4), nil
	case "5":
		return Rank(5), nil
	case "6":
		return Rank(6), nil
	case "7":
		return Rank(7), nil
	case "8":
		return Rank(8), nil
	case "9":
		return Rank(9), nil
	case "10":
		return Rank(10), nil
	case "J":
		return JackRank, nil
	case "Q":
		return QueenRank, nil
	case "K":
		return KingRank, nil
	default:
		return 0, fmt.Errorf("invalid rank: %+v", s)
	}
}
