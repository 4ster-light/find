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

func PrintResults(results []string) {
	if len(results) == 0 {
		return
	}

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = 80 // fallback to reasonable default
	}
	if termWidth < 40 {
		termWidth = 40
	}

	// Calculate available space for path column
	// Fixed width: 3 (Nº) + 7 (borders and spaces) = 10
	const fixedWidth = 10
	maxPathLength := termWidth - fixedWidth

	topBorder := "┌─────┬" + strings.Repeat("─", maxPathLength) + "┐"
	middleBorder := "├─────┼" + strings.Repeat("─", maxPathLength) + "┤"
	bottomBorder := "└─────┴" + strings.Repeat("─", maxPathLength) + "┘"

	if maxPathLength < 10 {
		fmt.Printf("\n%s\nTerminal too narrow to display results.\n",
			color.Colorize("[RESULTS]", color.ColorResults))
		return
	}

	fmt.Printf("\n%s\n", color.Colorize("[RESULTS]", color.ColorResults))
	fmt.Println(topBorder)
	fmt.Printf("│ %-3s │ %-*s │\n", "Nº", maxPathLength-2, "Path")
	fmt.Println(middleBorder)

	for i, path := range results {
		// Strip existing color codes for length calculation
		strippedPath := color.StripColorCodes(path)
		displayPath := strippedPath

		// Truncate path if necessary
		pathWidth := utf8.RuneCountInString(strippedPath)
		if pathWidth > maxPathLength-2 {
			// Try to preserve the last part of the path
			segments := strings.Split(strippedPath, string(os.PathSeparator))
			if len(segments) > 1 {
				lastPart := segments[len(segments)-1]
				if utf8.RuneCountInString(lastPart) > maxPathLength-5 {
					// If even the last part is too long, do simple truncation
					displayPath = truncateString(strippedPath, maxPathLength-5) + "..."
				} else {
					// Show ".../" + last part
					remainingSpace := maxPathLength - 5 - utf8.RuneCountInString(lastPart)
					if remainingSpace > 0 {
						prefix := truncateString(strings.Join(segments[:len(segments)-1], string(os.PathSeparator)), remainingSpace)
						displayPath = prefix + ".../" + lastPart
					} else {
						displayPath = ".../" + lastPart
					}
				}
			} else {
				displayPath = truncateString(strippedPath, maxPathLength-5) + "..."
			}
		}

		// Pad the path (subtract 2 from maxPathLength to fix extra space) and apply color
		paddedPath := fmt.Sprintf("%-*s", maxPathLength-2, displayPath)
		coloredPath := color.Colorize(paddedPath, color.ColorPath)
		fmt.Printf("│ %-3d │ %s │\n", i+1, coloredPath)
	}

	fmt.Println(bottomBorder)
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
