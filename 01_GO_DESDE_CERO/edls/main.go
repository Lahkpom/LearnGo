package main

import (
	"flag"
	"fmt"
)

func main() {
	//* creamos las flag
	//* -p
	//* las flag se inicializan así, y retornan un puntero

	flagPattern := flag.String("p", "", "Filter by pattern.")
	flagAll := flag.Bool("a", false, "All files including hide files.")
	flagNumbersRecord := flag.Int("n", 0, "Number of records.")

	//* order flags
	orderByTime := flag.Bool("t", false, "Sort by time, oldest first.")
	orderBySize := flag.Bool("s", false, "Sort by size, smallest first.")
	orderReverse := flag.Bool("r", false, "Reverse order while sorting.")

	//* Siempre usar el Parse para poder utilizar la variable
	flag.Parse()

	fmt.Println(*flagPattern)
	fmt.Println(*flagAll)
	fmt.Println(*flagNumbersRecord)
	fmt.Println(*orderByTime)
	fmt.Println(*orderBySize)
	fmt.Println(*orderReverse)

}
