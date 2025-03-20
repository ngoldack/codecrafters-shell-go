package parser

import "strings"

func Parse(command string) ([]string, error) {
	args := make([]string, 0)
	cmd, command, _ := strings.Cut(command, " ")
	args = append(args, cmd)

	var currentArg strings.Builder
	inQuotes := false
	quoteChar := rune(0)
	escaped := false

	for i := 0; i < len(command); i++ {
		c := rune(command[i])

		if escaped {
			currentArg.WriteRune(c)
			escaped = false
			continue
		}

		switch c {
		case escapeChar:
			if inQuotes && quoteChar == singleQuote {
				currentArg.WriteRune(c)
			} else {
				escaped = true
			}
		case singleQuote, doubleQuote:
			if !inQuotes {
				inQuotes = true
				quoteChar = c
			} else if c == quoteChar {
				inQuotes = false
				quoteChar = 0
			} else {
				currentArg.WriteRune(c)
			}
		case spaceChar:
			if inQuotes {
				currentArg.WriteRune(c)
			} else if currentArg.Len() > 0 {
				args = append(args, currentArg.String())
				currentArg.Reset()
			}
		default:
			currentArg.WriteRune(c)
		}
	}

	if currentArg.Len() > 0 {
		args = append(args, currentArg.String())
	}

	return args, nil
}

const singleQuote = '\''
const doubleQuote = '"'
const escapeChar = '\\'
const spaceChar = ' '

var escapedCharsInDoubleQuotes = []rune{
	'n',
	'\\',
	'$',
	'"',
}
