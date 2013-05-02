package main

import (
	"go/parser"
	"go/token"
	"log"
	"os"
	"strings"
)

func main() {
	fileSet := token.NewFileSet()
	comFileSet := token.NewFileSet() // The FileSet used to parse the COM declarations.
	mod := newModule(comFileSet)

	// Iterate over all the files specified on the command line.
	for _, filename := range os.Args[1:] {
		f, err := parser.ParseFile(fileSet, filename, nil, parser.ParseComments)
		if err != nil {
			log.Fatalln("error parsing the source file", filename, ":", err)
		}

		// Iterate over all the comments in the current file, picking out the ones that
		// start with "com".
		// Join them into a long string to be parsed, starting with a package declaration.
		chunks := []string{"package " + f.Name.Name}
		for _, cg := range f.Comments {
			for _, c := range cg.List {
				text := strings.TrimSpace(strings.Trim(c.Text, "/*"))
				if strings.HasPrefix(text, "com") {
					text = text[len("com"):]
					if text == "" || !strings.Contains(" \t\n", text[:1]) {
						continue
					}
					text = strings.TrimSpace(text)
					chunks = append(chunks, text)
				}
			}
		}
		comDecls := strings.Join(chunks, "\n")

		// Now parse the concatenated result as a Go source file.
		comAST, err := parser.ParseFile(comFileSet, filename, comDecls, parser.ParseComments)
		if err != nil {
			log.Fatalln("error parsing the COM declarations from", filename, ":", err)
		}

		err = mod.loadFile(comAST)
		if err != nil {
			log.Fatalln("error loading declarations from", filename, "into module:", err)
		}
	}

	err := mod.write(os.Stdout)
	if err != nil {
		log.Fatalln("error generating output:", err)
	}
}
