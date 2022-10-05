package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {

	content, err := os.Open("Hangman.txt") // Open the file Hangman in .txt

	if err != nil { // If an error occurs, print err
		log.Fatal(err)
		fmt.Println("err")
	}

	scanner := bufio.NewScanner(content) // Scan the file
	scanner.Split(bufio.ScanWords)
	tmp := make([]string, 0) // Store all the strings in the file in a slice
	for scanner.Scan() {
		tmp = append(tmp, scanner.Text())

	}

	min := 0
	max := 100
	fmt.Println(tmp[0])
	fmt.Println(tmp[len(tmp)-1])
	fmt.Println(Fragment3(tmp))
	fmt.Println(Fragment4(tmp))
	fmt.Println(rand.Intn(max-min) + min)

	content.Close() // Close the file when the program is finished

}

func Fragment3(arr []string) string {
	var oui string
	for index, i := range arr { // Loop in our slice
		if i == "before" {
			result, _ := strconv.Atoi(arr[index+1]) // Convert string to int
			oui = arr[result-1]
		}

	}
	return oui
}

func Fragment4(arr []string) string {
	var res string
	for index, j := range arr { // Loop in our slice
		if j == "now" {
			l := arr[index-1]         // Take the previous string
			i := int(l[1]) / len(arr) // Divide the rune by the size of our slice
			res = arr[i-1]            // Take the result minus one in our slice
		}
	}
	return res
}
