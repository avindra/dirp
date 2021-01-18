package dir

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Feeder starts reading from stdin
// and enters selection mode
func Feeder() {
	scanner := bufio.NewScanner(os.Stdin)

	config := make(map[string]string)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			firstChar := line[0:1]
			if firstChar == "#" {
				continue
			}
			split := strings.Split(line, "|")
			name := split[0]
			value := split[1]
			config[name] = value
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
