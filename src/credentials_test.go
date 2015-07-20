package main

import (
	"appengine/aetest"
	"testing"
)

func TestGetAwsConfig(t *testing.T) {
	c, err := aetest.NewContext(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()

	GetAwsConfig(c)
}
