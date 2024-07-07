package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

func main() {
	//* creamos las flag
	//* -p
	//* las flag se inicializan así, y retornan un puntero

	flagPattern := flag.String("p", "", "Filter by pattern.")
	flagAll := flag.Bool("a", false, "All files including hide files.")
	flagNumbersRecords := flag.Int("n", 0, "Number of records.")

	//* order flags
	hasOrderByTime := flag.Bool("t", false, "Sort by time, oldest first.")
	hasOrderBySize := flag.Bool("s", false, "Sort by size, smallest first.")
	hasOrderReverse := flag.Bool("r", false, "Reverse order while sorting.")

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

	if !*hasOrderBySize || !*hasOrderByTime {
		orderByName(fs, *hasOrderReverse)
	}

	if *hasOrderBySize && !*hasOrderByTime {
		orderBySize(fs, *hasOrderReverse)
	}

	if *hasOrderByTime && !*hasOrderBySize {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumbersRecords == 0 || *flagNumbersRecords > len(fs) {
		*flagNumbersRecords = len(fs)
	}

	printList(fs, *flagNumbersRecords)
}

func mySort[T string | int64](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort[string](strings.ToLower(files[i].name), strings.ToLower(files[j].name), isReverse)
	})
}
func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort[int64](files[i].modificationTime.Unix(), files[j].modificationTime.Unix(), isReverse)
	})
}

func orderBySize(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort[int64](files[i].size, files[j].size, isReverse)
	})
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]

		fmt.Printf("%s %s %s %10d %s %s %s%s\n", file.mode, file.userName, file.
			groupName, file.size, file.modificationTime.Format(time.DateTime), style.
			icon, file.name, style.symbol)
	}
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, err
	}
	return file
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
