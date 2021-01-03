package config

import (
	"os"
	"strconv"
)

func loadEnv() *Env {
	env := &Env{
		// default
		TimeLimit:     60,
		WordsFilePath: "../config/words/words.txt",
	}
	tls := os.Getenv("TIME_LIMIT")
	if tl, err := strconv.Atoi(tls); err == nil {
		env.TimeLimit = tl
	}
	if path := os.Getenv("WORDS_FILE_PATH"); len(path) != 0 {
		env.WordsFilePath = path
	}
	return env
}
