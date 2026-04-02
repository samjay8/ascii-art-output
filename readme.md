# Ascii Art Output

A Go program that converts text strings into ASCII art using different banner styles. The program supports three banner formats: standard, thinkertoy, and shadow, and can output the result to the console or save it to a file.

## Features

- Convert text to ASCII art using three different banner styles
- Support for newline characters (`\n`) in input strings
- Output to console or save to a file
- Simple command-line interface
- Error handling for invalid inputs and banner names

## Project Structure

```
ascii-art-output/
├── main.go                 # Main entry point and CLI handling
├── go.mod                  # Go module configuration
├── standard.txt            # Standard banner font file
├── thinkertoy.txt          # Thinkertoy banner font file
├── shadow.txt              # Shadow banner font file
└── output/
    ├── ascii_art.go        # ASCII art generation logic
    ├── readbanner.go       # Banner file reading functionality
    └── writefile.go        # File writing functionality
```

## Installation

1. Ensure you have Go installed (version 1.22.2 or higher)
2. Clone or download this repository
3. Navigate to the project directory

```bash
cd ascii-art-output
```

## Usage

The program accepts command-line arguments in the following format:

```bash
go run . [OPTION] [STRING] [BANNER]
```

### Basic Usage

Convert text to ASCII art using the default standard banner:

```bash
go run . "Hello"
```

### Using Different Banners

Specify a banner style (standard, thinkertoy, or shadow):

```bash
go run . "Hello" thinkertoy
go run . "Hello" shadow
```

### Output to File

Save the ASCII art to a file using the `--output` flag:

```bash
go run . --output=output.txt "Hello" standard
```

### Using Newlines

Include newline characters in your text to create multi-line ASCII art:

```bash
go run . "Hello\nWorld"
```

## Banner Styles

The project includes three banner styles:

1. **Standard** - A clean, classic ASCII art style
2. **Thinkertoy** - A playful, toy-like style
3. **Shadow** - A style with shadow effects

Each banner file contains ASCII representations of characters from space (ASCII 32) to tilde (ASCII 126).

## How It Works

### Main Program (`main.go`)

The main program handles command-line arguments and orchestrates the ASCII art generation:

- Parses command-line arguments to determine the input string, banner style, and output file
- Validates the number of arguments and the `--output` flag format
- Calls the appropriate functions to read the banner, generate ASCII art, and output the result

### ASCII Art Generation (`output/ascii_art.go`)

The `AsciiArt` function converts input text to ASCII art:

- Splits the input string by newline characters (`\n`)
- For each character, calculates its position in the banner file
- Reads 8 lines of ASCII art for each character (each character is 8 lines tall)
- Concatenates the ASCII art for all characters in each line
- Returns the complete ASCII art string

### Banner Reading (`output/readbanner.go`)

The `ReadBanner` function reads and validates banner files:

- Validates that the banner name is one of: standard, thinkertoy, or shadow
- Opens the corresponding `.txt` file
- Reads all lines into a string slice
- Returns the banner data for use in ASCII art generation

### File Writing (`output/writefile.go`)

The `WriteFile` function saves ASCII art to a file:

- Parses the `--output=<filename>` flag
- Validates the flag format and filename
- Writes the ASCII art to the specified file with 0644 permissions

## Examples

### Example 1: Basic Text (Standard Banner)

```bash
go run . "Hello"
```

Output:
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
```

### Example 2: Using Shadow Banner

```bash
go run . "Go" shadow
```

Output:
```
                  
  _|_|_|          
_|         _|_|   
_|  _|_| _|    _| 
_|    _| _|    _| 
  _|_|_|   _|_|   
                  
                  
```

### Example 3: Using Thinkertoy Banner

```bash
go run . "Hi" thinkertoy
```

Output:
```
       
o  o   
|  | o 
O--O   
|  | | 
o  o | 
       
       
```

### Example 4: Multi-line Text

```bash
go run . "Hello\nWorld"
```

Output:
```
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
__          __                 _       _  
\ \        / /                | |     | | 
 \ \  /\  / /    ___    _ __  | |   __| | 
  \ \/  \/ /    / _ \  | '__| | |  / _` | 
   \  /\  /    | (_) | | |    | | | (_| | 
    \/  \/      \___/  |_|    |_|  \__,_| 
                                          
                                          
```

### Example 5: Save to File

```bash
go run . --output=result.txt "ASCII Art" standard
```

The ASCII art will be saved to `result.txt`.

## Error Handling

The program handles various error scenarios:

- **Invalid banner name**: Displays "Error: Wrong Banner name" if an unsupported banner is specified
- **Missing output filename**: Displays usage information if the `--output` flag is used without a filename
- **Invalid flag format**: Displays usage information if the `--output` flag is malformed
- **File write errors**: Displays the error message if writing to a file fails

## Usage Message

When incorrect arguments are provided, the program displays:

```
Usage: go run . [OPTION] [STRING] [BANNER]

EX: go run . --output=<fileName.txt> something standard
```

## Technical Details

- **Language**: Go 1.22.2
- **Module Name**: asciiartoutput
- **Character Support**: ASCII characters from space (32) to tilde (126)
- **Character Height**: 8 lines per character
- **File Permissions**: Output files are created with 0644 permissions

## Dependencies

This project uses only Go standard library packages:
- `fmt` - Formatted I/O
- `os` - Operating system functionality
- `strings` - String manipulation
- `bufio` - Buffered I/O

## Authors

- Samuel Ojetunde
- Grazia Orji
- Oluwaseun Showunmi

## License

This project is open source and available for use and modification.
