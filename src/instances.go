package main

import (
	"appengine"
	"appengine/mail"
	"fmt"
	"github.com/aws/aws-sdk-go/service/ec2"
	"net/http"
)

func init() {
	http.HandleFunc("/instances", handleInstances)
	http.HandleFunc("/cron/instances", handleCronInstances)
}

func getInstanceStates(r *http.Request) map[string]int {
	svc := ec2.New(GetAwsConfig(appengine.NewContext(r)))
	resp, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		panic(err)
	}

	states := map[string]int{}
	for _, region := range resp.Regions {
		svc := ec2.New(GetAwsRegionalConfig(
			appengine.NewContext(r), *region.RegionName))
		resp, err := svc.DescribeInstances(&ec2.DescribeInstancesInput{})
		if err != nil {
			panic(fmt.Errorf("Region %s error: %s", *region.RegionName, err))
		}

		for _, reservation := range resp.Reservations {
			for _, instance := range reservation.Instances {
				states[*instance.State.Name]++
			}
		}
	}
	return states
}

func handleInstances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	for state, count := range getInstanceStates(r) {
		fmt.Fprintf(w, "%s: %d\n", state, count)
	}
}

func handleCronInstances(w http.ResponseWriter, r *http.Request) {
	states := getInstanceStates(r)
	msg := &mail.Message{
		Sender:  "Ninetan <ninetan@imoz.jp>",
		To:      []string{"i@imoz.jp"},
		Subject: fmt.Sprintf("%d EC2 instances are running", states["running"]),
		Body:    fmt.Sprintf("%v", states)}
	if err := mail.Send(appengine.NewContext(r), msg); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "OK")
}
