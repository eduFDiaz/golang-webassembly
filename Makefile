all: clean package
clean:
	rm -rf bin
	rm -f ./frontend/src/assets/wasm/lib.wasm
# dep:
# 	cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./frontend/src/assets/wasm/
package:
	GOOS=js GOARCH=wasm go build -o ./frontend/src/assets/wasm/lib.wasm main.go
# run:
# 	go run server.go