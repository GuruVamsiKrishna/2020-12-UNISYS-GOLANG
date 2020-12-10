dev:
	go run src/main.go

build:
	go build -o server src/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -o server-linux src/main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o server-win64.exe src/main.go