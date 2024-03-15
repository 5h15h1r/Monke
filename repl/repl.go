package repl

import (
	"bufio"
	"fmt"
	"io"
	"monke/lexer"
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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")

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
