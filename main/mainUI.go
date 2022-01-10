package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//Function for Main UI of the program
func mainUI() int {

	//variable for reading user input
	scanner := bufio.NewScanner(os.Stdin)

	//Main UI menu
	fmt.Println("***********************************************************************************************************")
	fmt.Println("                MacOS CSRPacker UI made by AppleCoreOne")
	fmt.Println("                Based Upon CSRPacker UI made by Octa/ERS_")
	fmt.Println("Custom CMD-based UI for use with CSRPacker. Easy decryption and encryption of nsb, scb, trb, and crc files.")
	fmt.Println("           I don't own or develop CSRPacker, only this UI.")
	fmt.Print("***********************************************************************************************************")
	fmt.Println()
	fmt.Println()
	fmt.Println("1 - Decrypt files (\"Original\" folder)")
	fmt.Println("2 - Encrypt files (\"Decrypted\" folder)")
	fmt.Println("3 - Remove all files from folders")
	fmt.Println("4 - Exit")
	fmt.Println()
	fmt.Print("Type your choice: ")
	//scans in user input
	scanner.Scan()

	//converts user input from a string to an int
	usrInput, _ := strconv.Atoi(scanner.Text())

	//return user input
	return usrInput
}
