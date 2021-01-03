package config

import (
	"bufio"
	"fmt"
	"os"
)

func loadWords(e *Env) (*Texts, error) {
	texts := &Texts{}
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(fmt.Sprintf("%s/%s", wd, e.WordsFilePath))
	if err != nil {
		return nil, err
	}
	lines := make([]string, 0, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	texts.Words = lines
	return texts, nil
}
