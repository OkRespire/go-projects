/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var id int = 0

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		fmt.Println("description is: ", description)
		dir, fileName := getDefaultDirectoryPath()

		fileDirPath := filepath.Join(dir, fileName)

		file, err := os.Open(fileDirPath)
		if err != nil {
			fmt.Println("Error while opening the file")
		}

		csvWriter := OpenCSVWriter(file)

		finalString := "\n" + strconv.Itoa(id) + "," + description + ","

		currentTime := time.Now()

		finalString += currentTime.Format(time.DateTime) + "," + "FALSE"

		fmt.Println(finalString)
		csvWriter.Write([]string{finalString})
		csvWriter.Flush()
		id++
		file.Close()
	},
}

var description string

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().StringVarP("description", "", "Allows you to specify what task you want to add")

	addCmd.Flags().StringVarP(&description, "description", "d", "", "Allows you to specify what task you want to add")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func OpenCSVWriter(file *os.File) *csv.Writer {
	writer := csv.NewWriter(file)
	return writer
}
