package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func cat(reader io.Reader, line int, numberNonBlanks bool) int {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text := scanner.Text()

		if line < 1 {
			fmt.Println(text)
		} else {
			if numberNonBlanks {
				if text == "" {
					fmt.Print("\n")
					continue
				}
			}
			fmt.Printf("%d %s\n", line, text)
			line++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return line
}

func main() {
	line := -1
	numberLines := flag.Bool("n", false, "Number the output lines, starting at 1.")
	numberNonBlanks := flag.Bool("b", false, "Number the non-blank output lines, starting at 1.")

	flag.Parse()
	if *numberLines || *numberNonBlanks {
		line = 1
	}

	filenames := flag.CommandLine.Args()

	if len(filenames) == 0 || filenames[0] == "-" {
		cat(os.Stdin, line, *numberNonBlanks)
	} else {
		for _, filename := range filenames {
			file, err := os.Open(filename)

			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			line = cat(file, line, *numberNonBlanks)
		}
	}
}
