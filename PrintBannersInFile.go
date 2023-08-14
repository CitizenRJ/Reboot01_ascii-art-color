package asciiArtOutPut

import (
	"bufio"
	"fmt"
	"os"
)

// Print the full outcome in a file.
func PrintBannersInFile(outputFileName string, banners, arr []string) {
	num := 0
	file, errs := os.Create(outputFileName)
	if errs != nil {
		fmt.Println("Failed to create file:", errs)
		os.Exit(2)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)

	for _, ch := range banners {
		num = num + 1
		if ch == "" {
			if num < len(banners) {
				fmt.Fprintln(writer, "")
				continue
			} else {
				continue
			}
		}
		for i := 0; i < 8; i++ {
			for _, j := range ch {
				n := (j-32)*9 + 1
				fmt.Fprint(writer, arr[int(n)+i])

			}
			fmt.Fprintln(writer, "")
		}
	}
	writer.Flush()
	fmt.Println("Wrote to file: " + outputFileName + ".")
	//     colorReset := "\033[0m"

    // colorRed := "\033[31m"
    // colorGreen := "\033[32m"
    // colorYellow := "\033[33m"
    // colorBlue := "\033[34m"
    // colorPurple := "\033[35m"
    // colorCyan := "\033[36m"
    // colorWhite := "\033[37m"
    
    // fmt.Println(string(colorRed), "test")
    // fmt.Println(string(colorGreen), "test")
    // fmt.Println(string(colorYellow), "test")
    // fmt.Println(string(colorBlue), "test")
    // fmt.Println(string(colorPurple), "test")
    // fmt.Println(string(colorWhite), "test")
    // fmt.Println(string(colorCyan), "test", string(colorReset))
    // fmt.Println("next")
}
