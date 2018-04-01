package main

import "fmt"

func diffIntersect(oldFilename, newFilename string) {
	oldLines := readHashes(oldFilename)
	foundInNew := newSet()

	scanLines(newFilename, func(line string) {
		h := hash(line)

		if _, ok := oldLines[h]; ok {
			if foundInNew.add(h) {
				fmt.Println(line)
			}
		}
	})
}
