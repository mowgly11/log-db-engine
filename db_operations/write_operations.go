package db_operations

import (
	//"fmt"
	"log"
	"os"
	"strings"
	"net/url"
	"github.com/mowgly11/log-db-engine/models"
	"github.com/mowgly11/log-db-engine/utils"
)

func Set(key string, value string, index map[string]models.IndexEntry) bool {
	entryName := utils.CreateOrSelectSegment()

	file, err := os.OpenFile(entryName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatal(err)
		return false
	}

	defer file.Close()

	info, err := file.Stat()

	if err != nil {
		log.Fatal(err)
		return false
	}

	var data strings.Builder

	data.WriteString("PUT ")
	data.WriteString(key)
	data.WriteRune(':')
	data.WriteString(url.QueryEscape(value))
	data.WriteRune('\n')

	if _, err := file.Write([]byte(data.String())); err != nil {
		log.Println("write error:", err)
		return false
	}

	indexInfo := models.IndexEntry{SegmentName: strings.Replace(entryName, "database/", "", 1), Offset: int(info.Size())}

	index[key] = indexInfo

	return true
}
