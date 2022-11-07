all: clean dep package run
clean:
	rm -rf bin
	rm -f lib.wasm
dep:
	cp /c/Program\ Files/Go/misc/wasm/wasm_exec.js .
package:
	GOOS=js GOARCH=wasm go build -o lib.wasm main.go
run:
	go run server.go