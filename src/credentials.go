package imosaws

import (
	"appengine"
	"appengine/urlfetch"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
)

func GetAwsConfig(r *http.Request) *aws.Config {
	config := *aws.DefaultConfig
	config.Region = AwsRegion
	config.Credentials =
	    credentials.NewStaticCredentials(AwsKey, AwsSecretKey, "")
	config.HTTPClient = urlfetch.Client(appengine.NewContext(r))
	return &config
}
