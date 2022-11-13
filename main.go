package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"syscall/js"
)

var joke string

func add() js.Func {
	myFunc := js.FuncOf(func(this js.Value, i []js.Value) any {

		int1, _ := strconv.Atoi(i[0].String())
		int2, _ := strconv.Atoi(i[1].String())
		println("Sum is", int1+int2)
		return int1 + int2
	})
	return myFunc
}

func factorial_wasm() js.Func {
	myFunc := js.FuncOf(func(this js.Value, i []js.Value) any {
		return factorial(strconv.Atoi(i[0].String()))
	})
	return myFunc
}

func factorial(a int) int {
	if a <= 1 {
		return a
	}
	return a * factorial_wasm(a-1)
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

			// fmt.Println("request is", string(b))
			js.Global().Get("document").Call("getElementById", i[0].String()).Set("innerHTML", string(b))
			joke = string(b)
		}()
		return joke
	})
	return myFunc
}

func registerCallbacks() {
	js.Global().Set("add", add())
	js.Global().Set("factorial_wasm", factorial_wasm())
	js.Global().Set("request", request())
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
