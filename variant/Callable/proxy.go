package Callable

import (
	"graphics.gd/variant"
	"graphics.gd/variant/Array"
)

// Proxy can be implemented to provide a foreign-managed array representation. This can be useful
// when you want to access a callable with its implementation hosted in another programming language.
type Proxy interface {
	Name(complex128) string
	Args(complex128) (args int, bind Array.Any)
	Call(complex128, ...variant.Any) variant.Any
	Bind(complex128, ...variant.Any) (Proxy, complex128)
}

func Through(proxy Proxy, state complex128) Function {
	return Function{proxy: proxy, state: state}
}
