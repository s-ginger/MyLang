package evaluator

import (
	"fmt"
	"mylang/ast" // твой AST
	"log"
)

type Value interface{}

type Env struct {
	Vars    map[string]Value
	Funcs   map[string]*parser.FunctionStmt
	Parent  *Env // для вложенных блоков/функций
}


func NewEnv(parent *Env) *Env {
	return &Env{
		Vars:   make(map[string]Value),
		Funcs:  make(map[string]*parser.FunctionStmt),
		Parent: parent,
	}
}


func EvalExpr(e parser.Expr, env *Env) (Value, error) {
	switch expr := e.(type) {

	case parser.Literal:
		return expr.Value, nil

	case parser.Variable:
		val, ok := env.LookupVar(expr.Name)
		if !ok {
			return nil, fmt.Errorf("undefined variable: %s", expr.Name)
		}
		log.Println(val)
		return val, nil

	case parser.Operation:
		left, err := EvalExpr(expr.Left, env)
		if err != nil {
			return nil, err
		}
		right, err := EvalExpr(expr.Right, env)
		if err != nil {
			return nil, err
		}
		log.Println(left, expr.Op, right)
		return applyBinaryOp(expr.Op, left, right)

	case parser.UnaryExpr:
		val, err := EvalExpr(expr.Expr, env)
		if err != nil {
			return nil, err
		}
		log.Println(expr.Op, val)
		return applyUnaryOp(expr.Op, val)

	case parser.CallExpr:
		return callFunction(expr.Name, expr.Args, env)

	case parser.ListExpr:
		var result []Value
		for _, el := range expr.Elements {
			v, err := EvalExpr(el, env)
			if err != nil {
				return nil, err
			}
			result = append(result, v)
		}
		log.Println(result)
		return result, nil

	case parser.IndexExpr:
		coll, err := EvalExpr(expr.Collection, env)
		if err != nil {
			return nil, err
		}
		index, err := EvalExpr(expr.Index, env)
		if err != nil {
			return nil, err
		}
		log.Println(coll, index)
		return getFromIndex(coll, index)

	default:
		return nil, fmt.Errorf("unknown expr: %T", expr)
	}
}

func ExecStmt(stmt parser.Stmt, env *Env) (Value, error) {
	switch s := stmt.(type) {

	case *parser.Block:
		var result Value
		for _, stmt := range s.Statements {
			r, err := ExecStmt(stmt, env)
			if err != nil {
				return nil, err
			}
			if _, ok := stmt.(*parser.ReturnStmt); ok {
				return r, nil
			}
			result = r
		}
		log.Println(result)
		return result, nil

	case *parser.Assign:
		val, err := EvalExpr(s.Value, env)
		if err != nil {
			return nil, err
		}
		env.Vars[s.Name] = val
		return nil, nil

	case *parser.Print:
		val, err := EvalExpr(s.Expr, env)
		if err != nil {
			return nil, err
		}
		fmt.Println(val)
		return nil, nil

	case *parser.FunctionStmt:
		env.Funcs[s.Name] = s
		return nil, nil

	case *parser.CallStmt:
		_, err := callFunction(s.Name, s.Args, env)
		return nil, err

	case *parser.ReturnStmt:
		return EvalExpr(s.Value, env)

	default:
		return nil, fmt.Errorf("unknown statement: %T", s)
	}
}

func callFunction(name string, args []parser.Expr, env *Env) (Value, error) {
	fn, ok := env.Funcs[name]
	if !ok {
		return nil, fmt.Errorf("undefined function: %s", name)
	}
	if len(args) != len(fn.Args) {
		return nil, fmt.Errorf("wrong number of args for function %s", name)
	}

	localEnv := NewEnv(env)
	for i, paramName := range fn.Args {
		val, err := EvalExpr(args[i], env)
		if err != nil {
			return nil, err
		}
		localEnv.Vars[paramName] = val
	}

	return ExecStmt(&fn.Body, localEnv)
}

func (env *Env) LookupVar(name string) (Value, bool) {
	val, ok := env.Vars[name]
	if ok {
		return val, true
	}
	if env.Parent != nil {
		return env.Parent.LookupVar(name)
	}
	return nil, false
}

func applyBinaryOp(op string, left, right interface{}) (interface{}, error) {
	switch op {
	case "+":
		switch l := left.(type) {
		case int:
			if r, ok := right.(int); ok {
				return l + r, nil
			}
		case float64:
			if r, ok := right.(float64); ok {
				return l + r, nil
			}
		case string:
			if r, ok := right.(string); ok {
				return l + r, nil
			}
		}
	case "-":
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				return l - r, nil
			}
		}
	case "*":
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				return l * r, nil
			}
		}
	case "/":
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				if r == 0 {
					return nil, fmt.Errorf("division by zero")
				}
				return l / r, nil
			}
		}
	case "==":
		return left == right, nil
	case "!=":
		return left != right, nil
	case "<":
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				return l < r, nil
			}
		}
	case ">":
		if l, ok := left.(int); ok {
			if r, ok := right.(int); ok {
				return l > r, nil
			}
		}
	}
	return nil, fmt.Errorf("unsupported binary operation: %v %s %v", left, op, right)
}

func applyUnaryOp(op string, val interface{}) (interface{}, error) {
	switch op {
	case "-":
		if v, ok := val.(int); ok {
			return -v, nil
		}
	case "!":
		if v, ok := val.(bool); ok {
			return !v, nil
		}
	}
	return nil, fmt.Errorf("unsupported unary operation: %s %v", op, val)
}

func getFromIndex(collection, index interface{}) (interface{}, error) {
	switch coll := collection.(type) {
	case []interface{}:
		if i, ok := index.(int); ok {
			if i >= 0 && i < len(coll) {
				return coll[i], nil
			}
			return nil, fmt.Errorf("index out of range")
		}
	}
	return nil, fmt.Errorf("unsupported indexing operation")
}



