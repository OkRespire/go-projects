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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")

		id := args[0]

		dir, fileName := getDefaultDirectoryPath()

		filePath := filepath.Join(dir, fileName)

		file, err := os.OpenFile(filePath, os.O_RDWR, 0666)
		if err != nil {
			log.Fatal(err)
		}

		csvReader := csv.NewReader(file)

		fileContents, err := csvReader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(fileContents, id)

		// TODO: Make sure the record with the correct id is deleted and updated on the file.
		for _, record := range fileContents {
			if record[0] == id {
			}
		}

		defer file.Close()
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
