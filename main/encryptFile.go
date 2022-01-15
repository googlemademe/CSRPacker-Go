package main

import (
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//Function to encrypt the files in the "Decrypted" folder and creates a gzip file in "Finished" folder
func encryptFiles(startFolder string, endFolder string) {

	fmt.Println()
	//FOR loop to read all the files in the "Decrypted" folder and create a gzip file in "Finished" folder
	for _, fileName := range readFolder(startFolder) {

		//Gets the current full working directory path
		currentDIR, err := os.Getwd()

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot access directory - %s - to save the file....\n", filepath.FromSlash(currentDIR))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		}

		//Read all the data from the file from the Reader buffer
		content, err := ioutil.ReadFile(fileName)

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot access directory - %s - to save the file....\n", filepath.FromSlash(currentDIR))
			fmt.Println("Exiting....")
			log.Fatal(err)
			os.Exit(3)
		}

		//Creating the "Finished" folder full directory path
		finishedFile := currentDIR + endFolder

		//Getting the filename for the file
		fileName := filepath.Base(fileName)

		//Creating the encrypted filename
		name := finishedFile + fileName

		//Removing all file extension for the file
		name = strings.Replace(name, ".txt", "", -1)

		//Creating the file
		textFile, _ := os.Create(name)

		//Write compressed json string data
		writeJson := gzip.NewWriter(textFile)

		//Generating the json string with hashkey
		jsonString := encryptJsonString(removeSpacesString(content))

		//Write the binary data to the file
		fileSize, err := writeJson.Write([]byte(jsonString))

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot write file - %s....\n", filepath.FromSlash(fileName))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		} else {
			//Prints out a successful write notice
			fmt.Println("Successfully Saved " + filepath.FromSlash(name) + fmt.Sprintf(" - Filesize: %v bytes", fileSize))
		}

		//Closed the file for writing
		writeJson.Close()

	}
	fmt.Println()
}
