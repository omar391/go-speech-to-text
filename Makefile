buildwin:
	CGO_ENABLED=1 go build -ldflags="-s -w" -o ./app.exe main.go

buildjs:
	CGO_ENABLED=0 GOARCH=wasm GOOS=js go build -ldflags="-s -w" -o ./nodejs/lib.wasm ./test/main.go
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./nodejs/

tinyjs:
	tinygo build -o ./nodejs/lib.wasm -target wasm ./test/main.go
	cp "$$(go env GOROOT)/misc/wasm/wasm_exec.js" ./nodejs/


runjs:
	node ./nodejs/index.js

alljs: buildjs runjs