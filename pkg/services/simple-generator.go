package services

type SimpleGenerator struct {
	Alphabet string
	Len      int
}

/*Gen generate url*/
func (gen *SimpleGenerator) Gen() <-chan string {
	c := make(chan string)
	go func(c chan string) {
		defer close(c)
		addLetter(c, "", gen.Alphabet, gen.Len)
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
