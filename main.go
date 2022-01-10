package main

import (
	"os"
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
	var folderArray = []string{currentDIR + originalFolder, currentDIR + decryptedFolder, currentDIR + finishedFolder}

	//FOR loop to check if required directores are present, if not it will create them
	for _, b := range folderArray {
		checkIfDirectoryExist(b)
	}

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
		//The program will safely exit
		os.Exit(0)
	default:
		//Wrong user input - the program will exit with an error
		os.Exit(-1)
	}

	//Safety check to successfully exit the program
	os.Exit(0)

}
