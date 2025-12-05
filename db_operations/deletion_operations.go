package db_operations

import (
	"log"
	"os"
	"strings"

	"github.com/mowgly11/log-db-engine/models"
	"github.com/mowgly11/log-db-engine/utils"
)

func Delete(key string, index map[string]models.IndexEntry) bool {
	entryName := utils.CreateOrSelectSegment()

	file, err := os.OpenFile(entryName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

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
