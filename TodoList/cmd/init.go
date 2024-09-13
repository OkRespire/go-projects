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

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialises a new ToDo List",
	Long:  "USAGE: <COMMAND> init",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
		dir, fileNamePath := getDefaultDirectoryPath()
		id := idDir()
		createDirectoryIfNotExists(dir)
		InitialiseToDo(dir, fileNamePath)
		InitialiseID(id, dir)
		csvFile, err := os.OpenFile(filepath.Join(dir, fileNamePath), os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal("cannot open file.")
		}
		WriteToCSV(csvFile)
	},
}

const DETECTED_OS = runtime.GOOS

func init() {
	rootCmd.AddCommand(initCmd)
}

func getDefaultDirectoryPath() (string, string) {
	var dir string
	var fileNamePath string

	switch DETECTED_OS {
	case "windows":
		dir = os.ExpandEnv("%APPDATA%\\Respire-ToDo\\")
		fileNamePath = "\\list.csv"
	case "linux":
		dir = os.ExpandEnv("$HOME/.config/Respire-ToDo/")
		fileNamePath = "/list.csv"
	case "darwin":
		dir = os.ExpandEnv("$HOME/.config/Respire-ToDo/")
		fileNamePath = "/list.csv"
	}

	return dir, fileNamePath
}

func createDirectoryIfNotExists(dir string) {
	if _, err := os.Stat(dir); err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(dir, 0750)
			if err != nil {
				fmt.Println("Failed to create directory:", err)
				return
			}
			fmt.Println("Directory created at", dir)
		}
	}
}

func InitialiseToDo(dir string, fileNamePath string) {
	var csvFile *os.File
	if _, err := os.Stat(filepath.Join(dir, fileNamePath)); err != nil {
		if os.IsNotExist(err) {
			csvFile, err := os.Create(filepath.Join(dir, fileNamePath))
			if err != nil {
				log.Fatal("Failed to create file", err)
			}
			csvFile.Close()
			fmt.Println("File created")
		}
		defer csvFile.Close()
	}
}

func WriteToCSV(csvFile *os.File) error {
	csvWriter := csv.NewWriter(bufio.NewWriter(csvFile))
	record := []string{"ID", "Task", "Created", "Done"}

	err := csvWriter.Write(record)
	if err != nil {
		fmt.Println("failed to write in csv file", err)
		return err
	}

	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		log.Fatal("failed to flush file", err)
	}
	fmt.Println("wrote in csv file")
	return nil
}

func InitialiseID(id string, dir string) {
	if _, err := os.Stat(filepath.Join(dir, id)); err != nil {
		if os.IsNotExist(err) {
			idFile, err := os.Create(filepath.Join(dir, id))
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("File created")
			f, err := idFile.WriteString("0")
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Successfully wrote in id file.\nThis took", f, "bytes")
		}
	}
}
