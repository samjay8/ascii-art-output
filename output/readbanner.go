package asciiartoutput

import (
	"bufio"
	"fmt"
	"os"
)

func ReadBanner(banner string) []string {
	
	if banner != "standard" && banner != "thinkertoy" && banner != "shadow" {
		fmt.Println("Error: Wrong Banner name")
		return nil
	}

	filename := banner + ".txt"

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Opening file", err)
		return nil
	}

	defer file.Close()

	readString := bufio.NewScanner(file)

	var keepString []string

	for readString.Scan() {
		keepString = append(keepString, readString.Text())
	}

	return keepString
}
