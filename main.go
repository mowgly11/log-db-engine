package main

import (
	"bufio"
	"fmt"
	"github.com/mowgly11/log-db-engine/classes"
	"log"
	"os"
	"strings"
)

func main() {
	memtable := make(map[string]string)
	file, err := os.OpenFile("database/segement-1.txt", os.O_APPEND, 0600)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		kvPair := strings.Split(line, ":")
		memtable[kvPair[0]] = kvPair[1]
	}

	var option int

Loop:
	for option != 5 {
		fmt.Println("=== Log Based Database ===")
		fmt.Println("1. Insert into db")
		fmt.Println("2. View all key-value pairs")
		fmt.Println("3. Search for a value pair by key")
		fmt.Println("4. Write data to disk")
		fmt.Println("5. Exit")
		fmt.Printf("select an option: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var key string
			var value string
			fmt.Printf("Write the key you want to add: ")
			fmt.Scan(&key)
			fmt.Printf("Write the value you want to add with the key %v: ", key)
			fmt.Scan(&value)

			status := db_operations.Set(&memtable, key, value)

			if status {
				fmt.Println("Successfully added the key-value pair to the db!")
			} else {
				fmt.Println("Failed to add the key-value pair to the db!")
			}
		case 2:
			fmt.Println("== MAP START ==")
			for k, v := range memtable {
				fmt.Printf("%s:%s\n", k, v)
			}
			fmt.Println("== MAP END ==")
		case 3:
			var key string
			fmt.Printf("Insert the key of the value you're looking for: ")
			fmt.Scan(&key)

			var value string = db_operations.Get(&memtable, key)
			if value != "" {
				fmt.Printf("Value found: %v \n", value)
			} else {
				fmt.Printf("No value was found associated with that key\n")
			}
		case 4:
			db_operations.WriteToDisk(file, &memtable)
		case 5:
			fmt.Println("Leaving...")
			break Loop
		default:
			fmt.Println("Invalid option")
		}
	}

	defer file.Close()
}
