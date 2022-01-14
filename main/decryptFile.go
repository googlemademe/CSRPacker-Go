package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//Function to decrypt the files in the "Original" folder and creates a json text file in "Decrypted" folder
func decryptFiles(startFolder string, endFolder string) {

	//FOR loop to read all the files in the "Original" folder and create a json text file in "Decrypted" folder
	for _, fileName := range readFolder(startFolder) {

		//Open a file in the "Original" folder
		gzipFile, perr := os.Open(fileName)

		//IF statement to check if there is any exception and exit the program
		if perr != nil {
			fmt.Println(perr)
			os.Exit(-1)
		}

		// Create a gzip reader on top of the file reader
		// Again, it could be any type reader though
		gzipReader, err := gzip.NewReader(gzipFile)

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}

		//Closes the gzip Reader
		defer gzipReader.Close()

		// Create a buffer reader to skip the first line
		scanner := bufio.NewReader(gzipReader)

		//Reads the first line
		scanner.ReadLine()

		//Defer the closing of our jsonFile so that we can parse it later on
		defer gzipFile.Close()

		//Read all the data from the file from the IO buffer
		byteValue, _ := ioutil.ReadAll(scanner)

		//Gets the current full working directory path
		currentDIR, ierr := os.Getwd()

		//IF statement to check if there is any exception and exit the program
		if ierr != nil {
			fmt.Println("Fatal Error cannot access directory to save the file....")
			fmt.Println("Exiting....")
			fmt.Println(ierr)
			os.Exit(3)
		}

		//Creating the "Decrypted" folder full directory path
		decryptedFile := currentDIR + endFolder

		//Getting the filename for the file
		fileName := filepath.Base(fileName)

		//Function call to indent the json data string for human readability
		indentedString, err := PrettyString(string(byteValue))

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Println("Fatal Error while indenting the JSON data string....")
			fmt.Println("Exiting....")
			log.Fatal(err)
			os.Exit(3)
		}
		//Write the json string to a text file
		_ = ioutil.WriteFile(decryptedFile+fileName+".txt", []byte(indentedString), 0644)

		//Prints out a successful write notice
		fmt.Println("Successfully Converted " + filepath.FromSlash(decryptedFile+fileName))
	}
}
