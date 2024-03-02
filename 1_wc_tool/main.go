package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	var commandLineOptions Options

	flag.BoolVar(&commandLineOptions.printBytes, "c", false, "Count bytes")
	flag.BoolVar(&commandLineOptions.printLines, "l", false, "Count lines")
	flag.BoolVar(&commandLineOptions.printWords, "w", false, "Count words")
	flag.BoolVar(&commandLineOptions.printChars, "m", false, "Count characters")
	flag.Parse()

	filepaths := flag.CommandLine.Args()

	run(filepaths, commandLineOptions)
}

func run(filepaths []string, options Options) {
	if len(filepaths) == 0 {
		// TODO: read from standard input if no filename is specified
		log.Println("TBD")
	} else {
		for _, filepath := range filepaths {
			if options.printBytes {
				count := ByteCounter(filepath)
				printResults(count, filepath)
			}
			if options.printLines {
				count := LineCounter(filepath)
				printResults(count, filepath)
			}
			if options.printWords {
				count := WordCounter(filepath)
				printResults(count, filepath)
			}
			if options.printChars {
				count := CharacterCounter(filepath)
				printResults(count, filepath)
			}
			if !options.printBytes && !options.printLines && !options.printWords && !options.printChars {
				lineCount := LineCounter(filepath)
				wordCount := WordCounter(filepath)
				byteCount := ByteCounter(filepath)
				printResultsMultiple(lineCount, wordCount, byteCount, filepath)
			}
		}
	}
}

func printResults(count int, filepath string) {
	fmt.Printf("%d %s\n", count, filepath)
}

func printResultsMultiple(lineCount int, wordCount int, byteCount int, filepath string) {
	fmt.Printf("%d %d %d %s\n", lineCount, wordCount, byteCount, filepath)
}
