package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//Main function to start the wrapper program
func main() {

	//Constant variables for required working folders
	const originalFolder = "/Original/"
	const decryptedFolder = "/Decrypted/"
	const finishedFolder = "/Finished/"

	//Getting the full path name for the current application directory
	currentDIR, _ := os.Getwd()

	// String array to hold the full path for the working directories
	folderArray := []string{filepath.FromSlash(strings.TrimSuffix((currentDIR + originalFolder), "/")), filepath.FromSlash(strings.TrimSuffix((currentDIR + decryptedFolder), "/")), filepath.FromSlash(strings.TrimSuffix((currentDIR + finishedFolder), "/"))}

	//FOR loop to check if required directores are present, if not it will create them
	for _, b := range folderArray {
		checkIfDirectoryExist(b)
	}

	for {
		//Main UI for user selection
		userInput := mainUI()

		//SWITCH statement to
		switch userInput {
		case 1:
			//Decrypts all files in the "Original" folder
			decryptFiles(originalFolder, decryptedFolder)
		case 2:
			//Encrypts all files in the "Finished" folder
			encryptFiles(decryptedFolder, finishedFolder)
		case 3:
			//Delete all files in the directories
			deleteAllFiles(folderArray)
		case 4:
			//Screen messages that the user wants to exit the program
			fmt.Println("Exiting...... Program......")
			fmt.Println("Thank you......")

			//OS call to close the program successfully
			os.Exit(0)
		default:
			//Wrong user input - the program will exit with an error
			fmt.Println("Invalid input...... Exiting......")
			os.Exit(3)
		}
	}

}
