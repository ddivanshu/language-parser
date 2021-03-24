package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/checker/decls"
	"github.com/google/cel-go/common/types"
	"github.com/google/cel-go/common/types/ref"
	"github.com/google/cel-go/interpreter/functions"
	exprpb "google.golang.org/genproto/googleapis/api/expr/v1alpha1"
)

func equals(args ...ref.Val) ref.Val {
	fmt.Print(args)
	if args[0] == args[1] {
		return types.True
	}
	return types.False
}
func main() {
	celDeclarations := cel.Declarations(
		decls.NewVar("repo", decls.NewMapType(decls.String, decls.Any)),
		decls.NewVar("context", decls.NewMapType(decls.String, decls.Any)),
		decls.NewFunction("beginsWith", decls.NewOverload(
			"beginsWith",
			[]*exprpb.Type{decls.String, decls.String, decls.String, decls.String, decls.String},
			decls.String)),
		decls.NewFunction("toLower", decls.NewOverload(
			"toLower",
			[]*exprpb.Type{decls.String},
			decls.String)),
	)

	celFunctions := cel.Functions(
		&functions.Overload{
			Operator: "toLower",
			Unary:    toLower,
		}, &functions.Overload{
			Operator: "beginsWith",
			Function: beginsWith,
		})

	context := `{"name":"cel","context":{"company":"github","Domain":"IT"} ,"repo":{"nodeVersion":1222 , "package":"a/c/c","docker":{"path":"dockerFilepath","image":"nodeImage"}}}`
	contextMap := map[string]interface{}{}
	json.Unmarshal([]byte(context), &contextMap)
	env, _ := cel.NewEnv(celDeclarations)
	file, err := os.Open("template.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		//rendering logic
		compiledTemplate, issue := env.Compile(scanner.Text())
		if issue.Err() != nil {
			log.Fatalln(issue.Err())
		}
		prg, err := env.Program(compiledTemplate, celFunctions)
		if err != nil {
			log.Fatal(err.Error())
		}
		out, _, _ := prg.Eval(contextMap)
		fmt.Println(out)
	}
}

func beginsWith(args ...ref.Val) ref.Val { return args[1] }
func toLower(arg ref.Val) ref.Val {
	return types.String("lowerCase").Add(arg)

}
