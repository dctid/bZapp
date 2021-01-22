export AWS_DEFAULT_REGION ?= us-east-1
export APP ?= bzapp

.PHONY: test

app: clean dynamo-start dev

clean:
	rm -f $(wildcard handlers/*/main)

deploy: BUCKET = pkgs-$(shell aws sts get-caller-identity --output text --query 'Account')-$(AWS_DEFAULT_REGION)
deploy: clean handlers
	@aws s3api head-bucket --bucket $(BUCKET) || aws s3 mb s3://$(BUCKET) --region $(AWS_DEFAULT_REGION)
	sam package --output-template-file out.yml --s3-bucket $(BUCKET) --template-file template.yml
	sam deploy --capabilities CAPABILITY_NAMED_IAM  --template-file out.yml --stack-name $(APP) --parameter-overrides SlackSigningSecret=$(SLACK_SIGNING_SECRET) SlackToken=$(SLACK_TOKEN)


dev-debug:
	make clean
	GCFLAGS="-N -l" make -j handlers
	GOARCH=amd64 GOOS=linux go build -o /tmp/delve/dlv github.com/derekparker/delve/cmd/dlv
	sam local start-api -d 5986 --debugger-path /tmp/delve

dev:
	make -j dev-watch dev-sam
dev-sam:
	sam local start-api -p 3001 --env-vars json/env.json
dev-watch:
	watchexec -f '*.go' 'make -j handlers'

HANDLERS=$(addsuffix main,$(wildcard handlers/*/))
$(HANDLERS): handlers/%/main: *.go handlers/%/main.go
	cd ./$(dir $@) && GOOS=linux go build -gcflags="${GCFLAGS}" -o main .

handlers: handlers-go
handlers-go: $(HANDLERS)

test:
	go test -v ./...

dynamo-start:
	make docker-up create-table

docker-up:
	docker-compose up -d --no-recreate

create-table:
	-aws dynamodb create-table --cli-input-json file://create-table.json --endpoint-url http://localhost:8000 > /dev/null

dynamo-stop:
	docker-compose down