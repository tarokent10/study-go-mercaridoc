package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	n := flag.Bool("n", false, "set a number")
	flag.Parse()
	args := flag.Args()

	c := newCatCommand(*n)
	c.cat(args)
}

type catCommand struct {
	setnum bool
	n      int32
}

func (c *catCommand) cat(files []string) {
	for _, fname := range files {
		if err := c.read(fname); err != nil {
			log.Fatal(err.Error())
		}
	}
}
func (c *catCommand) read(fname string) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer func() {
		f.Close() // not handle error
	}()

	s := bufio.NewScanner(f)
	buf := bytes.Buffer{}
	for s.Scan() {
		c.n++
		fmt.Fprint(&buf, c.sprintf(s.Text()))
	}
	if err := s.Err(); err != nil {
		return err
	}
	fmt.Fprint(os.Stdout, buf.String())
	return nil
}
func (c *catCommand) sprintf(text string) string {
	if c.setnum {
		return fmt.Sprintf("%d: %s\n", c.n, text)
	} else {
		return fmt.Sprintf("%s\n", text)
	}
}

func newCatCommand(n bool) catCommand {
	return catCommand{n, 0}
}
