package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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
		log.Println("TBD")
	} else {
		for _, filepath := range filepaths {
			file, err := os.Open(filepath)
			if err != nil {
				log.Fatal(err)
			}
			if options.printBytes {
				count := ByteCounter(filepath)
				printResults(count, filepath)
			}
			if options.printLines {
				count := LineCounter(file)
				printResults(count, filepath)
			}
			if options.printWords {
				count := WordCounter(file)
				printResults(count, filepath)
			}
		}
	}
}

func printResults(count int, filepath string) {
	fmt.Printf("%d %s\n", count, filepath)
}
