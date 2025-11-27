package db_operations

import (
	"os"
	"strings"
)

func Set(memtable *map[string]string, key string, value string) bool {
	(*memtable)[key] = value

	return true
}

func WriteToDisk(db_instance *os.File, memtable *map[string]string) {
	var data strings.Builder

	for key, value := range *memtable {
		data.WriteString(key)
		data.WriteRune(':')
		data.WriteString(value)
		data.WriteRune('\n')
	}

	db_instance.Write([]byte(data.String()))
}