package dirp

import (
	"fmt"
	"strings"
)

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
