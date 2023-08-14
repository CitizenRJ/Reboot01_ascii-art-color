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
	//color.Cyan("file")
}
