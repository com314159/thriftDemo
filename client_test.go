package main

import (
	"testing"
	"thriftDemo/gen-go/demo"
)

//GetUser
func TestGetUser(t *testing.T) {
	client := GetClient()
	rep, err := client.GetUser(ctx, 100)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}

//SayHello
func TestSayHello(t *testing.T) {
	client := GetClient()

	user := &demo.User{}
	user.Name = "thrift"
	user.Address = "address"

	rep, err := client.SayHello(ctx, user)
	if err != nil {
		t.Errorf("thrift err: %v\n", err)
	} else {
		t.Logf("Recevied: %v\n", rep)
	}
}
