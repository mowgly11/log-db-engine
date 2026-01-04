package utils

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mowgly11/log-db-engine/models"
	//"github.com/mowgly11/log-db-engine/models"
)

// WARNING: this is a resource-heavy function, only run this as a background process or otherwise it will cause the database to throttle
func CompactAndMerge() {
	compactionStorage := make(map[string]string)
	allSegmentsRaw, err1 := os.ReadDir("database")

	if err1 != nil {
		log.Fatal(err1)
	}

	// read each kv pair alone

	for i := range allSegmentsRaw {
		var segmentName strings.Builder

		segmentName.WriteString("database/segment-")
		segmentName.WriteString(strconv.Itoa(i + 1))
		segmentName.WriteString(".txt")

		file, err := os.Open(segmentName.String())
		if err != nil {
			log.Fatal(err)
		}

		reader := bufio.NewReader(file)

		for {
			line, length, err := ReadLineAndLen(reader)

			if err == io.EOF && length == 0 {
				break
			}

			if line != "" {
				colon := strings.IndexByte(line, ':')
				if colon != -1 {
					key := line[:colon]
					value := line[colon+1:]
					if strings.HasPrefix(key, "DELETE ") {
						delete(compactionStorage, strings.Replace(key, "DELETE ", "", 1))
					} else {
						key = strings.Replace(key, "PUT ", "", 1)
						compactionStorage[key] = value
					}
				}
			}
		}

		if err := file.Close(); err != nil {
			log.Println("error closing file:", err)
		}
	}

	var segmentNameHolder strings.Builder
	var segmentNumber int = 1

	segmentNameHolder.WriteString("database/segment-1.txt")

	var file *os.File
	var err2 error

	file, err2 = os.OpenFile(segmentNameHolder.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err2 != nil {
		log.Fatal(err2)
	}

	os.Truncate(segmentNameHolder.String(), 0)

	for key, value := range compactionStorage {
		info, err := file.Stat()

		var data strings.Builder

		if err != nil {
			log.Fatal(err)
		}

		data.WriteString("PUT ")
		data.WriteString(key)
		data.WriteRune(':')
		data.WriteString(value)
		data.WriteRune('\n')

		if _, err := file.Write([]byte(data.String())); err != nil {
			log.Println("write error:", err)
		}

		if (info.Size() / 1024) > int64(models.SEGEMENT_SIZE_LIMIT_KB) {
			file.Close()

			segmentNumber++

			segmentNameHolder.Reset()
			segmentNameHolder.WriteString("database/segment-")
			segmentNameHolder.WriteString(strconv.Itoa(segmentNumber))
			segmentNameHolder.WriteString(".txt")

			file, err2 = os.OpenFile(segmentNameHolder.String(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			os.Truncate(segmentNameHolder.String(), 0)

			if err2 != nil {
				log.Fatal(err2)
			}
		}
	}

	file.Close()

	for i, segment := range allSegmentsRaw {
		var segmentName strings.Builder

		segmentName.WriteString("database/")
		segmentName.WriteString(segment.Name())

		if i+1 > segmentNumber {
			err := os.Remove(segmentName.String())
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
