package asciiArtColor

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// converts HSL values to RGB values
func HSLtoRGB(hsl string) []string {
	var Crgb []int
	num := ""
	for i := 0; i < len(hsl); i++ {
		if IsNumeric(string(hsl[i])) {
			num = num + string(hsl[i])
		} else if !(IsNumeric(string(hsl[i]))) && num != "" {
			numCap, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error during hsl.")
				os.Exit(0)
			}
			Crgb = append(Crgb, numCap)
			num = ""
		} else {
			continue
		}
	}
	h := Crgb[0]
	s := Crgb[1]
	l := Crgb[2]
	//hsl
	s2 := float64(s) / float64(100)
	l2 := float64(l) / float64(100)
	c := (1 - math.Abs(2*l2-float64(1))) * s2
	x := c * (float64(1) - math.Abs(math.Mod(float64(h)/60, 2)-float64(1)))
	m := l2 - c/2
	var r, g, b float64
	switch {
	case h < 60:
		r = c
		g = x
		b = 0
	case h < 120 && h >= 60:
		r = x
		g = c
		b = 0
	case h < 180 && h >= 120:
		r = 0
		g = c
		b = x
	case h < 240 && h >= 180:
		r = 0
		g = x
		b = c
	case h < 300 && h >= 240:
		r = x
		g = 0
		b = c
	case h < 360 && h >= 300:
		r = c
		g = 0
		b = x
	}
	R := int(math.Round((r + m) * 255))
	G := int(math.Round((g + m) * 255))
	B := int(math.Round((b + m) * 255))
	Rr := strconv.FormatInt(int64(R), 10)
	Gg := strconv.FormatInt(int64(G), 10)
	Bb := strconv.FormatInt(int64(B), 10)
	arr := []string{string(Rr), string(Gg), string(Bb)}
	return arr
}

// HextoRGB converts a hexadecimal string to RGB values
func HextoRGB(hex string) []string {
	if hex[0:1] == "#" {
		hex = hex[1:]
	}
	r := string(hex)[0:2]
	g := string(hex)[2:4]
	b := string(hex)[4:6]
	R, _ := strconv.ParseInt(r, 16, 0)
	G, _ := strconv.ParseInt(g, 16, 0)
	B, _ := strconv.ParseInt(b, 16, 0)
	Rr := strconv.FormatInt(int64(R), 10)
	Gg := strconv.FormatInt(int64(G), 10)
	Bb := strconv.FormatInt(int64(B), 10)
	arr := []string{string(Rr), string(Gg), string(Bb)}
	return arr
}

// RGBtoNum
func RGBtoNum(hex string) []string {
	var Crgb []string
	num := ""
	for i := 0; i < len(hex); i++ {
		if IsNumeric(string(hex[i])) {
			num = num + string(hex[i])
		} else if !(IsNumeric(string(hex[i]))) && num != "" {
			numCap, err := strconv.Atoi(num)
			if err != nil {
				fmt.Println("Error during conversion in the Capitlizee.")
				os.Exit(0)
			}
			if numCap > 255 {
				fmt.Println("wrong rgb num.")
				os.Exit(1)
			} else {
				Crgb = append(Crgb, num)
				num = ""
			}

		} else {
			continue
		}
	}
	return Crgb
}
