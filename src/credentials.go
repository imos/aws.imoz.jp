package main

import (
	"appengine"
	"appengine/urlfetch"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetAwsConfig(c appengine.Context) *aws.Config {
	config := *aws.DefaultConfig
	config.Region = AwsRegion
	config.Credentials =
		credentials.NewStaticCredentials(AwsKey, AwsSecretKey, "")
	config.HTTPClient = urlfetch.Client(c)
	return &config
}
