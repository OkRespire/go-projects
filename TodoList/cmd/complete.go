/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Use:   "complete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		dir, fileName := getDefaultDirectoryPath()

		filePath := filepath.Join(dir, fileName)

		fileRead, err := os.OpenFile(filePath, os.O_RDONLY, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}

		csvReader := csv.NewReader(fileRead)

		fileContents, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fileContents, id)
		flag := false

		defer fileRead.Close()

		fileWrite, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, os.ModePerm)

		csvWriter := csv.NewWriter(fileWrite)

		for _, record := range fileContents {
			if record[0] == id {
				if record[3] == "true" {
					fmt.Println("Task already completed")
					flag = true
					break
				}
				record[3] = "true"
				csvWriter.WriteAll(fileContents)
				flag = true
			}
		}
		if !flag {
			fmt.Println("ID not found")
		}

		defer fileWrite.Close()
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
