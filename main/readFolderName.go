package main

import (
	"fmt"
	"os"
	"path/filepath"
)

//Function to walk thru the current working directory and return all files in the directory
func readFolder(folderName string) []string {

	//String slice to hold the names of all files in the current working directory
	var nameOfFiles []string

	//Get the current working directory full path
	currentDir, ierr := os.Getwd()

	//IF statement to check if there is any exception and exit the program
	if ierr != nil {
		fmt.Println(ierr)
		os.Exit(-1)
	}

	//Walk the current directory and get all names of the files
	err := filepath.Walk(currentDir+folderName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		//Add the file names to a slice
		nameOfFiles = append(nameOfFiles, path)
		return nil
	})

	//IF statement to check if there is any exception and exit the program
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	//return the file name slice from 1 element to end
	return nameOfFiles[1:]
}
