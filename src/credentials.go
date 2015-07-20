package main

import (
	"appengine"
	"appengine/urlfetch"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetAwsRegionalConfig(c appengine.Context, region string) *aws.Config {
	config := *aws.DefaultConfig
	config.Region = region
	config.Credentials =
		credentials.NewStaticCredentials(AwsKey, AwsSecretKey, "")
	config.HTTPClient = urlfetch.Client(c)
	return &config
}

func GetAwsConfig(c appengine.Context) *aws.Config {
	return GetAwsRegionalConfig(c, AwsRegion)
}
