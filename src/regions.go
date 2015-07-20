package main

import (
	"appengine"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"net/http"
)

func init() {
	http.HandleFunc("/regions", handleRegions)
}

func handleRegions(w http.ResponseWriter, r *http.Request) {
	svc := ec2.New(GetAwsConfig(appengine.NewContext(r)))
	resp, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "text/plain")
	for _, region := range resp.Regions {
		fmt.Fprintf(w, "%s\n", *region.RegionName)
	}
}
