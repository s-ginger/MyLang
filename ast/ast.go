package parser



// --- AST выражения ---
type Expr interface{}

type Literal struct {
	Kind  string      // "int", "float", "string", "bool"
	Value interface{}
}

type Variable struct {
	Name string
}

type Operation struct {
	Op   string
	Left, Right Expr
}

type CallExpr struct {
	Name string
	Args []Expr
}

type UnaryExpr struct {
	Op   string
	Expr Expr
}

type ListExpr struct {
	Elements []Expr
}

type IndexExpr struct {
	Collection Expr
	Index      Expr
}

// --- AST команды ---
type Stmt interface{}

type Block struct {
	Statements []Stmt
}

type Assign struct {
	Name  string
	Value Expr
}

type Print struct {
	Expr Expr
}

type IfStmt struct {
	Branches []IfBranch
	Else     Stmt // может быть Block, Function, и т.д.
}

type IfBranch struct {
	Cond Expr
	Then Block
}

type ForStmt struct {
	Init Stmt     // например: Assign
	Cond Expr
	Post Stmt     // например: Assign
	Body Block
}

type ReturnStmt struct {
	Value Expr
}

type FunctionStmt struct {
	Name string
	Args []string
	Body Block
}

type CallStmt struct {
	Name string
	Args []Expr
}

type BreakStmt struct{}
