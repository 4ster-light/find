package color

import "strings"

// Color codes for terminal output
const (
	ColorInfo    string = "\033[1;34m" // Bold blue for [INFO]
	ColorResults string = "\033[1;32m" // Bold green for [RESULTS]
	ColorPath    string = "\033[1;36m" // Cyan for file paths
	ColorError   string = "\033[1;31m" // Red for [ERROR]
	ColorSearch  string = "\033[1;33m" // Yellow for [SEARCHING]
	colorReset   string = "\033[0m"    // Reset color
)

// Colorize text with a given color code
func Colorize(text, color string) string {
	return color + text + colorReset
}

// Strip color codes from a string for analysis
func StripColorCodes(text string) string {
	stripped := text
	for _, code := range []string{
		ColorInfo,
		ColorResults,
		ColorPath,
		ColorError,
		ColorSearch,
		colorReset,
	} {
		stripped = strings.ReplaceAll(stripped, code, "")
	}
	return stripped
}
