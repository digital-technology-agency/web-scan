package models

/*Gen generate url*/
func Gen(alphabet string, len int) <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		addLetter(c, "", alphabet, len)
	}(c)
	return c
}

func addLetter(c chan string, combo string, alphabet string, len int) {
	if len <= 0 {
		return
	}
	var newComb string
	for _, ch := range alphabet {
		newComb = combo + string(ch)
		c <- newComb
		addLetter(c, newComb, alphabet, len-1)
	}
}
