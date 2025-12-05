package db_operations

import (
	"log"
	"os"
	"strings"

	"github.com/mowgly11/log-db-engine/models"
	"github.com/mowgly11/log-db-engine/utils"
)

func Delete(key string, index map[string]models.IndexEntry) bool {
	var entryName strings.Builder

	entry, _ := utils.SelectMostRecentSegment()

	if entry == nil {
		_, name, _ := utils.CreateSegment()
		entryName.WriteString(name)
	} else {
		entryInfo, err := entry.Info()

		if err != nil {
			log.Fatal(err)
			return false
		}

		if (entryInfo.Size() / 1024) >= int64(SEGEMENT_SIZE_LIMIT_KB) {
			_, name, _ := utils.CreateSegment()
			entryName.WriteString(name)
		} else {
			entryName.WriteString("database/")
			entryName.WriteString(entry.Name())
		}
	}

	file, err := os.OpenFile(entryName.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var data strings.Builder

	data.WriteString("DELETE ")
	data.WriteString(key)
	data.WriteRune(':')
	data.WriteRune('\n')

	if _, err := file.Write([]byte(data.String())); err != nil {
		log.Println("write error:", err)
		return false
	}

	delete(index, key)

	return true
}
