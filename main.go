package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "My sample Cobra app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from myapp!")
	},
}

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Change",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		oldString, _ := cmd.Flags().GetString("old")
		newString, _ := cmd.Flags().GetString("new")
		err := Change(path, oldString, newString)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("file has been successfully changed")
		}
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
	changeCmd.Flags().StringP("path", "p", "Path", "path to file")
	changeCmd.Flags().StringP("old", "o", "Old", "search string")
	changeCmd.Flags().StringP("new", "n", "New", "string for change")
}

func Change(path, old, new string) error {
	currentContent, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	newContent := strings.ReplaceAll(string(currentContent), old, new)
	if string(currentContent) == newContent {
		return fmt.Errorf("string not found")
	}

	err = os.WriteFile(path, []byte(newContent), fs.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	Execute()
}
