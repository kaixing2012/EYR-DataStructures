package main

import (
	"fmt"
	"golang/pkg/data_structures"
)

func main() {
	listOne := data_structures.LinkedList{}

	listOne.Insert("Edward")
	listOne.Insert("Herro")
	listOne.Insert("Freyja")
	fmt.Println(listOne.Delete("Freyja"))

	fmt.Println(listOne.Retrieve("Freyja"))
	fmt.Println(listOne.Size())
}
