package main

import (
	"fmt"
	"github.com/mowgly11/log-db-engine/db_operations"
)

func main() {
	index := make(map[string]int64)

	db_operations.BuildIndex(index)

	var option int = -1

Loop:
	for option != 3 && option != 0 {
		fmt.Println("=== Log Based Database ===")
		fmt.Println("1. Insert into db")
		fmt.Println("2. Search for a value pair by key")
		fmt.Println("3. Exit")
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

			status := db_operations.Set(key, value, index)

			if status {
				fmt.Println("Successfully added the key-value pair to the db!")
			} else {
				fmt.Println("Failed to add the key-value pair to the db!")
			}
		case 2:
			var key string
			fmt.Printf("Insert the key of the value you're looking for: ")
			fmt.Scan(&key)

			value, _ := db_operations.Get(key, index)
			if value != "" {
				fmt.Printf("value found: %v \n", value)
			} else {
				fmt.Printf("No value was found associated with that key\n")
			}
		case 3:
			fmt.Println("Leaving...")
			break Loop
		default:
			fmt.Println("Invalid option")
		}
	}
}
