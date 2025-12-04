package utils

import (
	//"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func SelectMostRecentSegment() (os.DirEntry, int) {
	entries, err := os.ReadDir("database")

	var MostRecentSegmentNumber int
	var segment os.DirEntry

	if err != nil {
		log.Fatal(err)
	}

	if len(entries) == 0 {
		return segment, 0
	}

	for _, entry := range entries {
		info, err := entry.Info()

		if err != nil {
			log.Fatal(err)
			continue
		}

		segmentNumber, _ := strconv.Atoi(strings.Replace(strings.Split(info.Name(), "-")[1], ".txt", "", 1))

		if segmentNumber > MostRecentSegmentNumber {
			MostRecentSegmentNumber = segmentNumber
			segment = entry
		}
	}

	return segment, MostRecentSegmentNumber
}

func CreateSegment() (*os.File, string, error) {
	_, segmentNumber := SelectMostRecentSegment()

	var segmentName strings.Builder
	segmentName.WriteString("database/segment-")
	segmentName.WriteString(strconv.Itoa(segmentNumber + 1))
	segmentName.WriteString(".txt")

	file, err := os.Create(segmentName.String())

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return file, file.Name(), err
}
