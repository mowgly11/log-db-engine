package db_operations

import (
	"log"
	"os"
	"strings"
)

func Set(key string, value string, index map[string]int64) bool {
	file, err := os.OpenFile("database/database.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		log.Fatal(err)
	}

	var data strings.Builder
	
	data.WriteString(key)
	data.WriteRune(':')
	data.WriteString(value)
	data.WriteRune('\n')
	
	if _, err := file.Write([]byte(data.String())); err != nil {
		log.Println("write error:", err)
		return false
	}

	index[key] = info.Size()

	return true
}
