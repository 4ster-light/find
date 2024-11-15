package main

import (
	"fmt"
	"time"

	command "github.com/4ster-light/finder/cmd"
	"github.com/4ster-light/finder/color"

	"github.com/spf13/cobra"
)

func main() {
	var filename, dir string
	var showDirs, showTime bool

	var rootCmd = &cobra.Command{
		Use:   "finder",
		Short: "Finder: A file search tool to locate files within directories",
		Run: func(cmd *cobra.Command, args []string) {
			if filename == "" {
				fmt.Println(color.Colorize("[ERROR] -- Please specify a filename to search with -f or --filename", color.ColorError))
				return
			}

			start := time.Now()
			results := command.SearchFiles(dir, filename, showDirs)
			elapsed := time.Since(start)

			if len(results) > 0 {
				command.PrintResults(results)
				if showTime {
					fmt.Printf("\n%s\nSearch completed in %s with %d result(s).\n", color.Colorize("[INFO]", color.ColorInfo), elapsed, len(results))
				} else {
					fmt.Printf("\n%s\n%d result(s) found.\n", color.Colorize("[INFO]", color.ColorInfo), len(results))
				}
			} else {
				if showTime {
					fmt.Printf("\n%s\nNo results found in %s.\n", color.Colorize("[INFO]", color.ColorInfo), elapsed)
				} else {
					fmt.Println(color.Colorize("\n[INFO]\nNo results found.", color.ColorInfo))
				}
			}
		},
	}

	rootCmd.Flags().StringVarP(&filename, "filename", "f", "", "Name of the file to search for")
	rootCmd.Flags().StringVarP(&dir, "dir", "d", ".", "Directory to search in (default is current directory)")
	rootCmd.Flags().BoolVarP(&showDirs, "show-dirs", "s", false, "Show directories being searched")
	rootCmd.Flags().BoolVarP(&showTime, "time", "t", false, "Show elapsed time for the search")

	rootCmd.Execute()
}
