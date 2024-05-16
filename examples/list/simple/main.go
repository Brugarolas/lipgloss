package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/list"
)

func main() {
	l := list.New(
		"A",
		"B",
		"C",
		list.New(
			"D",
			"E",
			"F",
		).Enumerator(list.Roman),
	)
	fmt.Println(l)
}
