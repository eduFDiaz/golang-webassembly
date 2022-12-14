# golang-webassembly
Working code for the TutorialEdge intro to webassembly as of today 11/06/2022
[![Go Webassembly tutorial](https://img.youtube.com/vi/4kBvvk2Bzis/0.jpg)](https://www.youtube.com/watch?v=4kBvvk2Bzis&t=32s)

This was run under windows 11, for unix use replace 

`cp /c/Program\ Files/Go/misc/wasm/wasm_exec.js .` 

in the Makefile with 

`cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .`

## then run to build and serve the project
```make
make all
```
Then open your favorite browser at http://localhost:8080/
![alt text](Screenshot-2022-11-06-231900.png "Title")
## Version
* go 1.19 
* syscall/js 

The add and subtract methods had to be slightly changed due to breaking changes in `syscall/js`
```golang
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
```

Big shoutout to https://github.com/sponsors/elliotforbes for his amazing work at TutorialEdge ????????????
