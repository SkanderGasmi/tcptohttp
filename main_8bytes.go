package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Open the file "messages.txt" for reading
	// If the file doesn't exist or can't be opened, log.Fatal will exit the program
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close() // Ensure the file is closed when the function exits

	for {
		// Create a buffer to read 8 bytes at a time
		data := make([]byte, 8)

		// Read up to 8 bytes from the file
		n, err := file.Read(data)
		if err != nil {
			// Exit the loop on EOF or any read error
			break
		}

		// Print the number of bytes read and the actual data
		fmt.Printf("Read %d bytes: %s\n", n, string(data[:n]))
	}
}
