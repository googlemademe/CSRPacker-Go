package main

import (
	"bufio"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func decryptFiles(startFolder string, endFolder string) {
	for _, b := range readFolder(startFolder) {
		gzipFile, perr := os.Open(b)
		if perr != nil {
			log.Fatal(perr)
		}

		// Create a gzip reader on top of the file reader
		// Again, it could be any type reader though
		gzipReader, err := gzip.NewReader(gzipFile)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(gzipReader)
		defer gzipReader.Close()

		// Create a buffer reader to skip the first line
		scanner := bufio.NewReader(gzipReader)
		scanner.ReadLine() // this moves to the next token

		// defer the closing of our jsonFile so that we can parse it later on
		defer gzipFile.Close()

		//byteValue, _ := ioutil.ReadAll(jsonFile)
		byteValue, _ := ioutil.ReadAll(scanner)

		var result map[string]interface{}

		json.Unmarshal([]byte(byteValue), &result)

		pwd, ierr := os.Getwd()
		if ierr != nil {
			log.Fatal(ierr)
		}
		decryptedFile := pwd + endFolder

		fileName := filepath.Base(b)
		jsonFile, _ := json.MarshalIndent(result, "", " ")
		_ = ioutil.WriteFile(decryptedFile+fileName+".txt", jsonFile, 0644)
		fmt.Println("Successfully Converted " + b)
	}
}
