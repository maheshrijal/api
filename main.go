package main

import "github.com/aws/aws-lambda-go/lambda"

func sayHello() (string, error){
	return "Hello @maheshrjl from λ!", nil
}

func main() {
	lambda.Start(sayHello)
}