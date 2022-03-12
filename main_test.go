package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSample(t *testing.T) {
	statusCode, _ := GetPolarisshare()
	assert.Equal(t, statusCode, 200)
}

func TestHeader(t *testing.T) {
	_, header := GetPolarisshare()
	assert.NotEqual(t, header, nil)
}
