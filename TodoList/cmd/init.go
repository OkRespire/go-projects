/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init called")
	},
}

const DETECTED_OS = runtime.GOOS

func init() {
	rootCmd.AddCommand(initCmd)

	dir := getDefaultDirectoryPath()
	createDirectoryIfNotExists(dir)
	InitialiseToDo(dir)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getDefaultDirectoryPath() string {
	var dir string

	switch DETECTED_OS {
	case "windows":
		dir = os.ExpandEnv("%LOCALAPPDATA%\\Respire-ToDo\\")
	case "linux":
		dir = os.ExpandEnv("$HOME/.config/Respire-ToDo/")
	case "darwin":
		dir = os.ExpandEnv("$HOME/.config/Respire-ToDo/")
	}

	return dir
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

func InitialiseToDo(dir string) *os.File {
	var csvFile *os.File
	if _, err := os.Stat("list.csv"); err != nil {
		if os.IsNotExist(err) {
			csvFile, err := os.Create(filepath.Join(dir, "list.csv"))
			if err != nil {
				log.Fatal("Failed to create file", err)
				return nil
			}
			csvFile.Close()
			fmt.Println("File created")
		}
	}

	return csvFile
}
