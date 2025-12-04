package db_operations

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func BuildHashIndex(index map[string]int) bool {
	entries, err := os.ReadDir("database")

	if err != nil {
		log.Fatal(err)
		return false
	}

	for _, entry := range entries {
		var entryName strings.Builder
		entryName.WriteString("database/")
		entryName.WriteString(entry.Name())

		file, err := os.Open(entryName.String())

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
					if strings.HasPrefix(key, "DELETE ") {
						delete(index, strings.Replace(key, "DELETE ", "", 1))
					} else {
						key = strings.Replace(key, "PUT ", "", 1)
						index[key] = int(nextOffset)
					}
				}
			}

			nextOffset += int64(length)

			if err == io.EOF {
				break
			}
		}
	}

	return true
}
