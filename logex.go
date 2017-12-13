package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("file.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", ":", err)
	}

	logger := log.New(file, "PREFIX: ", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println("This goes to file")

	log.SetPrefix("prefix- ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("This goes to stdout")

	logger.Println("PREFIX: and flags are not affected")
}
