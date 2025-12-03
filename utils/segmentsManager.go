package utils

import (
	//"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

const SEGEMENT_SIZE_LIMIT_KB int32 = 200

func SelectMostRecentSegment() (int) {
	entries, err := os.ReadDir("database")
	files := []int{}

	if err != nil {
		log.Fatal(err)
	}
	
	if len(entries) == 0 {
		return 0
	}

	for _, entry := range entries {
		info, err := entry.Info()

		if err != nil {
			log.Fatal(err)
			continue
		}

		segNum, _ := strconv.Atoi(strings.Replace(strings.Split(info.Name(), "-")[1], ".txt", "", 1))
		files = append(files, segNum)
	}

	segmentNumber := slices.Max(files)

	if err != nil {
		log.Fatal(err)
	}

	return segmentNumber
}

func CreateSegment() (*os.File, error) {
	segmentNumber := SelectMostRecentSegment()

	var segmentName strings.Builder
	segmentName.WriteString("database/segment-")
	segmentName.WriteString(strconv.Itoa(segmentNumber + 1))
	segmentName.WriteString(".txt")

	file, err := os.Create(segmentName.String())

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	return file, err
}
