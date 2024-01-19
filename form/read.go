package ascii

import (
	"bufio"

	"os"

	"strings"
)

func Line(number int) string {
    // Prepare an empty string to store the line of text
    line := ""
	var file *os.File
    
    // Create a reader tool to read the contents of a file
    reader := bufio.NewReader(file)
    
    // Read lines from the file until we reach the desired line number
    for r := 0; r < number; r++ {
        line, _ = reader.ReadString('\n')
    }
    
    // Remove the trailing newline character, if any
    line = strings.TrimSuffix(line, "\n")
    
    // Return the line of text
    return line
}