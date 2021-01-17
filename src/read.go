package dir

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func scanner() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if hasInput() {
		fmt.Print("hup")
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
