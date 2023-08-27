package asciiArtColor

import (
	"fmt"
	// "golang.org/x/image/colornames"
	"bufio"
	"log"
	"os"
)

const (
	fileLen = 144
)

// Print the full outcome in the triminal
func WordColors(color string) []string {
	edited := ""
	for i := 0; i < len(color); i++ {
		if color[i] != ' ' {
			edited = edited + string(color[i])
		}
	}
	color, edited = edited, color
	arr := []string{}
	var Crgb []string
	readFile, err := os.Open("colortt.txt")
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
		os.Exit(1)
	}
	match := false
	for j := 0; j < len(color); j++ {
		if !(IsAlpha(string(color[j]))) {
			fmt.Println("Error: wrong color name/type.")
			os.Exit(0)
		}
	}
	for i := 0; i < len(arr); i++ {
		word := arr[i]
		if color == word[0:len(color)] && len(word) == len(color)+9 {
			for j := 0; j < 3; j++ {
				num := (len(color)) + (3 * j)
				Crgb = append(Crgb, (word[num : num+3]))
			}
			match = true
		}

	}
	if match == false {
		fmt.Println("\033[1;31m"+edited, "isn't a color.\033[0m")
		os.Exit(1)
	}
	return Crgb
}
