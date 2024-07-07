package main

import (
	"flag"
	"os"
	"regexp"
	"strings"
)

func main() {
	//* creamos las flag
	//* -p
	//* las flag se inicializan así, y retornan un puntero

	flagPattern := flag.String("p", "", "Filter by pattern.")
	flagAll := flag.Bool("a", false, "All files including hide files.")
	flagNumbersRecords := flag.Int("n", 0, "Number of records.")

	//* order flags
	// orderByTime := flag.Bool("t", false, "Sort by time, oldest first.")
	// orderBySize := flag.Bool("s", false, "Sort by size, smallest first.")
	// orderReverse := flag.Bool("r", false, "Reverse order while sorting.")

	//* Siempre usar el Parse para poder utilizar la variable
	flag.Parse()

	path := flag.Arg(0)

	if path == "" {
		path = "."
	}

	dirs, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	fs := []file{}
	for _, dir := range dirs {
		isHidden := isHidden(dir.Name(), path)

		if isHidden && !*flagAll {
			continue
		}

		if *flagPattern != "" {
			isMatched, err := regexp.MatchString("(?i)"+*flagPattern, dir.Name())
			if err != nil {
				panic(err)
			}

			if isMatched {
				continue
			}
		}

		f, err := getFile(dir, isHidden)
		if err != nil {
			panic(err)
		}

		fs = append(fs, f)
	}

	if *flagNumbersRecords == 0 || *flagNumbersRecords > len(fs) {
		*flagNumbersRecords = len(fs)
	}

	printList(fs, *flagNumbersRecords)
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, zip) ||
		strings.HasSuffix(f.name, gz) ||
		strings.HasSuffix(f.name, tar) ||
		strings.HasSuffix(f.name, rar) ||
		strings.HasSuffix(f.name, deb)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, png) ||
		strings.HasSuffix(f.name, jpg) ||
		strings.HasSuffix(f.name, gif)
}

func isHidden(fileName, basePath string) bool {
	return strings.HasPrefix(fileName, ".")
}
