package main

import (
	asciiartoutput "asciiartoutput/output"
	"fmt"
	"os"
	"strings"
)

func main() {

	usage := `Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard`

	switch len(os.Args) {

	case 2:
		if strings.HasPrefix(os.Args[1], "--output") {
			fmt.Println(usage)
			return
		}
		font := asciiartoutput.ReadBanner("standard")
		result := asciiartoutput.AsciiArt(os.Args[1], font)
		fmt.Print(result)

	case 3:
		if strings.HasPrefix(os.Args[1], "--output") {
			fmt.Println(usage)
			return
		}
		font := asciiartoutput.ReadBanner(os.Args[2])
		result := asciiartoutput.AsciiArt(os.Args[1], font)
		fmt.Print(result)

	case 4:
		font := asciiartoutput.ReadBanner(os.Args[3])
		result := asciiartoutput.AsciiArt(os.Args[2], font)
		err := asciiartoutput.WriteFile(os.Args[1], result)
		if err != nil {
			fmt.Println(err)
			return
		}

	default:
		fmt.Println(usage)
	}

}