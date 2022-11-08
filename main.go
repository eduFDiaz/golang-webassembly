package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
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

func request() js.Func {
	myFunc := js.FuncOf(func(this js.Value, i []js.Value) any {

		go func() {
			req, err := http.NewRequest("GET", "https://icanhazdadjoke.com", nil)
			if err != nil {
				log.Fatalln(err)
			}

			req.Header.Set("Accept", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Fatalln(err)
			}

			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			// b, err := ioutil.ReadAll(resp.Body)  Go.1.15 and earlier
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("request is", string(b))
		}()
		return nil
	})
	return myFunc
}

func registerCallbacks() {
	js.Global().Set("add", add())
	js.Global().Set("subtract", subtract())
	js.Global().Set("request", request())
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
