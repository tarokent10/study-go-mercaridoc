package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/tarokent10/go-utils/imgutil"
)

// fconv converts file type.
// useage
// fconv [directory name] {-s [source type(default jpeg)] -d [dest type(default png)]}
func main() {
	args := new(argsMgr)
	args.parse()
	fmt.Printf("convert files in %s(%s -> %s)\n", args.dirname, args.srcType, args.destType)
	if err := convert(args.srcType, args.destType, args.dirname); err != nil {
		log.Fatal(err.Error())
	}
}

func convert(s, d filetype, dir string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if dir != path && info.IsDir() {
				// 自身もwalk対象となるのでチェックしないと無限ループになる
				return convert(s, d, path)
			}
			pos := strings.LastIndex(path, ".")
			if pos != -1 && path[pos:] == s.withExt() {
				return convertFile(s, d, path)
			}
			return nil
		})
	return err
}
func convertFile(s, d filetype, path string) error {
	var err error
	if s == fileTypePNG {
		switch d {
		case fileTypeJPEG:
			_, err = imgutil.Png2Jpeg(path)
		case fileTypeGIF:
			_, err = imgutil.Png2Gif(path)
		}
	} else if s == fileTypeJPEG {
		switch d {
		case fileTypePNG:
			_, err = imgutil.Jpeg2Png(path)
		case fileTypeGIF:
			_, err = imgutil.Jpeg2Gif(path)
		}
	} else if s == fileTypeGIF {
		switch d {
		case fileTypePNG:
			_, err = imgutil.Gif2Png(path)
		case fileTypeJPEG:
			_, err = imgutil.Gif2Jpeg(path)
		}
	}
	return err
}

type filetype string

func (f *filetype) is(s string) bool {
	return string(*f) == strings.ToLower(s)
}

func (f *filetype) withExt() string {
	return fmt.Sprintf(".%s", string(*f))
}

const (
	fileTypePNG  filetype = "png"
	fileTypeJPEG filetype = "jpeg"
	fileTypeGIF  filetype = "gif"
)

type argsMgr struct {
	dirname  string
	srcType  filetype
	destType filetype
}

func (m *argsMgr) parse() {
	// default
	src := fileTypeJPEG
	dest := fileTypePNG

	s := flag.String("s", "jpeg", "fconv -s jpeg")
	d := flag.String("d", "png", "fconv -d png")
	flag.Parse()
	args := flag.Args()

	if !m.validate(s, d, args) {
		log.Fatal("arguments are not valid!!")
	}
	if s != nil {
		src, _ = m.toFileType(*s)
	}
	if d != nil {
		dest, _ = m.toFileType(*d)
	}

	m.dirname = args[0]
	m.srcType = src
	m.destType = dest

}
func (m *argsMgr) validate(s, d *string, args []string) bool {
	// 冗長なのでもう少しスマートに書きたかった..
	if len(args) != 1 {
		println("no or too many args!")
		return false
	}

	if s != nil {
		if _, err := m.toFileType(*s); err != nil {
			println(err.Error())
			return false
		}
	}
	if d != nil {
		if _, err := m.toFileType(*d); err != nil {
			println(err.Error())
			return false
		}
	}

	return true
}

func (m *argsMgr) toFileType(s string) (filetype, error) {
	for _, ft := range []filetype{fileTypePNG, fileTypeJPEG, fileTypeGIF} {
		if ft.is(s) {
			return ft, nil
		}
	}
	return "", errors.New("no such file type")
}
