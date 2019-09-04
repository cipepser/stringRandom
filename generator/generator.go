package generator

import (
	"bytes"
	"fmt"
	"github.com/cipepser/stringRandom/ast"
	"math/rand"
	"strconv"
	"time"
)

var (
	//UPPERS = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	//LOWERS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	DIGITS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	//SPACES = []string{" ", "\n", "\t"}
	//OETHERS =[]string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "`", "{", "|", "}", "~"}

	//CLASSES = map[]
)

func Generate(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		generateProgram(node)
	case *ast.ExpressionStatement:
		generateExpressionStatement(node)
	case *ast.DigitExpression:
		generateDigitExpression(node)
	default:
		panic("unknown node" + node.String())
	}
}

func generateProgram(program *ast.Program) {
	for _, statement := range program.Statements {
		Generate(statement)
	}
}

func generateExpressionStatement(node *ast.ExpressionStatement) {
	Generate(node.Expression)
}

func generateDigitExpression(node *ast.DigitExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(node.Range.Max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomDigit())
	}
	fmt.Println(out.String())
}

func generateRandomDigit() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(DIGITS))

	// TODO: DIGITSから取ってこないとダメなのでは...
	// それならstrconvいらないはず

	return strconv.Itoa(r)
}