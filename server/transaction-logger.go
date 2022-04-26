package server

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const path string = "transactions.txt"
const delimiter string = "|"

var file *os.File

func restoreData() {
	_, err := os.Stat(path)
	log.Println(err)
	if err != nil {
		// if file does not exist, we create it
		file, err = os.Create(path)

		if err != nil {
			log.Println("File creation failed.")
			log.Fatal(err)
		}
		log.Println("File created")
	} else {
		// otherwise, we restore transactions from it
		file, err = os.OpenFile(path, os.O_RDWR, 0644)

		if err != nil {
			log.Println("Couldn't open the file.")
			log.Fatal(err)
		}
		log.Println("File opened")
		// [not implemented]
	}
}

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
