package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	oldFilename = flag.String("old", "", "Old file path")
	newFilename = flag.String("new", "", "New file path")
	intersect   = flag.Bool("intersect", false, "Only show intersections")
)

func main() {
	flag.Parse()

	if flag.NArg() == 2 {
		*oldFilename = flag.Arg(0)
		*newFilename = flag.Arg(1)
	}

    if *oldFilename == "" || *newFilename == "" {
        flag.PrintDefaults()
        os.Exit(1)
    }

	oldFile, err := os.Open(*oldFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer oldFile.Close()

	newFile, err := os.Open(*newFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	if *intersect {
		diffIntersect()
	} else {
		diff()
	}
}

func diffIntersect() {
	oldLines := readHashes(*oldFilename)
	foundInNew := make(map[string]bool)

	f, err := os.Open(*newFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " \r\n")

		h := hash(text)

		if _, ok := oldLines[hash(text)]; ok {
			// only print common elements once
			if _, alreadySeen := foundInNew[h]; !alreadySeen {
				foundInNew[h] = true
				fmt.Println(text)
			}
		}
	}
}

func diff() {
	oldLines := readFile(*oldFilename)
	newLines := readFile(*newFilename)

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

func readFile(filename string) map[string]bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[string]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key := strings.Trim(scanner.Text(), " \r\n")
		m[key] = true
	}

	return m
}

func readHashes(filename string) map[interface{}]bool {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	m := make(map[interface{}]bool)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		key := strings.Trim(scanner.Text(), " \r\n")
		m[hash(key)] = true
	}

	return m
}

func hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
