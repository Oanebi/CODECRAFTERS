package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var history []string

func addToHistory(cmd string, input string) {
	entry := fmt.Sprintf("[%s] %s", cmd, input)
	history = append(history, entry)

	if len(history) > 2 {
		history = history[1:]
	}
}

func countspace(s string) int {
	return strings.Count(s, " ")
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
func lower(s string) string {
	return strings.ToLower(s)
}
func palindrome(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		fmt.Print(string(s[i]))
	}
	return " "
}

func countchar(s string) int {
	return len([]rune(s))

}
func countwords(s string) int {
	words := strings.Fields(s)
	return len(words)
}
func capitalize(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(w[:1]) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}
func snake(s string) string {
	words := strings.Fields(strings.ToLower(s))
	var result []string

	for _, w := range words {
		var clean strings.Builder

		for _, c := range w {
			if (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') {
				clean.WriteRune(c)
			}
		}

		if clean.Len() > 0 {
			result = append(result, clean.String())
		}
	}

	return strings.Join(result, "_")
}

var smallWords = map[string]bool{
	"a": true, "an": true, "the": true, "and": true,
	"but": true, "or": true, "for": true, "nor": true,
	"on": true, "at": true, "to": true, "by": true,
	"in": true, "of": true, "up": true, "as": true,
	"is": true, "it": true,
}

func Titlecase(s string) string {
	words := strings.Fields(s)

	for i, w := range words {
		l := strings.ToLower(w)

		if i == 0 || !smallWords[l] {
			words[i] = strings.ToUpper(l[:1]) + l[1:]
		} else {
			words[i] = l
		}
	}

	return strings.Join(words, " ")
}

func reverse(s string) string {
	words := strings.Fields(s)

	for i, w := range words {
		runes := []rune(w)

		for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
			runes[l], runes[r] = runes[r], runes[l]
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}
		if input == "fly" {
			fmt.Println(" upper lower, cap, title, snake, reverse, exit:")
		}

		if strings.ToLower(input) == "exit" {
			fmt.Println("Goodbye")
			break
		}
		if strings.ToLower(input) == "history" {
			fmt.Println("--- Last 5 Commands ---")
			if len(history) == 0 {
				fmt.Println("No history yet.")
			}
			for i, entry := range history {
				fmt.Printf("%d: %s\n", i+1, entry)
			}
			continue
		}

		parts := strings.Fields(input)

		if len(parts) < 2 {
			fmt.Println("No text provided")
			continue
		}

		command := strings.ToLower(parts[0])
		text := strings.Join(parts[1:], " ")

		addToHistory(command, text)

		switch command {
		case "upper":
			fmt.Println(Upper(text))
		case "lower":
			fmt.Println(lower(text))
		case "cap":
			fmt.Println(capitalize(text))
		case "title":
			fmt.Println(Titlecase(text))
		case "snake":
			fmt.Println(snake(text))
		case "reverse":
			fmt.Println(reverse(text))
		case "palindrome":
			fmt.Println(palindrome(text))
		case "countchar":
			fmt.Println(countchar(text))
		case "countwords":
			fmt.Println(countwords(text))
		case "countspace":
			fmt.Println(countspace(text))
		default:
			fmt.Println("Unknown command")
		}

	}
}
