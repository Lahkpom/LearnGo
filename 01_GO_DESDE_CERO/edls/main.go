package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"time"

	// Devuelve información sin importar el SO en el que se ejectue
	"github.com/AJRDRGZ/fileinfo"
	"github.com/Lahkpom/mainModules/constraints"
	"github.com/fatih/color"
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

			if !isMatched {
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

	if *hasOrderByTime {
		orderByTime(fs, *hasOrderReverse)
	}

	if *flagNumbersRecords == 0 || *flagNumbersRecords > len(fs) {
		*flagNumbersRecords = len(fs)
	}

	printList(fs, *flagNumbersRecords)
}

func getFile(dir fs.DirEntry, isHidden bool) (file, error) {
	info, err := dir.Info()
	if err != nil {
		return file{}, fmt.Errorf("dir.Infor(): %v", err)
	}

	userName, groupName := fileinfo.GetUserAndGroup(info.Sys())

	f := file{
		name:             dir.Name(),
		isDir:            dir.IsDir(),
		isHidden:         isHidden,
		userName:         userName,
		groupName:        groupName,
		size:             info.Size(),
		modificationTime: info.ModTime(),
		mode:             info.Mode().String(),
	}

	setFile(&f)

	return f, nil
}

func mySort[T constraints.Ordered](i, j T, isReverse bool) bool {
	if isReverse {
		return i > j
	}
	return i < j
}

func orderByName(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort[string](
			strings.ToLower(files[i].name),
			strings.ToLower(files[j].name),
			isReverse,
		)
	})
}

func orderBySize(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].size,
			files[j].size,
			isReverse,
		)
	})
}

func orderByTime(files []file, isReverse bool) {
	sort.SliceStable(files, func(i, j int) bool {
		return mySort(
			files[i].modificationTime.Unix(),
			files[j].modificationTime.Unix(),
			isReverse,
		)
	})
}

func printList(fs []file, nRecords int) {
	for _, file := range fs[:nRecords] {
		style := mapStyleByFileType[file.fileType]

		fmt.Printf("%11s %-8s %-8s %10d %s %s %s%s %s\n",
			file.mode,
			file.userName,
			file.groupName,
			file.size,
			file.modificationTime.Format(time.DateTime),
			style.icon,
			setColor(file.name, style.color),
			style.symbol,
			markHidden(file.isHidden),
		)
	}
}

func setFile(f *file) {
	switch {
	case isLink(*f):
		f.fileType = FILE_LINK
	case f.isDir:
		f.fileType = FILE_DIRECTORY
	case isExec(*f):
		f.fileType = FILE_EXECUTABLE
	case isCompress(*f):
		f.fileType = FILE_COMPRESS
	case isImage(*f):
		f.fileType = FILE_IMAGE
	default:
		f.fileType = FILE_REGULAR
	}
}

func setColor(nameFile string, styleColor color.Attribute) string {
	switch styleColor {
	case color.FgBlue:
		return blue(nameFile)
	case color.FgGreen:
		return green(nameFile)
	case color.FgRed:
		return red(nameFile)
	case color.FgMagenta:
		return magenta(nameFile)
	case color.FgCyan:
		return cyan(nameFile)
	default:
		return nameFile
	}
}

func isLink(f file) bool {
	return strings.HasPrefix(strings.ToLower(f.mode), "l")
}

func isExec(f file) bool {
	if runtime.GOOS == Windows {
		return strings.HasSuffix(f.name, EXE)
	}
	return strings.Contains(f.mode, "x")
}

func isCompress(f file) bool {
	return strings.HasSuffix(f.name, ZIP) ||
		strings.HasSuffix(f.name, GZ) ||
		strings.HasSuffix(f.name, TAR) ||
		strings.HasSuffix(f.name, RAR) ||
		strings.HasSuffix(f.name, DEB)
}

func isImage(f file) bool {
	return strings.HasSuffix(f.name, PNG) ||
		strings.HasSuffix(f.name, JPG) ||
		strings.HasSuffix(f.name, GIF)
}

func isHidden(fileName, basePath string) bool {
	filePath := fileName

	if runtime.GOOS == Windows {
		filePath = path.Join(basePath, filePath)
	}

	return fileinfo.IsHidden(filePath)
}

func markHidden(isHidden bool) string {
	if !isHidden {
		return ""
	}

	return yellow("ø")
}
