package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// getLinesChannel reads a file in fixed-size chunks (8 bytes) and sends
// each complete line to a channel. This allows concurrent processing
// of lines as they are read without blocking the main function.
func getLinesChannel(file io.ReadCloser) chan string {
	out := make(chan string, 1) // Buffered channel for asynchronous line delivery

	go func() {
		defer file.Close() // Ensure the file is closed when goroutine exits
		defer close(out)   // Close the channel when done reading

		var buffer string // Accumulates data until a full line is read
		for {
			// Read up to 8 bytes at a time
			data := make([]byte, 8)
			n, err := file.Read(data)
			if err != nil {
				if err == io.EOF {
					// Send any remaining data as the last line
					if len(buffer) > 0 {
						out <- buffer
					}
					break
				}
				log.Fatal("Error reading file:", err)
			}

			data = data[:n] // Slice to the number of bytes actually read

			// Process all lines within the current chunk
			for {
				i := bytes.IndexByte(data, '\n') // Find newline in the chunk
				if i == -1 {
					// No newline found; append all data to buffer
					buffer += string(data)
					break
				}

				// Newline found; send complete line
				buffer += string(data[:i])
				out <- buffer
				buffer = ""

				// Move past the newline for further processing
				data = data[i+1:]
				if len(data) == 0 {
					break
				}
			}
		}
	}()

	return out
}

func main() {
	// Open the file "messages.txt" for reading
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	// Get channel of lines from the file
	lines := getLinesChannel(file)

	// Process lines as they are read
	for line := range lines {
		fmt.Printf("Read line: %s\n", line)
	}
}
