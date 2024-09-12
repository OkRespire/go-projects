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
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Use:   "add",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to qickly create a Cobra application.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")

		fmt.Println("description is: ", args[0])
		dir, fileName := getDefaultDirectoryPath()

		idDir := idDir()

		fileDirPath := filepath.Join(dir, fileName)

		file, err := os.OpenFile(fileDirPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
		if err != nil {
			fmt.Println("Error while opening the file")
		}

		// TODO: MAKE THE IDFILE AND FILEID INTO ONE SINGLE VARIABLE
		fileId, err := os.OpenFile(filepath.Join(dir, idDir), os.O_RDONLY, 0444)
		if err != nil {
			fmt.Println(err)
		}

		idFile, err := os.OpenFile(filepath.Join(dir, idDir), os.O_RDWR, 0666)
		if err != nil {
			fmt.Println(err)
		}

		s := bufio.NewScanner(idFile)

		var id string
		for s.Scan() {
			id = s.Text()
		}
		fmt.Println(id)
		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal(err)
		}

		idInt++
		fmt.Println(idInt)
		id = strconv.Itoa(idInt)

		f, err := idFile.Write([]byte(id))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("\n%d bytes written\n", f)

		defer fileId.Close()
		defer idFile.Close()
		csvWriter := OpenCSVWriter(file)

		finalString := id + "," + args[0] + ","

		currentTime := time.Now()

		finalString += currentTime.Format(time.DateTime) + "," + "false"

		finalStringArr := strings.Split(finalString, ",")

		for _, a := range finalStringArr {
			fmt.Println(a)
		}
		err = csvWriter.Write(finalStringArr)
		if err != nil {
			log.Fatal("failed to write in csv file", err)
		}
		fmt.Println("Written to file at", fileDirPath)
		csvWriter.Flush()
		if err := csvWriter.Error(); err != nil {
			log.Fatal("failed to flush file", err)
		}

		defer file.Close()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().StringVarP("description", "", "Allows you to specify what task you want to add")

	// addCmd.Flags().StringVarP(&description, "description", "d", "", "Allows you to specify what task you want to add")
	// addCmd.MarkFlagRequired("description")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func OpenCSVWriter(file *os.File) *csv.Writer {
	writer := csv.NewWriter(bufio.NewWriter(file)) // Create a new CSV writer using the provided file)
	return writer
}

func idDir() string {
	var id string
	switch runtime.GOOS {
	case "windows":
		id = "\\id.txt"
	case "linux":
		id = "/id.txt"
	case "darwin":
		id = "/id.txt"
	}
	return id
}
