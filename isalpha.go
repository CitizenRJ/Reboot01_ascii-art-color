package asciiArtColor

func IsAlpha(s string) bool {
	for _, w := range s {
		if isLte(w) == false {
			return false
		}
	}
	return true
}

func isLte(woo rune) bool {
	if woo == ' ' {
		return false
	}
	return (woo >= 'A' && woo <= 'Z') || (woo >= 'a' && woo <= 'z') || (woo >= '0' && woo <= '9')
}
