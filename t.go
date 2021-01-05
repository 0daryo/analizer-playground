package main

func F() {}

type T struct {
	F func()
}

func (T) M() {}

var Fv = F

func Comparison() {
	var t T
	var fn func()
	if fn == nil || Fv == nil || t.F == nil {
		// no error; these func vars or fields may be nil
	}
	if F == nil { // want "comparison of function F == nil is always false"
		panic("can't happen")
	}
	if t.M == nil { // want "comparison of function M == nil is always false"
		panic("can't happen")
	}
	if F != nil { // want "comparison of function F != nil is always true"
		if t.M != nil { // want "comparison of function M != nil is always true"
			return
		}
	}
	panic("can't happen")
}
