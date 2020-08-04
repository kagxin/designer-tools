all:
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mkdir.exe mkdir/main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o mktxt.exe mktxt/main.go
clean:
	rm mkdir.exe mktxt.exe