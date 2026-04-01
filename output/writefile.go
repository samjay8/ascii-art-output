package asciiartoutput

import (
	"fmt"
	"os"
	"strings"
)

func WriteFile(flag string, asciiart string) error {

	output := strings.Split(flag, "=")

	if output[0] != "--output" {
		return fmt.Errorf(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
	}

	if output[1] == "" {
		return fmt.Errorf(`Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`)
	}

	err := os.WriteFile(output[1], []byte(asciiart), 0644)

	return err
}
