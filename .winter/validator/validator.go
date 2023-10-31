package validator

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"log"
	"os"
	pathlib "path"
	"path/filepath"
)

func StructureValidation() {
	getWd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}

	internalDir := pathlib.Join(getWd, "internal")
	serviceDir := pathlib.Join(internalDir, "controller", "v1")

	err = filepath.Walk(serviceDir, func(path string, info fs.FileInfo, err error) error {
		if info.Name() == "users.go" {
			fSet := token.NewFileSet()
			file, err := parser.ParseFile(fSet, pathlib.Join(serviceDir, info.Name()), nil, parser.ParseComments)
			if err != nil {
				log.Fatalln(err)
			}

			type Arg struct {
				Name string
				Type string
			}

			type Layer struct {
				Name string
				Args []Arg
			}

			var layer Layer

			ast.Inspect(file, func(x ast.Node) bool {
				switch x.(type) {
				case *ast.TypeSpec:
					s, _ := x.(*ast.TypeSpec)
					layer.Name = s.Name.String()
				//case *ast.FuncDecl:
				//	s, _ := x.(*ast.FuncDecl)
				//	fmt.Println(s)
				case *ast.StructType:
					s, _ := x.(*ast.StructType)
					for _, field := range s.Fields.List {
						t := field.Type.(*ast.SelectorExpr)
						layer.Args = append(layer.Args, Arg{
							Name: field.Names[0].String(),
							Type: t.X.(*ast.Ident).Name + "." + t.Sel.String(),
						})
					}
				}
				return true
			})

			var (
				fArgs []byte
				sArgs []byte
			)

			for _, arg := range layer.Args {
				fArgs = fmt.Appendf(fArgs, ", %s %s", arg.Name, arg.Type)
				sArgs = fmt.Appendf(sArgs, "\n        %s: %s,", arg.Name, arg.Name)
			}

			a := fmt.Sprintf("func New%s(%s) *%s {\n    return &%s{%s\n    }\n}",
				layer.Name, string(fArgs[2:]), layer.Name, layer.Name, string(sArgs),
			)

			fmt.Println(a)
		}

		return err
	})
	if err != nil {
		log.Fatalln(err)
	}
}
