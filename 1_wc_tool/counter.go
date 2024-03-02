package main

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

// TODO: use bufio.Scanner instead of os.Stat to parse file/text/standard input
func ReadStandardInput(scanner *bufio.Scanner, options Options) int {
	return 0
}

// TODO: use bufio.Scanner instead of os.Stat to parse file/text/standard input
func ReadFile(filepath string, options Options) int {
	return 0
}

func ByteCounter(filepath string) int {
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return int(fi.Size())
}

func LineCounter(filepath string) int {
	r, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			log.Fatal(err)
		}
	}
}

func WordCounter(filepath string) int {
	r, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, r)
	if err != nil {
		log.Fatal(err)
	}
	s := buf.String()
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(s, -1)
	return len(results)
}

func CharacterCounter(filepath string) int {
	r, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, r)
	if err != nil {
		log.Fatal(err)
	}
	s := buf.String()
	return utf8.RuneCountInString(s)
}
