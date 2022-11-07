package main

import (
	"strconv"
	"syscall/js"
)

func add() js.Func {
	myFunc := js.FuncOf(func(this js.Value, i []js.Value) any {
		value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
		value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

		int1, _ := strconv.Atoi(value1)
		int2, _ := strconv.Atoi(value2)

		js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1+int2)
		println("Sum is", int1+int2)
		return nil
	})
	return myFunc
}

func subtract() js.Func {
	myFunc := js.FuncOf(func(this js.Value, i []js.Value) any {
		value1 := js.Global().Get("document").Call("getElementById", i[0].String()).Get("value").String()
		value2 := js.Global().Get("document").Call("getElementById", i[1].String()).Get("value").String()

		int1, _ := strconv.Atoi(value1)
		int2, _ := strconv.Atoi(value2)

		js.Global().Get("document").Call("getElementById", i[2].String()).Set("value", int1-int2)
		println("Subtract is", int1-int2)
		return nil
	})
	return myFunc
}

func registerCallbacks() {
	js.Global().Set("add", add())
	js.Global().Set("subtract", subtract())
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
