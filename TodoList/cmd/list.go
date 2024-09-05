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
	"time"

	"github.com/mergestat/timediff"
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
		fmt.Println("list called")

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

		count := 0
		var tStr string
		// TODO: MAKE SURE THE TIME IS READABLE!
		for _, record := range records {
			if count != 0 {
				t, err := time.Parse("2006-01-02 15:04:05", record[2])
				if err != nil {
					log.Fatal(err)
				}
				tStr = timediff.TimeDiff(time.Now(), timediff.WithStartTime(t))

				fmt.Println(record[0], record[1], tStr, record[3])
			} else {
				fmt.Println(record[0], record[1], record[2], record[3])
				count++

			}
		}
	},
}

var all bool

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVarP(&all, "all", "a", true, "Displays all the tasks completed or not")
}

// func printRecords(records [][]string) {
//   if all {
//     fmt.Println("All tasks")
//   } else {
//     fmt.Println("Incomplete tasks")
//   }
//   for _, record := range records {
//     if !all {
//
//     }
//   }
// }
