package main

import (
	"go/token"
	"go/parser"
	"go/ast"
)

func main() {
	ExamplePrint()
}

// This example shows what an AST looks like when printed for debugging.
func ExamplePrint() {
	// src is the input for which we want to print the AST.
	src := `
package main
func main() {
	println("Hello, World!")
}
`

	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	// Print the AST.
	ast.Print(fset, f)

	// Output:
	//      0  *ast.File {
	//      1  .  Package: 2:1
	//      2  .  Name: *ast.Ident {
	//      3  .  .  NamePos: 2:9
	//      4  .  .  Name: "main"
	//      5  .  }
	//      6  .  Decls: []ast.Decl (len = 1) {
	//      7  .  .  0: *ast.FuncDecl {
	//      8  .  .  .  Name: *ast.Ident {
	//      9  .  .  .  .  NamePos: 3:6
	//     10  .  .  .  .  Name: "main"
	//     11  .  .  .  .  Obj: *ast.Object {
	//     12  .  .  .  .  .  Kind: func
	//     13  .  .  .  .  .  Name: "main"
	//     14  .  .  .  .  .  Decl: *(obj @ 7)
	//     15  .  .  .  .  }
	//     16  .  .  .  }
	//     17  .  .  .  Type: *ast.FuncType {
	//     18  .  .  .  .  Func: 3:1
	//     19  .  .  .  .  Params: *ast.FieldList {
	//     20  .  .  .  .  .  Opening: 3:10
	//     21  .  .  .  .  .  Closing: 3:11
	//     22  .  .  .  .  }
	//     23  .  .  .  }
	//     24  .  .  .  Body: *ast.BlockStmt {
	//     25  .  .  .  .  Lbrace: 3:13
	//     26  .  .  .  .  List: []ast.Stmt (len = 1) {
	//     27  .  .  .  .  .  0: *ast.ExprStmt {
	//     28  .  .  .  .  .  .  X: *ast.CallExpr {
	//     29  .  .  .  .  .  .  .  Fun: *ast.Ident {
	//     30  .  .  .  .  .  .  .  .  NamePos: 4:2
	//     31  .  .  .  .  .  .  .  .  Name: "println"
	//     32  .  .  .  .  .  .  .  }
	//     33  .  .  .  .  .  .  .  Lparen: 4:9
	//     34  .  .  .  .  .  .  .  Args: []ast.Expr (len = 1) {
	//     35  .  .  .  .  .  .  .  .  0: *ast.BasicLit {
	//     36  .  .  .  .  .  .  .  .  .  ValuePos: 4:10
	//     37  .  .  .  .  .  .  .  .  .  Kind: STRING
	//     38  .  .  .  .  .  .  .  .  .  Value: "\"Hello, World!\""
	//     39  .  .  .  .  .  .  .  .  }
	//     40  .  .  .  .  .  .  .  }
	//     41  .  .  .  .  .  .  .  Ellipsis: -
	//     42  .  .  .  .  .  .  .  Rparen: 4:25
	//     43  .  .  .  .  .  .  }
	//     44  .  .  .  .  .  }
	//     45  .  .  .  .  }
	//     46  .  .  .  .  Rbrace: 5:1
	//     47  .  .  .  }
	//     48  .  .  }
	//     49  .  }
	//     50  .  Scope: *ast.Scope {
	//     51  .  .  Objects: map[string]*ast.Object (len = 1) {
	//     52  .  .  .  "main": *(obj @ 11)
	//     53  .  .  }
	//     54  .  }
	//     55  .  Unresolved: []*ast.Ident (len = 1) {
	//     56  .  .  0: *(obj @ 29)
	//     57  .  }
	//     58  }
}
