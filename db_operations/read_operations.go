package db_operations

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func ReadLineAndLen(r *bufio.Reader) (string, int, error) {
	line, err := r.ReadString('\n')

	if err != nil && err != io.EOF {
		return "", 0, err
	}

	length := len(line)

	trimmed := strings.TrimRight(line, "\r\n")

	if err == io.EOF {
		if length > 0 {
			return trimmed, length, io.EOF
		}
		return "", 0, io.EOF
	}

	return trimmed, length, nil
}

func OpenFile(path string) *os.File {
	file, err := os.Open(path);

	if err != nil {
		log.Fatal(err)
	}

	return file
}

func Get(key string, index *map[string]int) {

}