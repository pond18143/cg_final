package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Print("Enter text: ")
	var input string
	fmt.Scanln(&input)
	fmt.Print(input)

	// Open file on disk.
	f, _ := os.Open("./ramen.png")

	// Read entire JPG into byte slice.
	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)

	// Encode as base64.
	encoded := base64.StdEncoding.EncodeToString(content)

	// Print encoded data to console.
	// ... The base64 image can be used as a data URI in a browser.
	//fmt.Println("ENCODED: " + encoded)

	length := len(encoded)
	a := encoded[0:188]
	b := input
	c := encoded[191:length]

	//fmt.Println("ENCODED: " + a )
	//
	//fmt.Println("ENCODED: " + b )
	//
	//fmt.Println("ENCODED: " + c )
	fmt.Println("Token: " + a + b + c )

}