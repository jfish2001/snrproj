package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jfish2001/snrproj/tree/master/monkey-interpreter-master/eval"
	"github.com/jfish2001/snrproj/tree/master/monkey-interpreter-master/lexer"
	"github.com/jfish2001/snrproj/tree/master/monkey-interpreter-master/object"
	"github.com/jfish2001/snrproj/tree/master/monkey-interpreter-master/parser"
)

const prompt = ">> "

// Start starts Monkey REPL.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print("Hello! This is the Fisher programming language!")
		fmt.Println("Feel free to type in commands")
		if !scanner.Scan() {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program, env)
		if evaluated == nil {
			continue
		}

		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, msg)
		io.WriteString(out, "\n")
	}
}
