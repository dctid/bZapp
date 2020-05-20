.PHONY: deps clean build

deps:
	go get -u ./...

clean: 
	rm -rf ./bin/slash
	
build:
	GOOS=linux GOARCH=amd64 go build -o ./bin/slash ./slash
	GOOS=linux GOARCH=amd64 go build -o ./bin/interactive ./interactive

run:	clean build
	sam local start-api