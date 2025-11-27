package db_operations

import (
	"log"
	"os"
	"strings"
)

func Set(key string, value string) bool {
	file, err := os.OpenFile("database/database.txt", os.O_APPEND, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var data strings.Builder

	data.WriteString(key)
	data.WriteRune(':')
	data.WriteString(value)
	data.WriteRune('\n')

	file.Write([]byte(data.String()))

	return true
}
