//go:build ignore
// +build ignore

// 本文件不参与项目整体编译, 因此使用 go:build ignore 忽略
// 因此本文件可以使用 main 作为包名直接运行

// 在本目录下执行本文件:
//
//	go run partial_gen.go
//
// 以生成 partial_more.go 文件
package main

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"strings"

	"code.gopub.tech/commons/funcs"
	"code.gopub.tech/commons/iters"
)

func main() {
	buf := new(bytes.Buffer)

	fmt.Fprintln(buf, "// code generated by `go run partial_gen.go`; DO NOT EDIT.")
	fmt.Fprintln(buf)
	fmt.Fprintln(buf, "package funcs")
	fmt.Fprintln(buf)

	for i := 2; i <= 10; i++ {
		fmt.Fprintln(buf, makeType(i))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, makeOf(i))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, makeMethod(i))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, makePartialFun(i))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, makePartialnFunc(i))
		fmt.Fprintln(buf)
		fmt.Fprintln(buf, makePartialRFun(i))
		fmt.Fprintln(buf)
	}

	// fmt.Println(buf.String())

	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("format err=%+v\n", err)
	} else {
		err = os.WriteFile("partial_more.go", source, 0660)
		if err != nil {
			fmt.Printf("write err=%+v\n", err)
		}
	}
}

func makeType(n int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`type Func%d[%s, R any] func(%s) R`, n, p, p)
}

func makeParams(n, skip int, fn funcs.Function[int, string]) string {
	return strings.Join(iters.Maps(
		iters.Range(1, n+1).
			Filter(func(i int) bool { return i != skip }),
		fn,
	).ToSlice(), ", ")
}

func printType(i int) string    { return fmt.Sprintf("T%d", i) }
func printArg(i int) string     { return fmt.Sprintf("t%d", i) }
func printArgType(i int) string { return fmt.Sprintf("t%d T%d", i, i) }

func makeOf(n int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func Of%d[
	%s, R any,
	F ~func(%s) R,
](fn F) Func%d[%s, R] {
	return Func%d[%s, R](fn)
}`, n, p, p, n, p, n, p)
}

func makeMethod(n int) string {
	buf := new(bytes.Buffer)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(buf, makeMethodN(n, i))
		fmt.Fprintln(buf)
	}
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func (fn Func%d[%s, R]) Call(%s) R {
	return fn(%s)
}

func (fn Func%d[%s, R]) Partial(t1 T1) Func%d[%s, R] {
	return Partial%d(fn, t1)
}

%s

func (fn Func%d[%s, R]) PartialR(t%d T%d) Func%d[%s, R] {
	return Partial%dR(fn, t%d)
}`,
		n, p, makeParams(n, 0, printArgType), makeParams(n, 0, printArg),
		n, p, n-1, makeParams(n, 1, printType), n,
		buf.String(),
		n, p, n, n, n-1, makeParams(n-1, 0, printType), n, n)
}

func makeMethodN(n, i int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func  (fn Func%d[%s, R]) Partial%d(t%d T%d) Func%d[%s, R] {
	return Partial%dN%d(fn, t%d)
}`,
		n, p, i, i, i, n-1, makeParams(n, i, printType),
		n, i, i)
}

func makePartialFun(n int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func Partial%d[
	%s, R any, 
	F ~func(%s) R,
](fn F, t1 T1) Func%d[%s, R] {
	return func(%s) R {
		return fn(%s)
	}
}`,
		n, p, p, n-1, makeParams(n, 1, printType),
		makeParams(n, 1, printArgType), makeParams(n, 0, printArg),
	)
}

func makePartialRFun(n int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func Partial%dR[
	%s, R any,
	F ~func(%s) R,
](fn F, t%d T%d) Func%d[%s, R] {
	return func(%s) R {
		return fn(%s)
	}
}`,
		n, p, p, n, n, n-1, makeParams(n-1, 0, printType),
		makeParams(n-1, 0, printArgType), makeParams(n, 0, printArg))
}

func makePartialnFunc(n int) string {
	buf := new(bytes.Buffer)
	for i := 1; i <= n; i++ {
		fmt.Fprintln(buf, makePartialN(n, i))
		fmt.Fprintln(buf)
	}
	return buf.String()
}

func makePartialN(n, i int) string {
	p := makeParams(n, 0, printType)
	return fmt.Sprintf(`func Partial%dN%d[
	%s, R any,
	F ~func(%s) R,
](fn F, t%d T%d) Func%d[%s, R] {
	return func(%s) R {
		return fn(%s)
	}
}`,
		n, i, p, p, i, i, n-1, makeParams(n, i, printType),
		makeParams(n, i, printArgType), makeParams(n, 0, printArg))
}
