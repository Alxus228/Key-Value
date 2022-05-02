package server

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// String path contains the name of the file
const path string = "transactions.txt"

// We use delimiter to join or split transactions while writing or reading them from the file, accordingly.
const delimiter string = "|"

var file *os.File

// restoreData creates the file if it doesn't exist and uses
// client package to reestablish data of the key-value vault.
func restoreData() {
	//Firstly, we check the state of our file.
	_, err := os.Stat(path)
	if err != nil {
		// if file does not exist, we create it
		file, err = os.Create(path)

		if err != nil {
			log.Println("File creation failed.")
			log.Fatal(err)
			return
		}

		log.Println("File created")
	} else {
		// otherwise, we restore transactions from it
		file, err = os.OpenFile(path, os.O_RDWR, 0644)

		if err != nil {
			log.Println("Couldn't open the file.")
			log.Fatal(err)
			return
		}

		log.Println("File opened")
		// [not implemented]
	}
}

// logTransaction writes into the file all vital information about successful PUT and DELETE requests.
func logTransaction(method string, key interface{}, value interface{}) {
	transaction := strings.Join([]string{
		method,
		fmt.Sprintf("%v", key),
		fmt.Sprintf("%v", value),
		time.Now().Format(time.UnixDate) + "\n",
	}, delimiter)

	_, err := file.WriteString(transaction)
	if err != nil {
		log.Println("Writing transaction log into the 'transactions.exe' went wrong.")
		log.Println(err)
	}
}
