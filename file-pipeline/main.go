



// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Blessing Anebi]
// Squad:  [Goroutines]

 // ═══════════════════════════════════════════

// SQUAD PIPELINE CONTRACT
// Squad: Goroutines
// ───────────────────────────────────────────
//
// Input line types:
//
//   1. Normal report lines
//   2. Lines in ALL CAPS
//   3. Lines in all lowercase
//   4. Lines starting with TODO:
//   5. Lines starting with CLASSIFIED:
//   6. Lines that are only dashes or blank
//   7. Lines with leading/trailing whitespace
//   8. Lines containing numbers and symbols
//
// Transformation rules (in order):
//
//   1. Trim all leading and trailing whitespace
//   2. Remove lines that are only dashes or blank
//   3. Replace TODO: with ✦ ACTION:
//   4. Replace CLASSIFIED: with [REDACTED]:
//   5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
//
//   Header: yes — SENTINEL FIELD REPORT — PROCESSED
//   Line numbering format: 001., 002., 003. (three-digit zero-padded)
//   Summary block: yes — Lines Processed, Lines Written, Lines Removed
//
// Terminal summary fields:
//
//   ✦ Lines read    : <number>
//   ✦ Lines written : <number>
//   ✦ Lines removed : <number>
//   ✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func trim(s string) string {
	return strings.TrimSpace(s)
}

func Dashes(s string) bool {
	if s == "" {
		return true
	}
	for _, ch := range s {
		if ch != '-' {
			return false
		}
	}
	return true
}


func handleTodo(s string) string {
	return strings.ReplaceAll(s, "TODO:", "✦ ACTION:")
}

func handleClassified(s string) string {
	if strings.HasPrefix(s, "CLASSIFIED:") {
		return "[REDACTED]:" + s[len("CLASSIFIED:"):]
	}
	return s
}


func handleReverse(s string) string {
	if strings.Contains(s, "REVERSE") {
		words := strings.Fields(s)
		for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
			words[i], words[j] = words[j], words[i]
		}
		return strings.Join(words, " ")
	}
	return s
}

func main() {
	
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputPath, outputPath := os.Args[1], os.Args[2]

	
	inputFile, err := os.Open(inputPath)
	if err != nil {
		fmt.Printf("✗ Error: %v\n", err)
		return
	}
	defer inputFile.Close()

	
	outputFile, err := os.Create(outputPath)
	if err != nil {
		fmt.Printf("✗ Error: %v\n", err)
		return
	}
	defer outputFile.Close()


	linesRead, linesWritten, linesRemoved := 0, 0, 0
	
	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)

	
	fmt.Fprintln(writer, "SENTINEL FIELD REPORT — PROCESSED")

	for scanner.Scan() {
		linesRead++
		line := scanner.Text()

		
		line = trim(line)

		
		if Dashes(line) {
			linesRemoved++
			continue
		}

		
		line = handleTodo(line)

		
		line = handleClassified(line)

		
		line = handleReverse(line)

		
		linesWritten++
		fmt.Fprintf(writer, "%03d. %s\n", linesWritten, line)
	}

	writer.Flush()

	
	fmt.Println("───────────────────────────────────────────")
	fmt.Printf("✦ Lines read    : %d\n", linesRead)
	fmt.Printf("✦ Lines written : %d\n", linesWritten)
	fmt.Printf("✦ Lines removed : %d\n", linesRemoved)
	fmt.Println("✦ Rules applied : Trim whitespace, Remove blank/dash lines, Replace TODO, Replace CLASSIFIED, Reverse REVERSE lines")
	fmt.Println("───────────────────────────────────────────")
}