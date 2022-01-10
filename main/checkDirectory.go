package main

import "os"

//Function to detects if the required directories - "Original", "Decrypted", and "Finished" folder - exists.
//If not, they are created
func checkIfDirectoryExist(path string) {
	//IF statement to check if the folder exist
	if _, err := os.Stat(path); os.IsNotExist(err) {
		//If does not exist, creates the folder
		err := os.Mkdir(path, 0755)
		//Any error that is present will exit the program
		if err != nil {
			os.Exit(-1)
		}
	}
}
