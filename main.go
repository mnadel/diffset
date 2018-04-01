package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"log"
	"os"
	"strings"
)

var (
	oldFilename = flag.String("old", "", "Old file path")
	newFilename = flag.String("new", "", "New file path")
	intersect   = flag.Bool("intersect", false, "Only show intersections")
)

type visitor func(string)

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

	if *intersect {
		diffIntersect(*oldFilename, *newFilename)
	} else {
		diff(*oldFilename, *newFilename)
	}
}

func scanLines(filename string, visit visitor) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := trim(scanner.Text())
		visit(text)
	}
}

func readHashes(filename string) map[string]bool {
	m := make(map[string]bool)

	scanLines(filename, func(line string) {
		m[hash(line)] = true
	})

	return m
}

func readFile(filename string) map[string]bool {
	m := make(map[string]bool)

	scanLines(filename, func(line string) {
		m[line] = true
	})

	return m
}

func trim(text string) string {
	return strings.Trim(text, " \r\n")
}

func hash(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	hashed := h.Sum(nil)

	return hex.EncodeToString(hashed)
}
