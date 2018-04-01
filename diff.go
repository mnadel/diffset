package main

import "fmt"

func diff(oldFilename, newFilename string) {
	oldLines := readFile(oldFilename)
	newLines := readFile(newFilename)

	numOldLines := int64(0)
	numNewLines := int64(0)

	for k := range oldLines {
		numOldLines++

		if _, ok := newLines[k]; !ok {
			fmt.Println("-", k)
		}
	}

	for k := range newLines {
		numNewLines++

		if _, ok := oldLines[k]; !ok {
			fmt.Println("+", k)
		}
	}
}
