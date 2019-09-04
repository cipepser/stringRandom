package repl

import (
	"bufio"
	"fmt"
	"github.com/cipepser/stringRandom/generator"
	"github.com/cipepser/stringRandom/lexer"
	"github.com/cipepser/stringRandom/parser"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		program := parser.New(l)
		generator.Generate(program.Parse())

		//io.WriteString(out, "\n")
	}
}
