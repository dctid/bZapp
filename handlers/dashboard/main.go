package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dctid/bZapp"
)

func main() {
	lambda.Start(bZapp.NotifyAPIGateway(bZapp.Dashboard))
}
