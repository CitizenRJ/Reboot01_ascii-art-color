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
		fmt.Println("check")
		num := 0
		help := 0
		for i := 0; i < ArgsLen; i++ {
			num = num + 1
			if strings.Contains(args[i], StrFlagArr[0]) {
				outputFile = outputFileCheck(args[i])
				if ColorColor == "" {
					if i+1 < ArgsLen {
						haha = append(haha[:num-1], haha[num:]...)
						num = num - 1
					} else {
						haha = haha[:num-1]
					}

					help = help + 1
				} else {
					outputFile = ""
					continue
				}
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				if outputFile == "" {
					if i+1 < ArgsLen {
						fmt.Println("try",haha)
						haha = append(haha[:num-1], haha[num:]...)
						fmt.Println("try",haha)
						num = num - 1
					} else {
						haha = haha[:num-1]
					}
					help = help + 1
				} else {
					continue
				}
			}
		}
		if help == 0 {

		}
	} else if ArgsLen > 2 && ArgsLen < 6 {
		num := 0
		help := 0
		for i := 0; i < ArgsLen; i++ {
			num = num + 1
			if strings.Contains(args[i], StrFlagArr[0]) {
				if ColorColor == "" {
					outputFile = outputFileCheck(args[i])
					if i+1 < ArgsLen {
						haha = append(haha[:num-1], haha[num:]...)
					} else {
						haha = haha[:num-1]
					}
					num = num - 1
				} else {
					continue
				}
				help = (help + 1)
			} else if strings.Contains(args[i], StrFlagArr[1]) {
				ColorColor = ColorColorCheck(args[i])
				if outputFile == "" {
					haha = append(haha[:num-1], haha[num:]...)
					num = num - 1
					if i+1 < ArgsLen {
						Str = haha[num]
						haha = append(haha[:num], haha[num+1:]...)
						num = num - 1
					} else {
						haha = haha[:num]
					}
				} else if outputFile != "" {
					if i+1 < ArgsLen {
						haha = append(haha[:num], haha[num+1:]...)
					} else {
						haha = haha[:num]
					}
					num = num - 1
				} else if i+1 < ArgsLen {
					Str = args[i+1]
					if i+2 < ArgsLen {
						haha = append(haha[:num], haha[num+2:]...)
						num = num - 2
					} else {
						haha = haha[:num+1]
					}
				}
				help = help + 1
			}
		}
		if help == 0 {
			fmt.Println("Error: Invalid arguments.")
			return
		}
	} else if ArgsLen > 5 {
		fmt.Println("Error: Invalid arguments.")
		return
	}
	ArgsLen = len(haha)
	// fmt.Println(haha, "4")
	if ArgsLen == 0 && Str != "" {
		haha = append(haha, Str)
		haha = append(haha, "standard")
		ArgsLen = len(haha)
		Str = ""
	} else if ArgsLen == 0 && Str == "" {
		
		haha = append(haha, os.Args[1])
		haha = append(haha, "standard")
		ArgsLen = len(haha)
	} else if ArgsLen == 1 {
		if Str != "" && Fonts(haha[ArgsLen-1]) {
			lol := haha[ArgsLen-1]
			haha[ArgsLen-1] = Str
			haha = append(haha, lol)
			fmt.Println(haha)
			Str = ""
			font = lol
		} else {
			// fmt.Println("1",haha)
		haha = append(haha, "standard")
		}
		
		ArgsLen = len(haha)
	} else if ArgsLen == 2 {
		font = haha[ArgsLen-1]
		if Fonts(haha[ArgsLen-1]) {
			font = haha[ArgsLen-1]

		} else {
			fmt.Println("\""+font+"\"", "is Not a valid font.")
			os.Exit(0)
		}
	} else {
		fmt.Println("Error: Invalid arguments.", haha)
		return
	}
	
	for i := 0; i < len(args); i++ {
		if !(asciiArtColor.IsValid(args[i])) {
			fmt.Println(args[i], "isn't a valid character/argument.")
			return
		}
	}

	text = haha[ArgsLen-2]
	if len(Str) > len(text) {
		fmt.Println("the \"", Str, "\" should be less or equal than \"", text, "\".")
		return
	} else if !(strings.Contains(text, Str)) {
		fmt.Println("the string", "'"+text+"'", "should contains", "'"+Str+"'", "to Print.")
		os.Exit(1)
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
	outputFile := strings.TrimPrefix(output, "--output=")
	if !strings.HasSuffix(outputFile, ".txt") {
		fmt.Println("Output file name should end with .txt")
		os.Exit(1)
	}
	return outputFile
}

func ColorColorCheck(color string) string {
	textColor := strings.TrimPrefix(color, "--color=")
	return textColor
}
