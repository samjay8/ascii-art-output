package main

import (
	"os"
	"strings"
	"testing"

	asciiartoutput "asciiartoutput/output"
)

func cleanupFile(t *testing.T, filename string) {
	t.Helper()
	os.Remove(filename)
}

// ReadBanner Tests
func TestReadBanner(t *testing.T) {
	tests := []struct {
		banner   string
		hasError bool
	}{
		{"standard", false},
		{"shadow", false},
		{"thinkertoy", false},
		{"invalidfont", true},
	}
	for _, tt := range tests {
		t.Run(tt.banner, func(t *testing.T) {
			font := asciiartoutput.ReadBanner(tt.banner)
			if tt.hasError && len(font) != 0 {
				t.Errorf("expected empty result for %s but got content", tt.banner)
			}
			if !tt.hasError && len(font) == 0 {
				t.Errorf("expected font content for %s but got empty", tt.banner)
			}
		})
	}
}

// AsciiArt Tests
func TestAsciiArt(t *testing.T) {
	font := asciiartoutput.ReadBanner("standard")
	tests := []struct {
		name          string
		input         string
		expectedLines int
	}{
		{"single word", "hello", 8},
		{"multiline", "hello\nworld", 16},
		{"empty input", "", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := asciiartoutput.AsciiArt(tt.input, font)
			lines := strings.Split(strings.TrimRight(result, "\n"), "\n")
			if tt.expectedLines == 0 && result != "" {
				t.Errorf("expected empty output but got: %s", result)
			}
			if tt.expectedLines > 0 && len(lines) != tt.expectedLines {
				t.Errorf("expected %d lines but got %d", tt.expectedLines, len(lines))
			}
		})
	}
}

// WriteFile Tests
func TestWriteFile(t *testing.T) {
	font := asciiartoutput.ReadBanner("standard")
	result := asciiartoutput.AsciiArt("hello", font)

	tests := []struct {
		name     string
		flag     string
		content  string
		hasError bool
	}{
		{"valid flag", "--output=test_output.txt", result, false},
		{"invalid flag", "--wrongflag=test.txt", result, true},
		{"empty content", "--output=empty_output.txt", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := asciiartoutput.WriteFile(tt.flag, tt.content)
			if tt.hasError && err == nil {
				t.Error("expected error but got nil")
			}
			if !tt.hasError && err != nil {
				t.Errorf("expected no error but got: %v", err)
			}
			if !tt.hasError {
				defer cleanupFile(t, strings.TrimPrefix(tt.flag, "--output="))
			}
		})
	}
}

// Integration Tests
func TestIntegration(t *testing.T) {
	tests := []struct {
		name     string
		text     string
		banner   string
		output   string
		hasError bool
	}{
		// Simulates: go run . hello
		{"one arg standard font", "hello", "standard", "", false},
		// Simulates: go run . hello shadow
		{"two args custom font", "hello", "shadow", "", false},
		// Simulates: go run . --output=result.txt hello shadow
		{"three args output to file", "hello", "shadow", "result.txt", false},
		// Simulates: go run . --output=result.txt (no text)
		{"invalid no text", "", "standard", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			font := asciiartoutput.ReadBanner(tt.banner)
			result := asciiartoutput.AsciiArt(tt.text, font)
			if tt.output != "" {
				defer cleanupFile(t, tt.output)
				err := asciiartoutput.WriteFile("--output="+tt.output, result)
				if tt.hasError && err == nil {
					t.Error("expected error but got nil")
				}
				if !tt.hasError && err != nil {
					t.Errorf("expected no error but got: %v", err)
				}
			}
		})
	}
}
