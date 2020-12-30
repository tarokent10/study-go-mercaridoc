package config

import (
	"os"
	"strconv"
)

// Env os env info
type Env struct {
	timeLimit     int
	wordsFilePath string
}

func loadEnv() *Env {
	env := &Env{
		// default
		timeLimit:     60,
		wordsFilePath: "../config/words/words.txt",
	}
	tls := os.Getenv("TIME_LIMIT")
	if tl, err := strconv.Atoi(tls); err == nil {
		env.timeLimit = tl
	}
	if path := os.Getenv("WORDS_FILE_PATH"); len(path) != 0 {
		env.wordsFilePath = path
	}
	return env
}
