package piscine

func StrRev(s string) string {

	runes := []rune(s)
	length := len(runes)
	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}
	reversedString := string(runes)
	return reversedString
}
