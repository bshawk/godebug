package main

import "github.com/mailgun/godebug/lib"

var regression_in_go_scope = godebug.EnteringNewScope(regression_in_go_contents)

func main() {
	ctx, _ok := godebug.EnterFunc(main)
	if !_ok {
		return
	}
	godebug.Line(ctx, regression_in_go_scope, 7)

	foo := func(i int) int {
		var result1 int
		fn := func(ctx *godebug.Context) {
			result1 = func() int {
				scope := regression_in_go_scope.EnteringNewChildScope()
				scope.Declare("i", &i)
				godebug.Line(ctx, scope, 8)
				return i
			}()
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
		return result1
	}(3)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("foo", &foo)
	godebug.Line(ctx, scope, 10)

	_ = foo
	{
		scope := scope.EnteringNewChildScope()

		for _, s := range []string{"foo"} {
			godebug.Line(ctx, scope, 14)
			scope.Declare("s", &s)
			godebug.Line(ctx, scope, 15)
			_ = s
		}
		godebug.Line(ctx, scope, 14)
	}
	godebug.Line(ctx, scope, 19)

	c := make(chan bool)
	scope.Declare("c", &c)
	godebug.Line(ctx, scope, 20)
	go func() {
		fn := func(ctx *godebug.Context) {
			godebug.Line(ctx, scope, 21)
			c <- true
		}
		if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
			defer godebug.ExitFunc(ctx)
			fn(ctx)
		}
	}()
	godebug.Line(ctx, scope, 23)
	<-c
	godebug.Line(ctx, scope, 26)

	defer println("Hello")
	defer godebug.Defer(ctx, scope, 26)
	godebug.Line(ctx, scope, 29)

	if false {
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 30)
		s := "hello"
		godebug.ElseIfExpr(ctx, scope, 30)
		if s == "hello" {
			godebug.Line(ctx, scope, 31)
			println(s)
		}
	}
	godebug.Line(ctx, scope, 35)

	m := map[string]int{"test": 5}
	scope.Declare("m", &m)
	godebug.Line(ctx, scope, 36)
	if false {
	} else {
		godebug.ElseIfSimpleStmt(ctx, scope, 37)
		_, ok := m["test"]
		godebug.ElseIfExpr(ctx, scope, 37)
		if ok {
			godebug.Line(ctx, scope, 38)
			println("test")
		}
	}
	godebug.SetTraceGen(ctx)
	godebug.Line(ctx, scope, 43)

	const n = 10
	scope.Constant("n", n)
	godebug.Line(ctx, scope, 44)
	_ = n
	godebug.Line(ctx, scope, 46)

	name1(5)
	godebug.Line(ctx, scope, 47)
	name2()
	godebug.Line(ctx, scope, 48)
	T{}.name3()
}

func _switch() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = _switch()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 53)

	switch {
	case godebug.Case(ctx, regression_in_go_scope, 54):
		fallthrough
	case false:
		godebug.Line(ctx, regression_in_go_scope, 55)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 56)
		godebug.Line(ctx, regression_in_go_scope, 57)
		return 5
	}
}

func _select() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = _select()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Select(ctx, regression_in_go_scope, 63)

	select {
	case <-godebug.Comm(ctx, regression_in_go_scope, 64):
		panic("impossible")
	case <-make(chan bool):
		godebug.Line(ctx, regression_in_go_scope, 64)
		godebug.Line(ctx, regression_in_go_scope, 65)
		return 4
	default:
		godebug.Line(ctx, regression_in_go_scope, 66)
		godebug.Line(ctx, regression_in_go_scope, 67)
		return 5
	case <-godebug.EndSelect(ctx, regression_in_go_scope):
		panic("impossible")
	}
}

func name1(_name1 int) {
	ctx, _ok := godebug.EnterFunc(func() {
		name1(_name1)
	})
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name1", &_name1)
	godebug.Line(ctx, scope, 73)
	if true {
		godebug.Line(ctx, scope, 74)
		_ = _name1
	}
}

