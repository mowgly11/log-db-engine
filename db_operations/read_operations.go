package db_operations

import (
	"log"
	"os"
	"strings"

	"github.com/mowgly11/log-db-engine/models"
)

func OpenFile(path string) *os.File {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Get(key string, index map[string]models.IndexEntry) (string, error) {
	value, ok := index[key]

	if !ok {
		return "", nil
	}

	var filePath strings.Builder
	filePath.WriteString("database/")
	filePath.WriteString(value.SegmentName)

	file := OpenFile(filePath.String())
	defer file.Close()

	_, err := file.Seek(int64(value.Offset), 0)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, 1)
	result := make([]byte, 0)

	for {
		_, err := file.Read(buffer)
		if err != nil {
			return string(result), err
		}

		if buffer[0] == '\n' {
			break
		}

		result = append(result, buffer[0])
	}

	return strings.Split(string(result), ":")[1], nil
}
