package command

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/term"

	"github.com/4ster-light/finder/color"
)

// Traverse the file system and return a list of all files that match the filename
func SearchFiles(dir string, filename string, showDirs bool) []string {
	var results []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("%s -- Unable to access %s: %v\n", color.Colorize("[ERROR]", color.ColorError), path, err)
			return nil
		}

		// Skip symlinks to avoid cycles
		if info.Mode()&os.ModeSymlink != 0 {
			return nil
		}

		// Print directories if showDirs is enabled
		if info.IsDir() && showDirs {
			fmt.Printf("%s   %s\n", color.Colorize("[SEARCHING]", color.ColorSearch), color.Colorize(path, color.ColorPath))
		}

		// Append matching file paths to results
		if info.Name() == filename && !info.IsDir() {
			results = append(results, path)
		}

		return nil
	})

	if err != nil {
		fmt.Printf("%s -- Search encountered an issue: %v\n", color.Colorize("[ERROR]", color.ColorError), err)
	}

	return results
}

// Print the results of the search to the terminal.
// Table format and colors are handled internally
func PrintResults(results []string) {
	if len(results) == 0 {
		return
	}

	termWidth := getTerminalWidth()
	maxPathLength := calculateMaxPathLength(termWidth)

	if maxPathLength < 10 {
		printTooNarrowMessage()
		return
	}

	printTableHeader(maxPathLength)

	for i, path := range results {
		displayPath := formatPath(path, maxPathLength)
		printTableRow(i+1, displayPath, maxPathLength)
	}

	printTableFooter(maxPathLength)
}

// * HELPER PRIVATE FUNCTIONS

func getTerminalWidth() int {
	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || termWidth < 40 {
		return 40 // fallback to reasonable default
	}
	return termWidth
}

func calculateMaxPathLength(termWidth int) int {
	const fixedWidth = 10 // 3 (Nº) + 7 (borders and spaces)
	return termWidth - fixedWidth
}

func printTooNarrowMessage() {
	fmt.Printf("\n%s\nTerminal too narrow to display results.\n",
		color.Colorize("[RESULTS]", color.ColorResults))
}

func printTableHeader(maxPathLength int) {
	topBorder := "┌─────┬" + strings.Repeat("─", maxPathLength) + "┐"
	middleBorder := "├─────┼" + strings.Repeat("─", maxPathLength) + "┤"

	fmt.Printf("\n%s\n", color.Colorize("[RESULTS]", color.ColorResults))
	fmt.Println(topBorder)
	fmt.Printf("│ %-3s │ %-*s │\n", "Nº", maxPathLength-2, "Path")
	fmt.Println(middleBorder)
}

func printTableFooter(maxPathLength int) {
	bottomBorder := "└─────┴" + strings.Repeat("─", maxPathLength) + "┘"
	fmt.Println(bottomBorder)
}

func formatPath(path string, maxPathLength int) string {
	strippedPath := color.StripColorCodes(path)
	pathWidth := utf8.RuneCountInString(strippedPath)

	if pathWidth <= maxPathLength-2 {
		return strippedPath
	}

	return truncatePath(strippedPath, maxPathLength)
}

func truncatePath(path string, maxPathLength int) string {
	const ellipsis = "..."
	segments := strings.Split(path, string(os.PathSeparator))

	if len(segments) <= 1 {
		return truncateString(path, maxPathLength-len(ellipsis)) + ellipsis
	}

	lastPart := segments[len(segments)-1]
	remainingSpace := maxPathLength - len(ellipsis) - utf8.RuneCountInString(lastPart)

	if remainingSpace <= 0 {
		return ellipsis + string(os.PathSeparator) + lastPart
	}

	prefix := truncateString(strings.Join(segments[:len(segments)-1], string(os.PathSeparator)), remainingSpace)
	return prefix + ellipsis + string(os.PathSeparator) + lastPart
}

func printTableRow(index int, displayPath string, maxPathLength int) {
	paddedPath := fmt.Sprintf("%-*s", maxPathLength-2, displayPath)
	coloredPath := color.Colorize(paddedPath, color.ColorPath)
	fmt.Printf("│ %-3d │ %s │\n", index, coloredPath)
}

func truncateString(s string, length int) string {
	if length <= 0 {
		return ""
	}
	runes := []rune(s)
	if len(runes) <= length {
		return s
	}
	return string(runes[:length])
}
