package main

import (
	"context"
	"strings"
	"testing"

	pb "github.com/peterdeme/go-app/generatedcode"
	"github.com/stretchr/testify/assert"
)

var server = GreeterServer{}

func TestSayHello(t *testing.T) {
	name := "Mike"
	resp, err := server.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	assert.Nil(t, err)
	assert.True(t, strings.Contains(resp.Message, name))
}
