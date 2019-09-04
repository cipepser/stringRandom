package main

import (
	"fmt"
	"github.com/cipepser/stringRandom/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the RegExp generator!\n", user.Username)
	fmt.Printf("Type some RegExp\n")
	repl.Start(os.Stdin, os.Stdout)
}
