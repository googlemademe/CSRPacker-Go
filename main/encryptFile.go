package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

//Function to encrypt the files in the "Decrypted" folder and creates a gzip file in "Finished" folder
func encryptFiles(startFolder string, endFolder string) {

	//FOR loop to read all the files in the "Decrypted" folder and create a gzip file in "Finished" folder
	for _, b := range readFolder(startFolder) {

		//Open a file in the "Decrypted" folder
		textFile, perr := os.Open(b)

		//IF statement to check if there is any exception and exit the program
		if perr != nil {
			fmt.Println(perr)
			os.Exit(-1)
		}

		//Create a Reader and use ReadAll to get all the bytes from the file
		reader := bufio.NewReader(textFile)

		//Closes the Reader
		defer textFile.Close()

		//Read all the data from the file from the Reader buffer
		byteValue, _ := ioutil.ReadAll(reader)

		//Variable for JSON UnMarshal variable
		var result map[string]interface{}

		//Reads the JSON data into a json unmarshal variable
		json.Unmarshal([]byte(byteValue), &result)

		//Gets the current full working directory path
		currentDIR, ierr := os.Getwd()

		//IF statement to check if there is any exception and exit the program
		if ierr != nil {
			fmt.Println(ierr)
			os.Exit(-1)
		}

		//Creating the "Finished" folder full directory path
		finishedFile := currentDIR + endFolder

		//Getting the filename for the file
		fileName := filepath.Base(b)

		//Generating the json string from the json unmarshal data
		jsonFile, _ := json.Marshal(result)

		//Creating the encrypted filename
		name := finishedFile + fileName

		//Removing all file extension for the file
		name = strings.Replace(name, ".txt", "", -1)

		//Creating the file
		textFile, _ = os.Create(name)

		//Write compressed json string data
		writeJson := gzip.NewWriter(textFile)

		//Generating the json string with hashkey
		jsonString := GetSignature(string(jsonFile), getHashkey()) + "\r" + string(jsonFile)

		//Write the binary data to the file
		writeJson.Write([]byte(jsonString))

		//Closed the file for writing
		writeJson.Close()

		//Prints out a successful write notice
		fmt.Println("Successfully Saved " + filepath.FromSlash(name))
	}
}
