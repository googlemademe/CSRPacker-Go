package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
)

//Function to delete all files in all directories
func deleteAllFiles(folderArray []string) {

	//FOR loop to recursively walk the directory
	for _, folderName := range folderArray {
		//get the current working directory
		dir, _ := ioutil.ReadDir(folderName)
		//FOR loop for all the files in the current working directory
		for _, fileName := range dir {
			//removes the file in the current working directory
			os.RemoveAll(path.Join([]string{folderName, fileName.Name()}...))
			//display successfull deletion to window
			fmt.Println(filepath.FromSlash((path.Join([]string{folderName, fileName.Name()}...))))
		}
	}
}
