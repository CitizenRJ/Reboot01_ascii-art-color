package main

import (
	"asciiArtColor"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	fileLen = 855
)

func main() {
	args := os.Args[1:]
	ArgsLen := len(args)
	font := "standard"
	StrFlagArr := []string{"--output=", "--color="}
	outputFile := ""
	ColorColor := ""
	text := ""
	// var Str []string
	Str := ""
	haha := os.Args[1:]
	// var color []string
	if ArgsLen < 1 {
		fmt.Println(len(os.Args), "is Not a valid amount of arguments.")
		return
	} else if ArgsLen > 0 && ArgsLen < 3 {
		num := 0
		help := 0
		for i := 0; i < ArgsLen; i++ {
			num = num + 1
			if strings.Contains(args[i], StrFlagArr[0]) {
				outputFile = outputFileCheck(args[i])
				haha = append(haha[:num], haha[num+1:]...)
				num = num + 1
				help = help + 1
				// ArgsLen := len(args)
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				haha = append(haha[:num], haha[num+1:]...)
				num = num + 1
				help = help + 1
				// ArgsLen := len(args)
			}
		}
	} else if ArgsLen > 2 && ArgsLen < 6 {
		num := 0
		help := 0
		for i := 0; i < ArgsLen; i++ {
			num = num + 1
			// fmt.Println(i, args[i])
			if strings.Contains(args[i], StrFlagArr[0]) {
				if ColorColor == "" {
					outputFile = outputFileCheck(args[i])
					haha = append(haha[:num], haha[num+1:]...)
					num = num - 1
				} else {
					continue
				}
				help = (help + 1)
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				if outputFile == "" {
					haha = append(haha[:num], haha[num+1:]...)
						num = num - 1
					// fmt.Println("test")
					if i+1 < ArgsLen {
						Str = args[i+1]
						// fmt.Println("test2", haha, i, "num=", num)
						haha = append(haha[:num], haha[num+1:]...)
						num = num - 1
						// fmt.Println("test3", haha, i)
					} else {
						haha = haha[:num]
					}
					// fmt.P
				} else if outputFile != "" {
					//fmt.Println("test")
					haha = append(haha[:num], haha[num+1:]...)
					num = num - 1
				} else if i+1 < ArgsLen {
					Str = args[i+1]
					if i+2 < ArgsLen {
						// fmt.Println("test4")
						haha = append(haha[:num], haha[num+2:]...)
						num = num - 2
					} else {
						haha = haha[:num+1]
					}

					// Str = append(Str,args[i+1])
				}
				help = help + 1
			}
		}
		if help == 0 {
			fmt.Println("Error: Invalid arguments.")
			return
		} else if ArgsLen > 5 {
			fmt.Println("Error: Invalid arguments.")
			return
		}

		ArgsLen = len(haha)
		if ArgsLen == 1 {
			haha = append(haha, "standard")
			ArgsLen = len(haha)
		} else if ArgsLen == 2 {
			if Fonts(haha[ArgsLen-1]) {
				font = haha[ArgsLen-1]
			} else {
				fmt.Println(font, "is Not a valid font.", haha)
				os.Exit(0)
			}
		} else {
			fmt.Println("Error: Invalid arguments.", haha)
			return
		}

		for i := 0; i < ArgsLen; i++ {
			if !(asciiArtColor.IsValid(args[i])) {
				fmt.Println(args[i], "isn't a valid character/argument.")
				return
			}
		}

		text = haha[ArgsLen-2]
		fmt.Println(text, "999", haha)
		if len(Str) > len(text) {
			fmt.Println("the \"", Str, "\" should be less or equal than \"", text, "\".")
			return
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
		if outputFile != "" {
			asciiArtColor.PrintBannersInFile(outputFile, argsArr, arr)
		} else if ColorColor != "" {
			asciiArtColor.PrintBannersWithColors(Str, ColorColor, argsArr, arr)
		} else {
			asciiArtColor.PrintBanners(argsArr, arr)
		}
	}
}

func Fonts(argFont string) bool {
	// font := ""
	switch argFont {
	case "shadow":
		argFont = "shadow"
	case "thinkertoy":
		argFont = "thinkertoy"
	case "standard":
		argFont = "standard"
	default:
		return false
	}
	return true
}

// outputFile Error manegment.
func outputFileCheck(output string) string {
	outputFile := strings.Split(output, "--output=")
	NameLen := len(outputFile)
	ErrorMsg := "you should print the run like this example: \nEX: go run . --output=<fileName.txt> something standard."
	outputFileName := ""
	if outputFile[0] != "" {
		fmt.Println(ErrorMsg)
		os.Exit(1)
	} else if NameLen > 2 {
		fmt.Println(ErrorMsg)
		os.Exit(1)
	} else if NameLen == 2 && outputFile[1] != "" {
		outputFileName = outputFile[1]
		lenlen := len(outputFileName)
		if lenlen < 5 {
			fmt.Println(outputFileName, "is Not a valid output File Name.")
			os.Exit(1)
		} else if !(outputFileName[lenlen-1] == 't' && outputFileName[lenlen-2] == 'x' && outputFileName[lenlen-3] == 't' && outputFileName[lenlen-4] == '.') {
			fmt.Println("output File Name should end with <.txt> .")
			os.Exit(1)
		}
	}
	outputFileName = outputFile[1]
	return outputFileName
}

func ColorColorCheck(color string) string {
	textColor := strings.Split(color, "--color=")
	NameLen := len(textColor)
	ErrorMsg := "you should print the run like this example: \nEX: go run . --color=<fileName.txt> something standard."
	colors := ""
	if textColor[0] != "" {
		fmt.Println(ErrorMsg)
		os.Exit(1)
	} else if NameLen > 2 {
		fmt.Println(ErrorMsg)
		os.Exit(1)
	} else if NameLen == 2 && textColor[1] != "" {
		colors = textColor[1]
	}
	return colors
}
