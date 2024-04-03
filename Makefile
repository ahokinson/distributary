build:
	GOARCH=amd64 GOOS=linux go build -o build/distributary-linux main.go
	GOARCH=amd64 GOOS=darwin go build -o build/distributary-darwin main.go
	GOARCH=amd64 GOOS=windows go build -o build/distributary-windows.exe main.go

clean:
	go clean
	rm build/distributary-linux
	rm build/distributary-darwin
	rm build/distributary-windows.exe
