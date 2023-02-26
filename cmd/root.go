/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func isEmptyJSON(b []byte) bool {
	data := make(map[string]interface{})
	if err := json.Unmarshal(b, &data); err != nil {
		log.Println(err)
		return false
	}
	if len(data) != 0 {
		log.Println("not empty json")
		return false
	}
	return true
}

const (
	tmpDir = "emptyjson"
	tmpTXT = "emptyjson.txt"
)

func checkFile(path string, f os.FileInfo, err error) error {
	if err != nil || f.IsDir() {
		return err
	}

	var isEmpty bool
	log.Println("path:", path)
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	defer func() {

		if isEmpty {
			// add filepath to emptyjson.txt
			emptyJSONlistfile.WriteString(path + "\n")

			// move file to emptyjson directory
			filename := filepath.Base(path)
			if err := os.Rename(path, filepath.Join(".", tmpDir, filename)); err != nil {
				log.Println(err)
			}
		}
	}()

	isEmpty = isEmptyJSON(data)
	return nil
}

func emptyJSON(cmd *cobra.Command, args []string) {
	fmt.Println("directory to search:", jsonDir)

	err := filepath.Walk(jsonDir, checkFile)
	if err != nil {
		log.Fatal(err)
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "emptyjson",
	Short: "find empty json file and move to 'emptyjson' directory",
	Long: `find empty json file and move to 'emptyjson' directory.
it also make 'emptyjson.txt' file which contains empty json file list. example:

$ emptyjson --dir=myjsonfiles`,

	Run: emptyJSON,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var jsonDir string
var emptyJSONlistfile *os.File

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.emptyjson.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringVarP(&jsonDir, "dir", "d", "", "Directory to check recursively (required)")
	rootCmd.MarkFlagRequired("jsonDir")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := os.Mkdir(tmpDir, os.ModePerm); err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	var err error
	path := filepath.Join(".", tmpDir, tmpTXT)
	if err := os.Remove(path); err != nil {
		log.Fatal(err)
	}
	emptyJSONlistfile, err = os.OpenFile(filepath.Join(".", tmpDir, tmpTXT), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
