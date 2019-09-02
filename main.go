package main

import (
	"./lexer"
	"fmt"
)

var (
	//UPPERS = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//LOWERS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	DIGITS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//SPACES = []string{" ", "\n", "\t"}
	//OETHERS =[]string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "`", "{", "|", "}", "~"}

	//CLASSES = map[]
)

func main() {
	b := `\d{1}`
	l := lexer.New(b)

	for l.HasNext() {
		fmt.Println(l.NextToken())
	}
	fmt.Println(l.NextToken())
}
