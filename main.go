package main

import (
	"fmt"
	"log"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
)

func main() {
	d := cel.Declarations(decls.NewVar("name", decls.String))
	//d := cel.Declarations()

	env, _ := cel.NewEnv(d)

	ast, iss := env.Compile(`"fds reds trfdsz $name "+  name`)
	//output fds reds trfdsz $name CEL
	//ast, iss := env.Compile(`"Hello world! I'm {{name}} ."`)

	// Check iss for compilation errors.
	if iss.Err() != nil {
		log.Fatalln(iss.Err())
	}
	prg, _ := env.Program(ast)
	out, _, _ := prg.Eval(map[string]interface{}{
		"name": "CEL",
	})
	fmt.Println(out)

}
