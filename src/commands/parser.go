package commands

import (
	"errors"
)

var unbalancedQuotes = "Failed to parse command: unbalanced quotes"

func parseCommandParts(command string) ([]string, error) {
	parts := []string{}

	current := ""
	quote := ""
	escape := false

	for _, c := range command {
		char := string(c)

		// Handle escapes first
		if escape {
			escape = false

			if char == "n" {
				current += "\n"
			} else {
				current += char
			}

			continue
		}

		if char == "\\" {
			escape = true
			continue
		}

		// Handle quotes
		if quote == "" && (char == "\"" || char == "'") {
			quote = char
			continue
		}

		if quote != "" {
			if quote == char {
				parts = append(parts, current)
				current = ""
				quote = ""
				continue
			}

			current += char
			continue
		}

		if char == " " {
			if current == "" {
				continue
			}

			parts = append(parts, current)
			current = ""
			continue
		}

		current += char
	}

	if quote != "" {
		return nil, errors.New("Failed to parse command: unbalanced quotes")
	}

	if escape {
		return nil, errors.New("Failed to parse command: invalid escape at end of command")
	}

	if current != "" {
		parts = append(parts, current)
	}

	return parts, nil
}
