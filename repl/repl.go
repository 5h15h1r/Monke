package repl

import (
	"bufio"
	"fmt"
	"io"
	"monke/evaluator"
	"monke/lexer"
	"monke/object"
	"monke/parser"
)

const PROMPT = ">>"
const MONKE = `
_
,.-" "-.,
/   ===   \
/  =======  \
__|  (o)   (0)  |__      
/ _|    .---.    |_ \         
| /.----/ O O \----.\ |       
\/     |     |     \/        
|                   |            
|                   |           
|                   |          
_\   -.,_____,.-   /_         
,.-"  "-.,_________,.-"  "-.,
/          |       |          \  
|           l.     .l           | 
|            |     |            |
l.           |     |           .l             
|           l.   .l           | \,     
l.           |   |           .l   \,    
|           |   |           |      \,  
l.          |   |          .l        |
|          |   |          |         |
|          |---|          |         |
|          |   |          |         |
/"-.,__,.-"\   /"-.,__,.-"\"-.,_,.-"\
|            \ /            |         |
|             |             |         |
\__|__|__|__/ \__|__|__|__/ \_|__|__/ 
`

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}

	}
}

func printParseErrors(out io.Writer, errors []string) {

	io.WriteString(out, MONKE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")

	for _, msg := range errors {

		io.WriteString(out, "\t"+msg+"\n")
	}
}
