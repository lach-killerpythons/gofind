// GOFINDER
// search for string [Arg1] in folder [Arg2]
package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"fmt"
)

const (
	ColorBlack  = "\u001b[30m"
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorBlue   = "\u001b[34m"
	ColorReset  = "\u001b[0m"
)

func colorize(myColor string, myMessage string) {
	fmt.Println(myColor, myMessage)
}

const targetFileType = ".go"

var hitBool = false

func fileSearcher(myFile string, myStr string) {

	var myHits []string

	f, err := os.Open(myFile)
	counter := 0

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	//fmt.Println(&scanner, "buffon")

	for scanner.Scan() {

		nextLine := scanner.Text()

		if strings.Contains(nextLine, myStr) {
			hitStr := " [LINE:" + strconv.Itoa(counter) + "]" + `"` + nextLine + `"`
			myHits = append(myHits, hitStr)
		}
		counter += 1

	}
	if len(myHits) > 0 {
		//fmt.Println(myFile, "--", myHits)
		fmt.Println(ColorYellow, myFile)
		for _, j := range myHits {
			fmt.Println(ColorGreen, j)
			hitBool = true
			//fmt.Println("\n")
		}
		fmt.Println("\n")
	}

}

func setDir(input string) error {
	var b1 error = nil
	b1 = os.Chdir(input)
	return b1
}

func myFiles(myFileSuffix string) []string { // unused function to get files from a single folder
	var myFiles []string
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		//fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() != true {
			if strings.Contains(file.Name(), myFileSuffix) {
				myFiles = append(myFiles, file.Name())
			}
		}
	}

	return myFiles

}

func getFilesTyp(myFileSuffix string) []string { //walk current directory for all files of .suffix
	var myFiles []string
	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() != true {
				// get current working directory + add to myFile list
				if strings.Contains(info.Name(), myFileSuffix) {
					ff, _ := os.Getwd()
					myFiles = append(myFiles, ff+`\`+path)

				}
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return myFiles // return list of files
}

func main() { // search for string [Arg1] in folder [Arg2]
	fmt.Println("starting!")
	argsUno := ""
	argsDos := ""
	nArgs := len(os.Args) - 1
	if nArgs >= 2 {
		argsUno = os.Args[1]
		argsDos = os.Args[2]
	}

	if argsUno != "" {
		fmt.Println("Finding String -", argsUno)
	} else {
		fmt.Println(ColorBlue, `search for "string" [Arg1] in "folder" [Arg2]`)
	}
	if argsDos != "" {
		fmt.Println(ColorBlue, "Searching Directory -", argsDos)
		if setDir(argsDos) == nil { // actually change the directory to Arg[2]
			myDir, _ := os.Getwd() // actually sets this directory if valid
			fmt.Println(ColorBlue, myDir, "is a valid dir 📁✔️")
			myF := getFilesTyp(targetFileType)
			for _, f := range myF {
				fileSearcher(f, argsUno)
			}

		} else {
			fmt.Println(ColorRed, argsDos, "is an invalid dir 🗑️🚫")
		}
	}

	if !hitBool {
		fmt.Println(ColorRed, "No Results!")
	}

	fmt.Println(ColorReset, "All Done!")

}
