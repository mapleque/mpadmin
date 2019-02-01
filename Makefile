.PHONY:all run build
all: build

run:
	cd main && go run main.go

build:
	-mkdir bin
	GOOS=linux GOARCH=amd64 go build -o bin/linux main/main.go
	GOOS=darwin GOARCH=amd64 go build -o bin/mac main/main.go
	GOOS=windows GOARCH=amd64 go build -o bin/win64.exe main/main.go
	GOOS=windows GOARCH=386 go build -o bin/win32.exe main/main.go

