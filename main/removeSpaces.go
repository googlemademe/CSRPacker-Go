package main

import (
	"strings"
)

func removeSpacesString(byteValue []byte) []byte {

	//String replacement calls to remove all newlines, carriage returns, null, and whitespaces
	noSpaceString := strings.Replace(string(byteValue), "\n", "", -1)
	noSpaceString = strings.Replace(noSpaceString, "\r", "", -1)
	noSpaceString = strings.Replace(noSpaceString, "  ", "", -1)
	noSpaceString = strings.Replace(noSpaceString, "\x00", "", -1)

	return []byte(noSpaceString)
}
