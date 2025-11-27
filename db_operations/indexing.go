package db_operations

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func BuildIndex(index *map[string]int) bool {
	file, err := os.Open("database/database.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var nextOffset int64 = 0

	for {
		line, length, err := ReadLineAndLen(reader)

		if err == io.EOF && length == 0 {
			break
		}

		if line != "" {
			colon := strings.IndexByte(line, ':')
			if colon != -1 {
				key := line[:colon]
				(*index)[key] = int(nextOffset)
			}
		}

		nextOffset += int64(length)

		if err == io.EOF {
			break
		}
	}

	return true
}
