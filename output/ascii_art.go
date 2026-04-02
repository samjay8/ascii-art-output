package asciiartoutput

import (
	"fmt"
	"strings"
)

func AsciiArt(input string, bannerlines []string) string {
	
	if input == "" {
		return ""
	}

	if input == `\n` {
		return "\n"
	}

	textsplit := strings.Split(input, `\n`)
	var rowString strings.Builder

	for _, char := range textsplit {
		if char == "" {
			fmt.Println()
		}

		length := 8
		for row := range length {
			for col := 0; col < len(char); col++ {
				post := int(char[col]-32)*9 + 1
				rowString.WriteString(bannerlines[post+row])
			}
			rowString.WriteRune('\n')
		}
	}
	return rowString.String()
}
