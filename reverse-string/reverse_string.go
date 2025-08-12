package reverse

func Reverse(input string) string {
	/*
		A string is a sequence of UTF-8 bytes, and some characters (like こ, é) use multiple bytes.
		Indexing or reversing by bytes can split those bytes apart, corrupting the text.

		when we work with runes: Each rune = one full Unicode code point (one character).
		Indexing works per character, so reversing doesn’t break multi-byte ones.
	*/
	r := []rune(input)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		// go lets u swap like this without a temp value
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
