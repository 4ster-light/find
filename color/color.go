package color

import "strings"

const (
	ColorInfo    = "\033[1;34m" // Bold blue for [INFO]
	ColorResults = "\033[1;32m" // Bold green for [RESULTS]
	ColorPath    = "\033[1;36m" // Cyan for file paths
	ColorError   = "\033[1;31m" // Red for [ERROR]
	ColorSearch  = "\033[1;33m" // Yellow for [SEARCHING]
	colorReset   = "\033[0m"    // Reset
)

func Colorize(text, color string) string {
	return color + text + colorReset
}

func StripColorCodes(text string) string {
	stripped := text
	for _, code := range []string{ColorInfo, ColorResults, ColorPath, ColorError, ColorSearch, colorReset} {
		stripped = strings.ReplaceAll(stripped, code, "")
	}
	return stripped
}
