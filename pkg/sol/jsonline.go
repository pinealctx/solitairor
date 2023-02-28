package sol

import (
	"os"
	"strings"
)

func ReadFileLine(filename string) ([]string, error) {
	var f, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var cs = strings.Split(string(f), "\n")
	return cs, nil
}
