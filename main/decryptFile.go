package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Function to decrypt the files in the "Original" folder and creates a json text file in "Decrypted" folder
func decryptFiles(startFolder string, endFolder string) {

	fmt.Println()
	//FOR loop to read all the files in the "Original" folder and create a json text file in "Decrypted" folder
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

		//Creating the "Decrypted" folder full directory path
		decryptedFile := filepath.FromSlash(currentDIR + endFolder)

		//Open a file in the "Original" folder
		gzipFile, err := os.Open(filepath.FromSlash(fileName))

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot open file - %s....\n", filepath.FromSlash(fileName))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		}

		// Create a gzip reader on top of the file reader
		// Again, it could be any type reader though
		gzipReader, err := gzip.NewReader(gzipFile)
		//gzipReader, err := os.Open("users.json")

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot open file - %s....\n", filepath.FromSlash(fileName))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		}

		//Closes the gzip Reader
		defer gzipReader.Close()

		// Create a buffer reader to skip the first line
		scanner := bufio.NewReader(gzipReader)

		//Reads the first line
		scanner.ReadString('\n')

		//Defer the closing of our jsonFile so that we can parse it later on
		defer gzipFile.Close()

		//Read all the data from the file from the IO buffer
		byteValue, err := ioutil.ReadAll(scanner)

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot read file - %s....\n", filepath.FromSlash(fileName))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		}

		//String replace function call to remove backslashes from the json data string
		noSpaceString := strings.Replace(string(byteValue), "\\", "", -1)

		//String replace fucntion call to remove NULL characters
		noSpaceString = strings.Replace(noSpaceString, "\x00", "", -1)

		//Function to indent to json data string for human readability
		prettyJsonString, err := PrettyString(noSpaceString)

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Println("Fatal Error cannot beautify the json string....")
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		}

		//Getting the filename for the file
		fileName := filepath.Base(fileName)

		//byteValue, _ := ioutil.ReadAll(out)
		//Write the json string to a text file
		err = ioutil.WriteFile(filepath.FromSlash((decryptedFile + fileName + ".txt")), []byte(prettyJsonString), 0644)

		//IF statement to check if there is any exception and exit the program
		if err != nil {
			fmt.Printf("Fatal Error cannot write file - %s....\n", filepath.FromSlash(fileName))
			fmt.Println("Exiting....")
			fmt.Println(err)
			os.Exit(3)
		} else {
			//Prints out a successful write notice
			fmt.Println("Successfully Converted " + filepath.FromSlash(decryptedFile+fileName))
		}
	}
	fmt.Println()
}
