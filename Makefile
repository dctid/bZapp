.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./build/slash
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./build/slash ./slash

run:	clean build
	sam local start-api