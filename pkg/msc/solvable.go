package pile

import (
	"fmt"
	"github.com/pinealctx/neptune/jsonx"
	"strings"
	"sync"
)

func ConvertKlondikes(sourceFile string) ([]*Converted, error) {
	var solvableList, err = ParseSolvableFileWithPrefix(sourceFile, "klondike")
	if err != nil {
		return nil, err
	}
	var convertedList = make([]*Converted, len(solvableList))
	for i, it := range solvableList {
		convertedList[i], err = it.ConvertSolitaire()
		if err != nil {
			return nil, err
		}
	}
	return convertedList, nil
}

func ParseSolvableFileWithPrefix(filename, prefix string) ([]*Solvable, error) {
	var list []*Solvable
	var err = jsonx.LoadJSONFile2Obj(filename, &list)
	if err != nil {
		return nil, fmt.Errorf("load json file error: %w", err)
	}
	var dist = make([]*Solvable, 0, len(list))
	for _, it := range list {
		if strings.HasPrefix(it.Name, prefix) {
			dist = append(dist, it)
		}
	}
	return dist, nil
}

type Solvable struct {
	Name       string   `json:"name"`
	Difficulty string   `json:"difficulty"`
	Seeds      []string `json:"seeds"`
}

func (s *Solvable) ConvertSolitaire() (*Converted, error) {
	var l = len(s.Seeds)
	var decks = make([][]byte, l)
	var errors = make([]error, l)
	var wg = sync.WaitGroup{}
	wg.Add(l)
	for i := 0; i < l; i++ {
		go func(index int) {
			defer wg.Done()
			decks[index], errors[index] = s.convertSolitaire(s.Seeds[index])
		}(i)
	}
	wg.Wait()
	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}
	return &Converted{
		Name:       s.Name,
		Difficulty: s.Difficulty,
		Decks:      decks,
	}, nil
}

func (s *Solvable) convertSolitaire(seed string) ([]byte, error) {
	var cards, err = Seed(seed)
	if err != nil {
		return nil, err
	}
	return cards.ToSol().ToBytes(), nil
}

type Converted struct {
	Name       string   `json:"name"`
	Difficulty string   `json:"difficulty"`
	Decks      [][]byte `json:"decks"` // sol card base64
}
