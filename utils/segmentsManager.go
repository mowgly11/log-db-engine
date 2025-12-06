package utils

import (
	//"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const SEGEMENT_SIZE_LIMIT_KB int32 = 1

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
	segmentName.WriteString("database\\segment-")
	segmentName.WriteString(strconv.Itoa(segmentNumber + 1))
	segmentName.WriteString(".txt")

	file, err := os.Create(segmentName.String())

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return file, file.Name(), err
}

func CreateOrSelectSegment() string {
	var entryName strings.Builder

	entry, _ := SelectMostRecentSegment()

	if entry == nil {
		_, name, _ := CreateSegment()
		entryName.WriteString(name)
	} else {
		entryInfo, err := entry.Info()

		if err != nil {
			log.Fatal(err)
			return ""
		}

		if (entryInfo.Size() / 1024) >= int64(SEGEMENT_SIZE_LIMIT_KB) {
			_, name, _ := CreateSegment()
			entryName.WriteString(name)
		} else {
			entryName.WriteString("database\\")
			entryName.WriteString(entry.Name())
		}
	}
	
	return entryName.String()
}