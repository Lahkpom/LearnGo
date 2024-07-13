package main

import "fmt"

/*
type Course struct {
	Name string
	Price float64
	IsFree bool
	UserIDs []uint
	Classes map[uint]string
}
*/

func (c Course) PrintClasses() {
	fmt.Println("Las clases de este curso son:")
	for _, class := range c.Classes {
		fmt.Println(class)
	}
}
