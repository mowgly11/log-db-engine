package writer

import (
	"os"
	"strings"
)

type Writer struct {
	DB_PATH string
}

func AppendRecord(memtable *map[string]string, key string, value string) (bool, error) {
	var err error

	(*memtable)[key] = value

	return true, err
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