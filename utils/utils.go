package utils

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

func MustClose(readCloser io.ReadCloser) {
	if err := readCloser.Close(); err != nil {
		panic(err)
	}
}

func ReadLines(file string) (lines []string) {
	// Open the passed file relative to caller directory
	_, fl, _, _ := runtime.Caller(1)
	f, err := os.Open(path.Join(path.Dir(fl), file))
	defer MustClose(f)
	if err != nil {
		panic(err)
	}

	// Returns file as array of lines
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}

func ReadAll(file string) string {
	_, fl, _, _ := runtime.Caller(1)
	f, err := os.Open(path.Join(path.Dir(fl), file))
	defer MustClose(f)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	return strings.TrimSpace(string(data))
}
