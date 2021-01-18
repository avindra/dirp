package dir

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// ConfigSelection represents a list
// of dirs to select from
type ConfigSelection map[string]string

// Feeder starts reading from stdin
// and enters selection mode
func Feeder() {
	scanner := bufio.NewScanner(os.Stdin)

	config := make(ConfigSelection)

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

// Selector issues a picker provided a config
func Selector(config ConfigSelection) {
	// config is now ready
	names := make([]string, 0, len(config))
	for k := range config {
		names = append(names, k)
	}

	result, err := Fuzz(strings.NewReader(strings.Join(names, "\n")))
	if err == nil && len(result) > 0 {
		choice := config[result]
		fmt.Print(choice)
	}
}
