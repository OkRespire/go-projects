/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		dir, fileName := getDefaultDirectoryPath()

		filePath := filepath.Join(dir, fileName)
		csvFile, err := os.OpenFile(filePath, os.O_RDONLY, 0444)
		if err != nil {
			log.Fatal(err)
		}

		csvReader := csv.NewReader(bufio.NewReader(csvFile))

		records, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		w := tabwriter.NewWriter(os.Stdout, 25, 0, 2, ' ', 0)
		for _, record := range records {
			var s string
			if record[0] == "" {
				continue
			}
			if all {
				s = record[0] + "\t" + record[1] + "\t" + record[2] + "\t" + record[3]
			} else {
				s = record[0] + "\t" + record[1] + "\t" + record[2]
			}

			fmt.Fprintln(w, s)
			w.Flush()
		}
	},
}

var all bool

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&all, "all", "a", false, "Displays all the tasks completed or not")
}
