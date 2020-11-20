package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// ファイル学習用.ファイル名を指定して作成.
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("作成するファイル名を入力してください\n > ")
	for scanner.Scan() {
		filename := scanner.Text()
		if len(filename) > 0 {
			createFile(filename)
			break
		} else {
			fmt.Printf("入力が不正です\n")
		}
	}
}

func createFile(name string) {
	filepath := filepath.Join(name)
	fmt.Printf("outputfile: %s/%s\n", getCurrent(), filepath)
	df, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		if err := df.Close(); err != nil {
			log.Fatal(err.Error())
		}
	}()
	if _, err := io.WriteString(df, "go!go!"); err != nil {
		log.Fatal(err.Error())
	}
}

func getCurrent() string {
	var err error
	var current string
	if current, err = os.Getwd(); err != nil {
		log.Fatal(err.Error())
	}
	return current
}
