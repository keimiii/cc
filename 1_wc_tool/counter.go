package main

import (
	"bytes"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

func ByteCounter(filepath string) int {
	fi, err := os.Stat(filepath)
	if err != nil {
		log.Fatal(err)
	}
	return int(fi.Size())
}

func LineCounter(r io.Reader) int {
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

func WordCounter(r io.Reader) int {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, r)
	if err != nil {
		log.Fatal(err)
	}
	s := buf.String()
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(s, -1)
	return len(results)
}