func name2() (_name2 string) {
	ctx, _ok := godebug.EnterFunc(func() {
		_name2 = name2()
	})
	if !_ok {
		return _name2
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name2", &_name2)
	godebug.Line(ctx, scope, 80)
	if true {
		godebug.Line(ctx, scope, 81)
		_name2 = "foo"
	}
	godebug.Line(ctx, scope, 83)
	return _name2
}

type T struct{}

func (_name3 T) name3() {
	ctx, _ok := godebug.EnterFunc(_name3.name3)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("name3", &_name3)
	godebug.Line(ctx, scope, 90)
	if true {
		godebug.Line(ctx, scope, 91)
		_ = _name3
	}
}

var nestedSwitch = func() {
	fn := func(ctx *godebug.Context) {
		godebug.Line(ctx, regression_in_go_scope, 96)
		var foo interface {
		} = 5
		scope := regression_in_go_scope.EnteringNewChildScope()
		scope.Declare("foo", &foo)
		godebug.Line(ctx, scope, 98)
		switch {
		default:
			godebug.Line(ctx, scope, 99)
			godebug.Line(ctx, scope, 100)
			switch foo.(type) {
			case int:
				godebug.Line(ctx, scope, 101)
			}
		}
	}
	if ctx, _ok := godebug.EnterFuncLit(fn); _ok {
		defer godebug.ExitFunc(ctx)
		fn(ctx)
	}
}

func init() {
	doFallthrough()
}

func doFallthrough() {
	ctx, _ok := godebug.EnterFunc(doFallthrough)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 112)
	fellthrough := false
	scope := regression_in_go_scope.EnteringNewChildScope()
	scope.Declare("fellthrough", &fellthrough)
	godebug.Line(ctx, scope, 113)
	switch {
	case godebug.Case(ctx, scope, 114):
		fallthrough
	case true:
		godebug.Line(ctx, scope, 115)
		fallthrough
	case godebug.Case(ctx, scope, 116):
		fallthrough
	case false:
		godebug.Line(ctx, scope, 117)
		fellthrough = true
	}
	godebug.Line(ctx, scope, 119)
	if !fellthrough {
		godebug.Line(ctx, scope, 120)
		panic("fallthrough statement did not work")
	}
}

func a() int {
	var result1 int
	ctx, _ok := godebug.EnterFunc(func() {
		result1 = a()
	})
	if !_ok {
		return result1
	}
	defer godebug.ExitFunc(ctx)
	godebug.Line(ctx, regression_in_go_scope, 125)
	return 0
}

func init() {
	switchInit()
}

func switchInit() {
	ctx, _ok := godebug.EnterFunc(switchInit)
	if !_ok {
		return
	}
	defer godebug.ExitFunc(ctx)
	godebug.SetTraceGen(ctx)
	{
		godebug.Line(ctx, regression_in_go_scope, 135)
		a := a()
		scope := regression_in_go_scope.EnteringNewChildScope()
		scope.Declare("a", &a)
		switch {
		default:
			godebug.Line(ctx, scope, 136)
			godebug.Line(ctx, scope, 137)
			_ = a
		}
	}
	godebug.Line(ctx, regression_in_go_scope, 139)
	_ = "the variable a should be out of scope"
}

var regression_in_go_contents = `package main

import "github.com/mailgun/godebug/lib"

func main() {
	// Nested scope in the first declaration in a function.
	foo := func(i int) int {
		return i
	}(3)
	_ = foo

	// String literal in range statement.
	// Blank identifier in range statement.
	for _, s := range []string{"foo"} {
		_ = s
	}

	// go statement with function literal.
	c := make(chan bool)
	go func() {
		c <- true
	}()
	<-c

	// String literal in defer statement.
	defer println("Hello")

	// String literal in else-if statement.
	if false {
	} else if s := "hello"; s == "hello" {
		println(s)
	}

	// Comma-ok in else-if
	m := map[string]int{"test": 5}
	if false {
	} else if _, ok := m["test"]; ok {
		println("test")
	}

	// Constant declaration.
	_ = "breakpoint"
	const n = 10
	_ = n

	name1(5)
	name2()
	T{}.name3()
}

func _switch() int {
	// Terminating switch statement in function with return value.
	switch {
	case false:
		return 4
	default:
		return 5
	}
}

func _select() int {
	// Terminating select statement in function with return value.
	select {
	case <-make(chan bool):
		return 4
	default:
		return 5
	}
}

// Function shares a name with an input parameter.
func name1(name1 int) {
	if true {
		_ = name1
	}
}

// Function shares a name with an output parameter.
func name2() (name2 string) {
	if true {
		name2 = "foo"
	}
	return name2
}

type T struct{}

// Function shares a name with its receiver
func (name3 T) name3() {
	if true {
		_ = name3
	}
}

var nestedSwitch = func() {
	var foo interface{} = 5
	// Type switch nested inside expression switch
	switch {
	default:
		switch foo.(type) {
		case int:
		}
	}
}

func init() {
	doFallthrough()
}

// Fallthrough should work.
func doFallthrough() {
	fellthrough := false
	switch {
	case true:
		fallthrough
	case false:
		fellthrough = true
	}
	if !fellthrough {
		panic("fallthrough statement did not work")
	}
}

func a() int {
	return 0
}

func init() {
	switchInit()
}

// Don't repeat switch initialization, use correct scope inside switch.
func switchInit() {
	godebug.SetTrace()
	switch a := a(); {
	default:
		_ = a
	}
	_ = "the variable a should be out of scope"
}
`
