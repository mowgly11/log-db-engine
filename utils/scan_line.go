package utils

import (
	"bufio"
	"io"
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