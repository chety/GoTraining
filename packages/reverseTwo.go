package packages

func reverseString(word string) string  {
	copyWord := []rune(word)
	for i,j := 0, len(copyWord) -1; i < len(copyWord) / 2 ; i , j = i + 1 , j - 1  {
		copyWord[i],copyWord[j] = copyWord[j],copyWord[i]
	}
	return string(copyWord)
}
