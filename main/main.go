package main

import (
	"asciiArtOutPut"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

// check amount of arguments
func main() {
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println(len(os.Args), "is Not a valid amount of arguments.")
		return
	}
	substr := "--output="
	str := os.Args[1]
	for i := 0; i < len(str)-len(substr)+1; i++ {
		if str[i:i+len(substr)] == substr && len(os.Args) == 3 {
			os.Args = append(os.Args, "standard")
		} else {
			continue
		}
	}
	args := os.Args[1:]
	ArgsLen := len(args)
	if !(asciiArtOutPut.IsValid(args[0])) {
		fmt.Println("This's Not a valid character.")
		return
	}

	font := "standard" //base font
	outputFileName := ""
	text := ""
	text = args[0] // "hello" == [0]

	if ArgsLen == 3 || ArgsLen == 2 {

		switch args[ArgsLen-1] {
		case "shadow":
			font = "shadow"
		case "thinkertoy":
			font = "thinkertoy"
		case "standard":
			font = "standard"
		default:
			fmt.Println(args[ArgsLen-1], "is Not a valid font.")
			return
		}
	}
	if len(os.Args) == 4 {
		text = args[0+1] // "hello" == [0]
		output := args[0]
		outputFile := strings.Split(output, "--output=")
		if len(outputFile) == 1 {
			fmt.Println("you should print the run like this example: ")
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
		}

		outputFileName = outputFile[1]
		if outputFile[0] != "" {
			fmt.Println("tf wrong w u.")
			return
		}
		NameLen := len(outputFileName)
		if NameLen < 5 {
			fmt.Println(outputFileName, "is Not a valid output File Name.")
			return
		} else if !(outputFileName[NameLen-1] == 't' && outputFileName[NameLen-2] == 'x' && outputFileName[NameLen-3] == 't' && outputFileName[NameLen-4] == '.') {
			fmt.Println("output File Name should end with <.txt> .")
			return
		}
	} else {

		text = args[0] // "hello" == [0]
		if len(args) == 2 {
			switch args[1] {
			case "shadow":
				font = "shadow"
			case "thinkertoy":
				font = "thinkertoy"
			case "standard":
				font = "standard"
			default:
				fmt.Println("Not a valid font")
				return
			}
		}

	}

	// Read the content of the file
	text = strings.ReplaceAll(text, "\\t", "   ")
	argsArr := strings.Split(strings.ReplaceAll(text, "\\n", "\n"), "\n")
	arr := []string{}
	readFile, err := os.Open("fonts/" + font + ".txt")
	defer readFile.Close()

	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		arr = append(arr, fileScanner.Text())
	}

	if len(arr) != fileLen {
		fmt.Println("File is corrupted.")
		return
	}
	larg := len(argsArr)
	if larg >= 2 {
		if argsArr[larg-1] == "" && argsArr[larg-2] != "" {
			argsArr = argsArr[:larg-1]
		}
	}
	if outputFileName != "" {
		asciiArtOutPut.PrintBannersInFile(outputFileName, argsArr, arr)
	} else {
		asciiArtOutPut.PrintBanners(argsArr, arr)
	}
}
