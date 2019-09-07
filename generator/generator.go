package generator

import (
	"bytes"
	"fmt"
	"github.com/cipepser/stringRandom/ast"
	"math/rand"
	"time"
)

const (
	INFINITE = 100
)

var (
	UPPERS = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	LOWERS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	DIGITS = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	SPACES = []string{" ", "\n", "\t"}
	OTHERS = []string{"!", "\"", "#", "$", "%", "&", "'", "(", ")", "*", "+", ",", "-", ".", "/", ":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "`", "{", "|", "}", "~"}

	WORD      []string
	ANY       []string
	NOTDIGITS []string
	NOTWORD   []string
	NOTSPACES []string
)

func init() {
	WORD = append(WORD, UPPERS...)
	WORD = append(WORD, LOWERS...)
	WORD = append(WORD, DIGITS...)
	WORD = append(WORD, "_")

	ANY = append(ANY, UPPERS...)
	ANY = append(ANY, LOWERS...)
	ANY = append(ANY, DIGITS...)
	ANY = append(ANY, OTHERS...)
	ANY = append(ANY, "_")

	NOTDIGITS = append(NOTDIGITS, UPPERS...)
	NOTDIGITS = append(NOTDIGITS, LOWERS...)
	NOTDIGITS = append(NOTDIGITS, OTHERS...)
	NOTDIGITS = append(NOTDIGITS, "_")

	NOTWORD = append(NOTWORD, OTHERS...)

	NOTSPACES = append(NOTSPACES, UPPERS...)
	NOTSPACES = append(NOTSPACES, LOWERS...)
	NOTSPACES = append(NOTSPACES, DIGITS...)
	NOTSPACES = append(NOTSPACES, OTHERS...)
	NOTSPACES = append(NOTSPACES, "_")
}

func Generate(node ast.Node) {
	switch node := node.(type) {
	case *ast.Program:
		generateProgram(node)
	case *ast.ExpressionStatement:
		generateExpressionStatement(node)
	case *ast.NumberExpression:
		generateNumberExpression(node)
	case *ast.DigitExpression:
		generateDigitExpression(node)
	case *ast.NotDigitExpression:
		generateNotDigitExpression(node)
	case *ast.StringExpression:
		generateStringExpression(node)
	case *ast.WordExpression:
		generateWordExpression(node)
	case *ast.NotWordExpression:
		generateNotWordExpression(node)
	case *ast.SpaceExpression:
		generateSpaceExpression(node)
	case *ast.NotSpaceExpression:
		generateNotSpaceExpression(node)
	case *ast.NewlineExpression:
		generateNewlineExpression(node)
	case *ast.TabExpression:
		generateTabExpression(node)
	case *ast.BackslashExpression:
		generateBackslashExpression(node)
	case *ast.DotExpression:
		generateDotExpression(node)
	case *ast.BlockExpression:
		generateBlockExpression(node)
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

func generateNumberExpression(node *ast.NumberExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(node.TokenLiteral())
	}
	fmt.Print(out.String())
}

func generateDigitExpression(node *ast.DigitExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomDigit())
	}
	fmt.Print(out.String())
}

func generateRandomDigit() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(DIGITS))

	return DIGITS[r]
}

func generateNotDigitExpression(node *ast.NotDigitExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomNotDigit())
	}
	fmt.Print(out.String())
}

func generateRandomNotDigit() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(NOTDIGITS))

	return NOTDIGITS[r]
}

func generateStringExpression(node *ast.StringExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(node.TokenLiteral())
	}
	fmt.Print(out.String())
}

func generateWordExpression(node *ast.WordExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomLetter())
	}
	fmt.Print(out.String())
}

func generateRandomLetter() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(WORD))

	return WORD[r]
}

func generateNotWordExpression(node *ast.NotWordExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateNotRandomLetter())
	}
	fmt.Print(out.String())
}

func generateNotRandomLetter() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(NOTWORD))

	return NOTWORD[r]
}

func generateSpaceExpression(node *ast.SpaceExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomSpace())
	}
	fmt.Print(out.String())
}

func generateRandomSpace() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(SPACES))

	return SPACES[r]
}

func generateNotSpaceExpression(node *ast.NotSpaceExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateNotRandomSpace())
	}
	fmt.Print(out.String())
}

func generateNotRandomSpace() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(NOTSPACES))

	return NOTSPACES[r]
}

func generateNewlineExpression(node *ast.NewlineExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString("\n")
	}
	fmt.Print(out.String())
}

func generateTabExpression(node *ast.TabExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString("\t")
	}
	fmt.Print(out.String())
}

func generateBackslashExpression(node *ast.BackslashExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString("\\")
	}
	fmt.Print(out.String())
}

func generateDotExpression(node *ast.DotExpression) {
	var out bytes.Buffer
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		out.WriteString(generateRandomAny())
	}
	fmt.Print(out.String())
}

func generateRandomAny() string {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(len(ANY))

	return ANY[r]
}

func generateBlockExpression(node *ast.BlockExpression) {
	rand.Seed(time.Now().UnixNano())

	max := node.Range.Max
	if max == ast.INFINITE {
		max = INFINITE
	}

	n := rand.Intn(max-node.Range.Min+1) + node.Range.Min
	for i := 0; i < n; i++ {
		Generate(&node.Block)
	}
}
